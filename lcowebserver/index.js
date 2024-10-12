// Importing required modules using 'import' syntax
import express from 'express';

const app = express();
const port = 8000;

// Middlewares to parse JSON and URL-encoded data
app.use(express.json());
app.use(express.urlencoded({ extended: true }));

// GET route for the root URL
app.get('/', (req, res) => {
  res.status(200).send("Welcome to ShardenduMishraCreate's server");
});

// GET route for '/get'
app.get('/get', (req, res) => {
  res.status(200).json({ message: "Hello from ShardenduMishraCreate's.in" });
});

// POST route to handle JSON data
app.post('/post', (req, res) => {
  let myJson = req.body; // Accessing the JSON sent in the request body
  res.status(200).send(myJson);
});

// POST route to handle form data
app.post('/postform', (req, res) => {
  res.status(200).send(JSON.stringify(req.body)); // Sending form data back as JSON
});

// Start the server
app.listen(port, () => {
  console.log(`Example app listening at http://localhost:${port}`);
});
