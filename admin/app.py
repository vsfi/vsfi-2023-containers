from flask import Flask, render_template, redirect, request
import psycopg2
import psycopg2.extras
import os

app = Flask(__name__)
db_url = os.getenv("DATABASE_URL", "postgres://root@db:26257/beer?sslmode=disable")
conn = psycopg2.connect(
    db_url,
    application_name="$ docs_simplecrud_psycopg2",
    cursor_factory=psycopg2.extras.RealDictCursor,
)


@app.route("/")
def queue():
    with conn.cursor() as cur:
        cur.execute("SELECT * FROM beer_queues")
        rows = cur.fetchall()
    conn.commit()
    queue = list()
    for row in rows:
        print(row)
        queue.append(dict(row))
    return render_template("queue.html", queue=queue)


@app.route("/status/<_id>")
def done(_id):
    state = request.args.get("state")
    print(_id, state)
    with conn.cursor() as cur:
        cur.execute(f"UPDATE beer_queues SET status= '{state}' WHERE id='{_id}'; ")
    conn.commit()
    return redirect("/")


if __name__ == "__main__":
    app.run(host='0.0.0.0')
