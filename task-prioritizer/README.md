python -m venv venv / python3 -m venv venv
source venv/bin/activate
pip install -r requirements.txt
go run cmd/main.go
uvicorn cmd.main:app --reload
после изменений:
pip freeze > requirements.txt
