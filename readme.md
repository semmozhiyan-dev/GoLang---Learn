# Go Conference Booking Application 🎟️

A simple command-line conference ticket booking application built with Go.

This project demonstrates fundamental Go programming concepts such as:

* Variables and Constants
* User Input
* Functions
* Loops
* Conditionals
* Slices
* String Manipulation
* Validation
* Switch Statements

---

## Features

✅ Book conference tickets

✅ Validate user input

✅ Check ticket availability

✅ Store bookings using slices

✅ Display remaining tickets

✅ Extract and display first names of attendees

✅ Stop bookings when tickets are sold out

---

## Project Structure

```text
.
├── main.go
└── README.md
```

---

## Technologies Used

* Go (Golang)
* Standard Library

  * fmt
  * strings

---

## How It Works

1. Displays conference information.
2. Accepts user details:

   * First Name
   * Last Name
   * Email
   * Number of Tickets
3. Validates user input.
4. Books tickets if valid.
5. Updates remaining ticket count.
6. Displays all attendee first names.
7. Ends when all tickets are sold.

---

## Validation Rules

### Name Validation

* First name must contain at least 2 characters.
* Last name must contain at least 2 characters.

### Email Validation

* Email must contain `@`.

### Ticket Validation

* Ticket count must be greater than 0.
* Ticket count must not exceed available tickets.

---

## Running the Project

### Clone Repository

```bash
git clone https://github.com/your-username/go-conference-booking.git
```

### Navigate to Project

```bash
cd go-conference-booking
```

### Run Application

```bash
go run main.go
```

---

## Example Output

```text
Welcome to Go Conference booking application!
We have a total of 50 tickets and 50 are still available.

Enter your first name: Sem
Enter your last name: Kumar
Enter your email: sem@example.com
Enter number of tickets: 2

Thank you Sem Kumar for booking 2 tickets.
You will receive a confirmation email at sem@example.com

48 tickets remaining for Go Conference

First names of bookings: [Sem]
```

---

## Concepts Practiced

* Functions
* Function Parameters
* Multiple Return Values
* Global Variables
* Loops
* Conditional Statements
* Slices
* String Operations
* User Input Handling
* Basic Application Design

---

## Future Improvements

* Use Structs for User Data
* Store Data in JSON Files
* Add Goroutines for Email Simulation
* Add Unit Tests
* Create REST API Version
* Connect to Database
* Build Web Interface

---

## Author

Sem

Aspiring DevOps Engineer | Learning Go, Docker, Kubernetes, AWS, CI/CD and Cloud Technologies 🚀
