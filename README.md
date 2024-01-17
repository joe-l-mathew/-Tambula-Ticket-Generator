# Ticket Generation System

A simple Go backend application for generating and managing tickets.

[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

## Table of Contents

- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
- [Usage](#usage)
  - [Generating Tickets](#generating-tickets)
  - [Fetching Tickets](#fetching-tickets)


## Getting Started

### Prerequisites

Before you begin, ensure you have the following installed:

- [Go](https://golang.org/doc/install)

### Installation

1. Clone the repository:

    ```bash
    git clone https://github.com/your-username/ticket-generation-system.git
    ```

2. Change to the project directory:

    ```bash
    cd ticket-generation-system
    ```

3. Install dependencies and build the project:

    ```bash
    go get
    go build main.go
    ```

4. Run the application:

    ```bash
    ./main
    ```

## Usage

The application exposes endpoints for generating and fetching tickets.

### Generating Tickets

To generate a ticket, make a GET request to:

```bash
http://localhost:8080/api/generate_ticket
```

### Fetching Tickets

To get  tickets, make a GET request to:
page is the page number
itemsPerPage is the number of item in each page

```bash
http://localhost:8080/api/tickets?page=1&itemsPerPage=2
```
