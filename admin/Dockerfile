FROM registry-01.vsfi.org/python:3-alpine

RUN apk add --no-cache --virtual .build-deps \
    gcc \
    python3-dev \
    musl-dev \
    postgresql-dev \
    && pip install --no-cache-dir psycopg2 \
    && apk del --no-cache .build-deps
 
WORKDIR /app
COPY requirements.txt .
RUN pip3 install -r requirements.txt

COPY app.py .
COPY templates/ templates/

CMD [ "python3", "app.py" ]