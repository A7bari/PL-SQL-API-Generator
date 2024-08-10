# PL/SQL Package Generator in Go

This project is a Go-based tool designed to **automate the generation of PL/SQL package** containing CRUD (Create, Read, Update, Delete) operations for tables defined in DDL (Data Definition Language) scripts. The tool consists of two main components: a **parser** and a **generator**.

## Overview

This tool automates the creation of PL/SQL package specifications for tables defined in DDL scripts. The tool consists of two main parts:

- ### Parser:

Reads the DDL of tables and extracts the tables along with their columns.

- ### Generator:

Generates a PL/SQL package for each table, containing procedures for insert, update, and get operations.

The tool is particularly useful for developers working in environments where PL/SQL is heavily used, as it simplifies and automates the repetitive task of creating standard CRUD procedures for database tables.

## Features

- DDL Parsing: Automatically parse DDL scripts to identify tables and their associated columns, the parser uses the State pattern to handle different states of the parsing process.

```go
// StateFn represents a function that processes input and returns the next state.
type StateFn func(*Parser) StateFn
```

- Package Generation: Generate consistent PL/SQL package specifications/body for each table, including basic CRUD operations.

- Extensible Design: Can be easily extended to include more complex procedures or additional database operations.
