const express = require('express');
const router = express.Router();

let tasks = [
  { id: 1, title: "Learn Node.js", completed: false },
  { id: 2, title: "Build a project", completed: false }
];

// Get all tasks
router.get('/', (req, res) => {
  res.json(tasks);
});

// Add a new task
router.post('/', (req, res) => {
  const newTask = {
    id: tasks.length + 1,
    title: req.body.title,
    completed: false
  };
  tasks.push(newTask);
  res.status(201).json(newTask);
});

// Update a task
router.put('/:id', (req, res) => {
  const task = tasks.find(t => t.id == req.params.id);
  if (!task) return res.status(404).json({ message: "Task not found" });

  task.title = req.body.title || task.title;
  task.completed = req.body.completed ?? task.completed;
  res.json(task);
});

// Delete a task
router.delete('/:id', (req, res) => {
  tasks = tasks.filter(t => t.id != req.params.id);
  res.json({ message: "Task deleted" });
});

module.exports = router;

