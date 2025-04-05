# Go Language Courses API

This is a REST API built using **Go** and **PostgreSQL** for managing language courses. The API supports full CRUD (Create, Read, Update, Delete) operations for courses, which include information like course name, description, language, and duration.

## Table of Contents
1. [Project Description](#project-description)
2. [Features](#features)
3. [Tech Stack](#tech-stack)
4. [Setup Instructions](#setup-instructions)
    - [Clone the Repository](#clone-the-repository)
    - [Install Dependencies](#install-dependencies)
    - [Setup Database](#setup-database)
    - [Run the Application](#run-the-application)
5. [API Endpoints](#api-endpoints)
    - [POST /courses](#post-courses)
    - [GET /courses](#get-courses)
    - [GET /courses/:id](#get-coursesid)
    - [PUT /courses/:id](#put-coursesid)
    - [DELETE /courses/:id](#delete-coursesid)
6. [License](#license)

## Project Description

This project provides a simple backend API to manage language courses. It uses **Go** for the backend and **PostgreSQL** for the database. The application allows users to:
- Create new courses
- Get the list of all courses
- Get a course by its ID
- Update an existing course
- Delete a course from the system

## Features
- Full CRUD operations for language courses.
- Basic validation and error handling for inputs.
- Soft delete feature via the `deleted_at` field.
- API built with the Go programming language.

## Tech Stack
- **Go**: The backend framework used to build the API.
- **PostgreSQL**: The relational database used to store the courses.
- **Gin**: A lightweight web framework for Go used to handle HTTP requests.
- **GORM**: ORM (Object-Relational Mapper) used to interact with the PostgreSQL database.

## Setup Instructions

### Clone the Repository
First, clone the repository to your local machine:

```bash
git clone https://github.com/AlbinaKhassanova/go-language-courses.git
cd go-language-courses
