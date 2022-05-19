const express = require('express');
const bodyParser = require('body-parser');
const mariadb = require('mariadb');
const session = require('express-session');
const SECRET = require('crypto').randomBytes(64).toString('hex');
const FL4G = require('./flag.json');

const app = express();

app.set('view engine', 'ejs');
app.set('trust proxy', 'uniquelocal');
app.use(bodyParser.urlencoded({ extended: true }));

app.use(
  session({
    secret: SECRET,
    saveUninitialized: false,
    resave: true,
    cookie: {
      expires: 120000,
    },
  })
);

const pool = mariadb.createPool({
  host: 'db',
  user: process.env.MARIADB_USER,
  password: process.env.MARIADB_PASSWORD,
  database: process.env.MARIADB_DATABASE,
  connectionLimit: 50,
});

function auth(req, res, next) {
  if (req.session.user) {
    next();
  } else {
    return res.redirect('/login');
  }
}

app.post('/login', async (req, res) => {
  let conn;
  const { username, password } = req.body;

  try {
    conn = await pool.getConnection();
    const users = await conn.query(
      `SELECT * FROM users WHERE username='${username}' AND password='${password}'`
    );
    if (users[0]) {
      req.session.user = users[0].username;
      return res.redirect('/');
    }
    res.render('login', { failed: true });
  } catch (err) {
    console.log(err);
    res.render('login', { failed: true });
  } finally {
    if (conn) conn.end();
  }
});

app.get('/login', (req, res) => {
  if (req.session.user) res.redirect('/');
  else res.render('login', { failed: false });
});

app.get('/', auth, (req, res) => {
  const username = req.session.user;
  res.render('index', { username, FL4G });
});

app.listen(3000);
