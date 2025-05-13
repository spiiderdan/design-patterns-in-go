# Design Patterns in Go 🧠

This repository contains examples of classic software design patterns implemented in Go, organized by category.

Each pattern includes:
- A clear, commented Go implementation
- A simple real-world analogy
- When to use it (and when not to)

## 📦 Categories

### Creational
- [Singleton](creational/singleton) – Ensures only one instance is used (e.g., Redis connection)
- [Factory Method](creational/factory) – Powers a URL shortener that selects between random and hash-based generation strategies
- [Builder](creational/builder)

### Structural
- [Adapter](structural/adapter)
- [Decorator](structural/decorator)

### Behavioral
- [Observer](behavioral/observer)
- [Strategy](behavioral/strategy)

## 🧪 Featured Demo

- ✨ [Factory Method (URL Shortener)](creational/factory):  
  Uses the Factory pattern to create short URLs with either a random or hashed strategy, depending on input. Demonstrates how to abstract creation logic and simplify extensibility.

## 🧑‍💻 Why this project?

I'm learning and practicing design patterns by building them one by one in Go — this is both a personal learning log and a portfolio project. Each pattern is paired with a small real-world example to make it easier to understand and apply.

## 🛠 Language
- Go (Golang)

## 📩 Contributions
Feel free to fork, star, or open an issue with ideas or suggestions!

