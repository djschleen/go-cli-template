![](img/go-cli-template128x128.png)

# go-cli-template

[![Go Report Card](https://goreportcard.com/badge/github.com/djschleen/go-cli-template)](https://goreportcard.com/report/github.com/djschleen/go-cli-template)

A go CLI template using cobra/spf

## Quickstart

After using the template to create your new repository, you should be good to go and can just run ```make build``` to make the ```go-cli-template``` binary. 

To test, run the following in your local repository folder:

``` bash
./go-cli-template hello
```

## Initializing Hookz

This repository has a pre-commit action pipeline in it that can be used with [Hookz](https://github.com/devops-kung-fu/hookz). Use the instructions there to install the ```hookz``` command and then execute the following in your local repository folder:

``` bash
hookz init --verbose --verbose-output
```
Now you will have a pre-commit action pipeline that checks go code quality, lints, runs cyclomatic complexity checks, and runs test cases before any code gets committed to the remote repository. If there is a problem, ```hookz``` will stop the commit process and let you address issues.

## Credits

A big thank-you to our friends at [Smashicons](https://www.flaticon.com/authors/smashicons) for the ```go-cli-template``` logo.

