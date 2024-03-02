**StayEase: Room Management System**

StayEase is a comprehensive Room Management System designed to streamline your hospitality operations. With a focus on simplicity and efficiency, StayEase offers essential features to manage room bookings, check-ins, check-outs, room availability, and billing processes. Whether you operate a hotel, hostel, or any lodging establishment, StayEase provides a reliable solution to enhance your guest management experience.

**Key Features:**

1. **Room Booking:** Easily manage room reservations with an intuitive booking system. Allow guests to book rooms hassle-free, ensuring a smooth reservation process.

2. **Check-in and Check-out:** Streamline guest arrivals and departures with efficient check-in and check-out functionalities. Keep track of guest details for a personalized and seamless experience.

3. **Room Availability:** StayEase provides real-time information on room availability, helping you optimize room occupancy and plan for peak periods.

4. **Billing:** Simplify billing processes with StayEase's integrated billing system. Generate accurate and transparent invoices for your guests, enhancing overall customer satisfaction.

**Getting Started:**

To set up StayEase, follow these steps:

1. **Installation:**
   ```bash
   go get -u github.com/bsaii/stay-ease
   ```

2. **Configuration:**
   Configure your database settings, room types, and other preferences in the `config` file.

3. **Usage:**
   Start using StayEase to manage your rooms efficiently. Refer to the documentation for detailed instructions on each feature.

**Contributing:**

Contributions are welcome! If you have ideas for improvements or new features, feel free to open an issue or submit a pull request.

**License:**

This project is licensed under the MIT License - see the [LICENSE](link-to-license) file for details.

StayEase empowers you to elevate your room management processes, providing a reliable and user-friendly solution for your lodging establishment.


<!-- For a minimal hostel or hotel management system in Go, you can start with a few key features and corresponding data structures (structs). Here are some suggested features along with the associated structs:

### Features:

1. **User Registration:**
   - Allow users to register with the system.

2. **Room Booking:**
   - Enable users to book rooms for a specified duration.

3. **Check-in/Check-out:**
   - Record user check-ins and check-outs.

4. **Room Availability:**
   - Provide information on room availability for a given date range.

5. **Billing:**
   - Generate bills for users based on their stay duration and room type.





Building user authentication in Go typically involves the following steps:

1. **Create a User Struct:**
   Define a struct that represents the user with fields such as `ID`, `Username`, `Email`, and `Password`. You may also include other fields based on your requirements.

    ```go
    type User struct {
        ID       uint
        Username string
        Email    string
        Password string
        // Other user details
    }
    ```

2. **Password Hashing:**
   When storing passwords, it's essential to hash them for security. Use a reliable password hashing library like `bcrypt` to hash and verify passwords.

    ```go
    import "golang.org/x/crypto/bcrypt"

    func hashPassword(password string) (string, error) {
        hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
        return string(hashedPassword), err
    }

    func comparePasswords(hashedPassword, password string) error {
        return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
    }
    ```

3. **User Registration:**
   Implement a function to handle user registration. Hash the user's password before storing it in the database.

    ```go
    func registerUser(username, email, password string) error {
        hashedPassword, err := hashPassword(password)
        if err != nil {
            return err
        }

        // Save user data to the database, including the hashed password
        // ...
        return nil
    }
    ```

4. **User Login:**
   Implement a function to handle user login. Verify the entered password against the stored hashed password.

    ```go
    func loginUser(email, password string) error {
        // Retrieve user data from the database based on the provided email
        // ...

        // Compare the entered password with the stored hashed password
        if err := comparePasswords(user.Password, password); err != nil {
            return err
        }

        // Authentication successful
        return nil
    }
    ```

5. **Session Management:**
   Implement session management to track authenticated users. You can use a session management library or create your own session mechanism.

6. **Middleware for Authentication:**
   Create middleware to protect routes that require authentication. Check the user's authentication status before allowing access to protected resources.

    ```go
    func isAuthenticated(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            // Check if the user is authenticated
            // ...

            // If authenticated, call the next handler
            next.ServeHTTP(w, r)
        })
    }
    ```

7. **Logout:**
   Implement a logout functionality that clears the user's session.

8. **Password Recovery (Optional):**
   If needed, implement a password recovery mechanism that allows users to reset their passwords.

9. **Security Considerations:**
   Ensure that your authentication implementation follows security best practices, such as protecting against SQL injection, cross-site scripting (XSS), and ensuring secure session management.

10. **Use Secure Connections:**
    Always use HTTPS to encrypt data in transit, especially when dealing with sensitive user information.

Remember that security is crucial, and you should stay informed about best practices and potential security vulnerabilities in authentication systems. Consider using well-established libraries or frameworks for authentication, as they often provide additional security features and have been thoroughly tested.



Certainly! Adding JWT (JSON Web Tokens) to your authentication system can enhance security and provide a stateless way to authenticate users. Additionally, incorporating token refresh functionality helps improve the user experience by extending the validity of authentication without requiring the user to log in again.

Here's a general outline of how you can integrate JWT with token refresh in your Go authentication system:

1. **JWT Generation and Verification:**
   Use a JWT library to generate and verify tokens. A popular choice in the Go ecosystem is the `github.com/dgrijalva/jwt-go` library.

    ```go
    import (
        "github.com/dgrijalva/jwt-go"
    )

    var jwtSecret = []byte("your-secret-key")

    func generateToken(userID uint) (string, error) {
        token := jwt.New(jwt.SigningMethodHS256)
        claims := token.Claims.(jwt.MapClaims)
        claims["userID"] = userID
        claims["exp"] = time.Now().Add(time.Hour * 2).Unix() // Token expiration time

        tokenString, err := token.SignedString(jwtSecret)
        return tokenString, err
    }

    func verifyToken(tokenString string) (*jwt.Token, error) {
        return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
            return jwtSecret, nil
        })
    }
    ```

2. **Token Refresh Functionality:**
   Implement a function to refresh the JWT token. When a user's token is close to expiration, generate a new token and return it.

    ```go
    func refreshToken(tokenString string) (string, error) {
        token, err := verifyToken(tokenString)
        if err != nil {
            return "", err
        }

        if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
            // Check if the token is close to expiration
            expirationTime := int64(claims["exp"].(float64))
            if time.Until(time.Unix(expirationTime, 0)) < time.Minute*5 {
                // Generate a new token with extended expiration time
                newToken, err := generateToken(uint(claims["userID"].(float64)))
                return newToken, err
            }
        }

        return "", errors.New("unable to refresh token")
    }
    ```

3. **Middleware for JWT Authentication:**
   Create middleware to authenticate users using JWT. Extract the token from the request, verify it, and set the user ID in the request context.

    ```go
    func authenticateJWT(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            tokenString := extractTokenFromRequest(r)

            token, err := verifyToken(tokenString)
            if err != nil {
                // Handle invalid token
                http.Error(w, "Unauthorized", http.StatusUnauthorized)
                return
            }

            if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
                userID := uint(claims["userID"].(float64))

                // Set the userID in the request context for use in handlers
                ctx := context.WithValue(r.Context(), "userID", userID)
                r = r.WithContext(ctx)

                // Call the next handler
                next.ServeHTTP(w, r)
            } else {
                // Handle invalid token
                http.Error(w, "Unauthorized", http.StatusUnauthorized)
                return
            }
        })
    }

    func extractTokenFromRequest(r *http.Request) string {
        // Extract the token from the request headers or query parameters
        // ...
        return token
    }
    ```

4. **Token Refresh Endpoint:**
   Create an endpoint for refreshing the JWT token. This endpoint should be protected by the `authenticateJWT` middleware.

    ```go
    func refreshTokenHandler(w http.ResponseWriter, r *http.Request) {
        tokenString := extractTokenFromRequest(r)
        newToken, err := refreshToken(tokenString)
        if err != nil {
            http.Error(w, "Unable to refresh token", http.StatusBadRequest)
            return
        }

        // Send the new token in the response
        w.Write([]byte(newToken))
    }
    ```

5. **Usage in Handlers:**
   Access the user ID from the request context in your route handlers.

    ```go
    func someProtectedHandler(w http.ResponseWriter, r *http.Request) {
        userID := r.Context().Value("userID").(uint)

        // Use the userID for further processing
        // ...
    }
    ```

This is a basic example, and you may need to adapt it based on your specific application structure and requirements. It's crucial to handle token storage securely, consider HTTPS for all communication, and ensure that your application is protected against common security threats.

ID:          1,
		RoomNumber:  "101",
		Type:        "Single",
		Description: "Cozy single room with a view",
		Capacity:    1,
		Price:       75.0,
		IsBooked:    false,
		BookedDates: []Booking{
			{
				ID:           1,
				RoomID:       1,
				UserID:       101,
				CheckInDate:  time.Date(2024, 3, 1, 14, 0, 0, 0, time.UTC),
				CheckOutDate: time.Date(2024, 3, 5, 12, 0, 0, 0, time.UTC),
				TotalCost:    300.0,
			},
			{
				ID:           2,
				RoomID:       1,
				UserID:       102,
				CheckInDate:  time.Date(2024, 4, 10, 15, 0, 0, 0, time.UTC),
				CheckOutDate: time.Date(2024, 4, 15, 11, 0, 0, 0, time.UTC),
				TotalCost:    375.0,
			},
		}, -->