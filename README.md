# Truncate

### A simple URL shortener built with Angular(Typescript) and Gin(Go)

To run the application:
- Clone the repository
- Setup an Auth0 application to manage logging in and out
- Setup the required environment variables in a .env file
  - AUTH0_DOMAIN
  - AUTH0_CLIENT_ID
  - AUTH0_CLIENT_SECRET
  - AUTH0_CLIENT_CALLBACK_URL
  - DB_NAME
  - DB_USER
  - DB_PASSWORD
  - ROUTER (0.0.0.0 for deployment to the web)
  - HANDLER (website URL)
- Save SSL certificate and private key as "cert.pem" and "key.pem"
- Build the Angular frontend in the web folder using "ng build"
  - Requires installing nodejs and npm
- Use "docker-compose up --build" to build and run the application
