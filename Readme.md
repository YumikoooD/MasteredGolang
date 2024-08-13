# Mastered Golang!

## Overview

Welcome to the **Mastered Golang** repository! This project is dedicated to enhancing my skills in Golang. The repository currently consists of two main sections: the "SimpleGoReminder" folder, which covers basic language concepts, and the "RestAPI" folder, which contains a fully functional REST API application. More projects will be added in the future as I continue to expand my knowledge and expertise in Golang.

## SimpleGoReminder

This folder includes simple and concise examples demonstrating fundamental Go concepts such as variables, functions, loops, and more. It's designed to serve as a quick reference or refresher on core Golang principles.

## RestAPI

The **RestAPI** project is built using the Gin framework and SQLite3. It provides a basic yet powerful example of how to create a RESTful API in Golang. The API supports full CRUD operations and includes user authentication with hashed passwords. Below are the main functionalities:

### API Endpoints

- **`GET /events`**: Retrieve a list of all events.
- **`GET /events/:id`**: Retrieve details of a specific event by ID.
- **`POST /events`**: Create a new event.
- **`PUT /events/:id`**: Update an existing event by ID.
- **`DELETE /events/:id`**: Delete an event by ID.
- **`POST /signup`**: Register a new user.
- **`POST /login`**: Authenticate a user and issue a token.

### Framework

- **Gin Framework**: A lightweight and high-performance web framework for building RESTful APIs.
- **SQLite3**: A simple and self-contained SQL database engine used for storing data.
- J

### The programm

The server will start on http://localhost:8080, and you can use tools like Postman or curl to interact with the API.