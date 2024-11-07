# QuickBin
QuickBin is a simple implementation used to store and retrieve text documents. This project is far from feature complete. Feel free to contribute any additional features this system is currently lacking!

## Goals
Support the major data storage platforms along with simple, small developer options.
- [x] Local file storage
- [ ] SQLite
- [ ] MySQL
- [ ] PostgreSQL
- [ ] MongoDB
- [ ] AWS S3
- [ ] Google Cloud Storage
- [ ] Azure Blob Storage
- [ ] DigitalOcean Spaces
- [ ] Backblaze B2 

Abuse-Proof
 - [x] Rate limiting
 - [x] Key length and namespace configuration
 - [ ] IP range blocking
 - [ ] CAPTCHA Support (Cloudflare turnstyle or Google reCAPTCHA)

Simple Setup
 - [ ] Single binary release
 - [ ] Docker image

## Installation
This is currently a work-in-progress project, so there are no releases available. You can build the project from source.
1. Clone the repository:
    ```sh
    git clone https://github.com/savag3life/QuickBin.git
    cd QuickBin
    ```

2. Install dependencies:
    ```sh
    go mod tidy
    ```

3. Configure your settings in `config.json`, a default config will be created by default on first run. Only a single file-options
    ```json
    {
      "host": "localhost",
      "port": 8080,
      "maxUploadSize": 1048576,
      "keyLength": 10,
      "keyNamespace": "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789",
      "rateLimitConfig": {
        "duration": "1m",
        "max": 10
      },
      "storageType": "file",
      "storageOptions": {
        "path": "data"
      }
    }
    ```
## Config Options
- `host`: The host to bind the server to
- `port`: The port to bind the server to
- `maxUploadSize`: The maximum size of the data that can be uploaded
- `keyLength`: The length of the generated key
- `keyNamespace`: The characters to use when generating the key
- `rateLimitConfig`: The rate limit configuration
    - `duration`: The duration of the rate limit
    - `max`: The maximum number of requests allowed in the duration
- `storageType`: The type of storage to use 
- `storageOptions`: The options for the storage type
  
### "File" Storage Options:
- `path`: The path to store the data
```json
{
  ...
    "storageType": "file",
    "storageOptions": {
        "path": "data"
    }
}
```

### "AWS S3" Storage Options:
- `region`: The region of the S3 bucket
- `bucket`: The name of the S3 bucket
- `accessKey`: The access key for the S3 bucket
- `secretKey`: The secret key for the S3 bucket
```json
{
  ...
  "storageType": "aws",
  "storageOptions": {
    "region": "us-east-1",
    "bucket": "bucket-id",
    "accessKey": "my-access-key",
    "secret": "my-secret-key"
  }
}
```

## Usage
1. Run the application:
    ```sh
    go run quick_bin.go
    ```
2. The server will start on the configured host and port.

## API Endpoints
- **POST /**: Save/Write data (rate-limited)
    - Request body: Raw data to be saved
    - Response: Generated key for the saved data

- **GET /raw/:key**: Retrieve raw data for the given key
    - Response: Raw data

## License
This project is licensed under the MIT License.