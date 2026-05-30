from fastapi import FastAPI, Request, Depends
from typing import Annotated
from fastapi.responses import HTMLResponse, RedirectResponse
from fastapi.templating import Jinja2Templates
from fastapi.staticfiles import StaticFiles
from internal.models.models import *
from internal.database.db import *
from hashlib import sha256

Base.metadata.create_all(bind=engine)
app = FastAPI()
templates = Jinja2Templates(directory="templates/fastapi")
app.mount("/static", StaticFiles(directory="static"), name="static")

def get_db():
    db = SessionLocal()
    try:
        yield db
    finally:
        db.close()

db_dependency = Annotated[Session, Depends(get_db)]

@app.get("/login", response_class=HTMLResponse)
def login(request: Request):
    return templates.TemplateResponse(
        "login.html",
        {"request": request}
    )

@app.post("/login")
async def login_post(request: Request, db: db_dependency):
    form = await request.form()
    username = form.get("username")
    password = form.get("password")
    hashed = sha256(password.encode()).hexdigest()

    user = db.query(User).filter(
        User.name == username,
        User.password == hashed
    ).first()
    if user is None:
        return {"error": "Неверное имя пользователя или пароль"}
    response = RedirectResponse(
        url="/tasks",
        status_code=303
    )
    response.set_cookie(
        key="user_id",
        value=str(user.id),
        httponly=True,
        path="/"
    )
    return response

@app.get("/register", response_class=HTMLResponse)
def registration(request: Request):
    return templates.TemplateResponse(
        "registration.html",
        {"request": request}
    )

@app.post("/register")
async def registration(request: Request, db: db_dependency):
    form = await request.form()
    username = form.get("username")
    password = form.get("password")
    if db.query(User).filter(User.name == username).first() is not None:
        return {"error": "Пользователь с таким именем уже существует"}
    try:
        new_user = User(name=username, password=sha256(password.encode()).hexdigest())
        db.add(new_user)
        db.commit()
        db.refresh(new_user)
        response = RedirectResponse(
        url="/tasks",
        status_code=303
        )
        response.set_cookie(
        key="user_id",
        value=str(new_user.id),
        httponly=True
        )
        return response
    except Exception as e:
        return {"error": str(e)}

@app.get("/logout")
def logout():
    response = RedirectResponse(
        url="/login",
        status_code=303
    )
    response.delete_cookie(key="user_id", path="/")
    return response