# cc1-term-paper
This repository is where all code for Cloud Computing 1's term paper resides. 

## What is the term paper topic? 
The topic of research is how [GitHub Actions](https://docs.github.com/en/actions/writing-workflows/quickstart) can be used to improve software testing and software integrity by automating these tasks. 

## How does it aim to do this? 

### Workflow
A codebase, most likely written in [Go](https://go.dev/doc/), will be hosted in this repository. [Workflow files](https://docs.github.com/en/actions/writing-workflows/about-workflows#about-workflows) will be used to show how GitHub Actions can be used to automate tasks, such as performing unit tests in a codebase, when given a particular [trigger](https://docs.github.com/en/actions/writing-workflows/choosing-when-your-workflow-runs/triggering-a-workflow#about-workflow-triggers) . 

### Runners
These workflows will also use two different types of [runners](https://docs.github.com/en/actions/writing-workflows/choosing-where-your-workflow-runs/choosing-the-runner-for-a-job) - GitHub hosted and self-hosted. The aim of this is to demonstrate how a developer can create a workflow for their given use case. For example, GitHub hosted runners may be sufficient for unit testing Go code, but may not be sufficient to quickly unit test as multiple libraries may be needed to do this. This is where self-hosted runners may be more ideal to use as the runner isn't tore down after workflow execution terminates. 

### Terraform
[Terraform](https://developer.hashicorp.com/terraform/intro) is a IaC provider that provisions infrastructure via providers, such as cloud platforms. It can be used in the case where the self-hosted runner is deployed in a cloud platform, such as [AWS](https://aws.amazon.com/), and generates cost while being idle - [setup-terraform](https://github.com/hashicorp/setup-terraform) can be used to manage allocation of self-hosted runners as needed using Terraform configuration files, keeping costs to a bare minimum. 

### URL Shortener
The go code for this application, along with the tests and documentation below, were generated using [ChatGPT](https://chatgpt.com/share/67220faf-25c8-8009-8cc7-df50e43b9d12). Click the link to view the chat log used to create this code. 

#### Overview
The URL Shortener is a simple web service written in Go that allows users to create shortened URLs. It provides two main functionalities:

1. URL Shortening: Given a long URL, the service generates a shorter version for easy sharing.
2. URL Redirection: When accessing the shortened URL, the service redirects the user to the original URL.
This project is designed to be simple and uses in-memory storage, making it ideal for local testing or as a foundation for further development with a persistent storage option.

#### Project Structure
```graphql
url_shortener/
│
├── main.go       # Main application code with the URL shortener logic.
├── main_test.go  # Test file for the URL shortener.
└── go.mod        # Module file for dependency management.
```
#### How to Run
##### Prerequisites
- Go installed (version 1.13 or higher recommended)
#### Steps to Run
1. Clone or Download the project folder `url_shortener`.
2. Initialize the Project (if not already initialized) by running:

```bash
Copy code
go mod init url_shortener
```
3. Run the Application:

```bash
Copy code
go run main.go
```
The server will start on http://localhost:8080.

4. Using the URL Shortener:

- Shorten a URL:

    - Send a GET request to http://localhost:8080/shorten?url=<YOUR_URL>
    - Example:
    ```bash
    Copy code
    curl "http://localhost:8080/shorten?url=http://example.com"
    ```
    - This will respond with a shortened URL in the format: http://localhost:8080/{shortURL}
- Access a Shortened URL:

    - Visit http://localhost:8080/{shortURL} (replace {shortURL} with the path from the shortened URL).
    - The service will redirect you to the original URL.
#### Testing the Application
The project includes unit tests to verify the core functionality of the URL shortener.

1. Run Tests:

```bash
go test
```
This command will automatically run all tests in the project and provide a summary of the results.

2. Test Coverage:

- Test for URL Shortening: Ensures that a shortened URL is generated when a valid URL is provided.
- Test for URL Redirection: Verifies that a shortened URL redirects to the correct original URL.
- Test for Non-Existent URLs: Ensures a 404 response when attempting to access a non-existent shortened URL.
3. Expected Test Output:

```
PASS
ok  	url_shortener	0.313s
```
A passing test output confirms that all functions are working as expected.

#### Example
1. Shorten a URL:

```bash
curl "http://localhost:8080/shorten?url=http://example.com"
```
Response:

```bash
Shortened URL: http://localhost:8080/abc123
```
2. Visit Shortened URL: Visiting http://localhost:8080/abc123 will redirect you to http://example.com.

#### Notes
This application stores URLs in memory, so the shortened URLs will be reset each time the server restarts. To make this service more robust, consider implementing persistent storage options, like a database.

## More to Come Later ...
This respository is currently a work in progress. Check back periodically for updates over the autumn semester of 2024...