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

## More to Come Later ...
This respository is currently a work in progress. Check back periodically for updates over the autumn semester of 2024...
