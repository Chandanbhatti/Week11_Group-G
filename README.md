**Building a Go API for Current Toronto Time with MySQL Database Logging**

*Set Up MySQL Database:*

Install MySQL and create a new database.

Create a table named time_log with at least two fields: id (primary key) and timestamp.

*API Development:*

Write a Go application with a web server.

Create an API endpoint /current-time that returns the current time in Toronto.

*Time Zone Conversion:*

Use Go's time package to handle the time zone conversion to Toronto's local time.

*Database Connection:*

Connect to your MySQL database from your Go application.

On each API call, insert the current time into the time_log table.

*Return Time in JSON:*

Format the response from the /current-time endpoint in JSON.

*Error Handling:*

Implement proper error handling for database operations and time zone conversions.


**Bonus Challenges**

Implement logging in your Go application to log events and errors.

Create an additional endpoint to retrieve all logged times from the database.
