# Readwise

**Readwise** is a Go-based application inspired by the real Readwise functionality. It is designed to help users manage and review their reading highlights seamlessly. The application parses JSON dumps of Kindle highlights generated using [Bookcision](https://www.bookcision.com/), stores them in a database, and provides daily insights through email.

---

## Features

1. **Highlight Parsing**
   - Upload Kindle highlights in JSON format.
   - Parse and store highlights, including associated book details.

2. **Daily Insights**
   - Automatically sends an email with random highlights every morning.
   - Highlights include book title, authors, and user notes.

3. **User Management**
   - Manage users with unique emails and user profiles.

4. **RESTful API Endpoints**
   - Parse highlights: `/api/v1/users/{userID}/parse-kindle-file` (POST).
   - Send daily insights: `/api/v1/cloud/send-daily-insights` (POST).

---

## Project Structure

```plaintext
├── api.go            # Main API server setup
├── config.go         # Configuration and environment variable handling
├── controller.go     # API controllers to handle requests
├── db.go             # MySQL database initialization and table management
├── mailer.go         # Email notifications using SendGrid
├── sendgrid.go       # SendGrid integration for sending daily insights
├── store.go          # Data storage interface and implementation
├── types.go          # Type definitions for data models
├── main.go           # Application entry point
├── daily.templ       # Email HTML template for daily insights
├── makefile          # Makefile for building and running the application
```

---

## Installation

### Prerequisites
- [Go](https://golang.org/) 1.16 or later
- MySQL database
- [SendGrid API Key](https://sendgrid.com/)

### Steps

1. Clone the repository:
   ```bash
   git clone <https://github.com/youngprinnce/readwise-in.go>
   cd readwise-in-go
   ```

2. Create a `.env` file:
   ```bash
   SENDGRID_API_KEY=<your-sendgrid-api-key>
   SENDGRID_FROM_EMAIL=<your-sendgrid-email>
   ```

3. Initialize and run the application:
   ```bash
   make build
   make run
   ```

---

## API Endpoints

### 1. Parse Kindle File
- **Endpoint:** `/api/v1/users/{userID}/parse-kindle-file`
- **Method:** POST
- **Description:** Upload a Kindle highlights JSON file to parse and store data.
- **Request Body:** Multipart file upload with the key `file`.

### 2. Send Daily Insights
- **Endpoint:** `/api/v1/cloud/send-daily-insights`
- **Method:** POST
- **Description:** Sends random highlights as a daily email to all users.

---

## Email Template

`daily.templ`
- HTML template used to format daily insights emails.

---

## Contributing

1. Fork the repository.
2. Create a new branch (`feature/your-feature-name`).
3. Commit your changes.
4. Open a pull request.

---

## License

This project is licensed under the MIT License. See the LICENSE file for details.
