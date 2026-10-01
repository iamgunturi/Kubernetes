const express = require('express');
const app = express();
const PORT = 3000;

app.use(express.json());

// Import routes
const taskRoutes = require('./routes/tasks');
app.use('/tasks', taskRoutes);

app.get('/', (req, res) => {
  res.send('Welcome to Node.js Project!');
});

app.listen(PORT, () => {
  console.log(`Server running at http://localhost:${PORT}`);
});

