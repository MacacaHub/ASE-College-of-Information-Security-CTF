import hashlib
import random
import string
import sqlite3


macaca = [
    ("sylvanus", ""),
    ("nemestrina", ""),
    ("maura", ""),
    ("tonkeana", ""),
    ("nigra", ""),
    ("sinica", ""),
    ("thibetana", ""),
    ("mulatta", ""),
    ("cyclopis", "banana"),
    ("fuscata", ""),
    ("leonina", ""),
    ("silenus", ""),
    ("nigriscens", ""),
    ("arctoides", ""),
    ("radiata", ""),
]


def generate_random_password():
    n = random.randint(10, 15)
    random_password = "".join(random.sample(string.printable, n))
    return random_password


def md5(password):
    m = hashlib.md5()
    m.update(password.encode())
    return m.hexdigest()


def main():
    con = sqlite3.connect("./app/files/macaca.db")
    cur = con.cursor()
    cur.execute(
        "create table if not exists `MacacaHub` (`id` integer not null primary key, `name` text, `password` text);"
    )

    for _id, (name, password) in enumerate(macaca):
        if not password:
            password = generate_random_password()
        cur.execute(
            f"insert or ignore into `MacacaHub` (`id`, `name`, `password`) values({_id}, '{name}', '{md5(password)}');"
        )

    con.commit()
    con.close()


if __name__ == "__main__":
    main()
