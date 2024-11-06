Test the API
1. Register new users

Endpoint Post /register

URL: http://localhost:8080/register

Body:
{
"username": "your_username",
"password": "your_password"
}
Description: Endpoint create new users to register. Replace "your_username" and "your_password" with the desired username and password for your user.


2. Login


Endpoint: Post/login

URL: http://localhost:8080/login

Body 

{
"username": "your_username",
"password": "your_password"
}

Description: After successful login, you will receive a JWT token. This token is required for authenticated endpoints.

3. Set Up Authorization in Postman
   1. Go to the Authorization tab in Postman.
   2. Choose Bearer Token.
   3. Paste the token you received from the login response.

4. Add a New Student (Teacher Only)
Endpoint: POST /students/

URL: http://localhost:8080/students/

BODY:
   {
   "name": "Student Name",
   "age": 18,
   "class_id": "class_id_value"
   }
Description: This endpoint allows teachers to add a new student. Make sure to set up the token in the Authorization header.


5. Update Student Information (Teacher Only)
   Endpoint: PUT /students/:id
   URL: http://localhost:8080/students/{student_id}
   Body:
   {
   "name": "Updated Student Name",
   "age": 19
   }
   Description: This endpoint allows teachers to update student details based on the student ID.

6. View Student Details
   Endpoint: GET /students/:id
   URL: http://localhost:8080/students/{student_id}
   Description: This endpoint retrieves the details of a specific student based on their ID. Requires authentication.

7. Search for Students
   Endpoint: GET /students/search
   URL: http://localhost:8080/students/search?name=StudentName
   Description: This endpoint allows you to search for students by name. Requires authentication.

8. Update Grade
   Endpoint: PATCH "/:id/grade
   URL: http://localhost:8080/students/:id/grade
   Description: This endpoint update grade for student.