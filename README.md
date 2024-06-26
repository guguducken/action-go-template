# Action Golang Template

This repository is a template repository for writing GitHub Actions in Golang. It is based on the [GitHub Action documentation](https://docs.github.com/en/actions/using-workflows/workflow-commands-for-github-actions) and references the [GitHub Node.js SDK](https://github.com/actions/toolkit/tree/main/packages/core).

In a sense, the way this repository is implemented has a bit of action-as-code flavor to it。

## Usage
### Use this template repository to create a repository
First. You need to use this template repository to create a repository, the details of which can be found in the [official GitHub documentation](https://docs.github.com/en/repositories/creating-and-managing-repositories/creating-a-repository-from-a-template)

### Pulling created repositories to local
You can pull the created repository locally using the following command
~~~bash
mkdir action-go-template;
cd action-go-template;
git init;
git remote add origin git@github.com:guguducken/action-go-template.git;
git pull origin main;
~~~
Just replace the repository information with your own.

### Create action.yaml
Personally I think this section is the key to what this repository can be called action-as-code, you can go through the following steps by modifying the code and generating action.yaml

1. Modifying information in the `Makefile`
    
    There is some information defined in the `Makefile` that is used to generate the `go module` and the `action yaml`, which needs to be changed to the actual repository information
    ~~~Makefile
    ...
    # You must modify it to yourself information
    AUTHOR=guguducken
    REPOSITORY=action-go-template
    DESCRIPTION=A template repository for creating action written in go language.
    ICON=arrow-up
    COLOR=blue
    ...
    ~~~
    For example, it can be modified to this
    ~~~Makefile
    ...
    # You must modify it to yourself information
    AUTHOR=guguducken
    REPOSITORY=upload-artifact-oss
    DESCRIPTION=Upload an artifact to oss.
    ICON=arrow-up
    COLOR=blue
    ...
    ~~~
2. Initialization go module settings
   
   After you have made the above changes, you can use the following commands to complete the initialization of the `go module`
   ~~~bash
   make preconfig
   ~~~
   After executing this command, you need to make sure that the information has been changed
3. Setting up GitHub action inputs, outputs, and other information
   
   This information is defined in the file `optools/action/action-helper.go`, the fields are annotated and can be modified according to the annotations.
   
   After modifying the corresponding information, you can generate the action yaml with the following command.
   ~~~bash
   make action
   ~~~
   If there are changes, the action yaml can be regenerated by running this command again
4. Add the code which needs

   Just add the actual code, `main.go` is the only entry point for the code.

   Finally, the binary can be generated by following command
   ~~~bash
   make build
   ~~~

5. Publish your action