# Go API for Testing Purposes

This Go API is designed for testing purposes and provides simple endpoints to interact with. It serves as a basic example of building a RESTful API using Go.

## Endpoints

- `/hello`: Returns a "Hello World!" message.
- `/hello/{name}`: Returns a personalized greeting message using the provided name.
- `/health`: Provides a health check endpoint returning a JSON response `{"status":"ok"}`.

## Building and Running the Image

### Requirements
- Go installed on your local machine to build the Go application.
- Podman installed on your local machine to build and run the container image.

### Building the Image
1. Clone the repository containing the Go API source code.
2. Navigate to the directory containing the source code.
3. Build the container image using Podman CLI:
    ```bash
    podman build -t go-api-image .
    ```

### Running the Image
1. Once the image is built, you can run the container using the following command:
    ```bash
    podman run -d -p 8080:8080 go-api-image
    ```
   This command runs the container in detached mode (`-d`) and forwards port 8080 from the container to port 8080 on the host.

2. Access the API endpoints:
   - Hello World: `http://localhost:8080/hello`
   - Personalized greeting: `http://localhost:8080/hello/YourName`
   - Health check: `http://localhost:8080/health`

### Additional Notes
- Make sure to replace `YourName` with the desired name in the personalized greeting endpoint.
- You can customize the code and endpoints according to your requirements.
- For more detailed instructions on building and running the Go API, refer to the source code and comments within the application files.

