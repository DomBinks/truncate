# Truncate

### A simple URL shortener built with Angular(Typescript) and Gin(Go)

To run the application:
- Clone the repository
- Setup an Auth0 application to manage logging in and out
- Setup the required environment variables in a .env file
  - AUTH_0_CLIENT_ID
  - AUTH_0_CLIENT_SECRET
  - AUTH_0_CLIENT_CALLBACK_URL
  - DB_NAME
  - DB_USER
  - DB_PASSWORD
  - ROUTER
  - HANDLER
- Build the Angular frontend in the web folder using "ng build"
- Use "docker-compose up --build" to build and run the application