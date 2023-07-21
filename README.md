# Confluence to Markdown Convertor
## Introduction

This is a fun project created as part of learning `GO` and also to convert some of my Confluence documents to Markdown. It is a simple Web-Based Go application which lets you download a Markdown file from a COnfluence Page.

## Installation

The project can run locally with no much additional installations. You need to have [Go](https://go.dev/doc/install) installed. You can clone the project [here](https://github.com/vinivasundharan/confluence-parser).

```
cd markdown
go build -o conf2md main.go
./conf2md
```

The application would be available on localhost:8080/conf2md

## Components

### Confluence API

The confluence API is used to get the information, content from the page. The API is authenticated with Basic Authentication. The functions related to the API are packaged in the module `confluence`.

### HTML to Markdown package

The go package [html-to-markdown](https://github.com/JohannesKaufmann/html-to-markdown) is used and extended for generating the desired Markdown from Confluence HTMl. The package gives the possibility to add new rules which makes it easier to add conversions for confluence macros. There are extensions for macros info and Links for Confluence pages (This does not have the implementation completed as of writing).

The extensions are packaged with `custom_md`. It is designed in a way that adding new rules are easy and maintainable.

### Frontend

A basic HTML is rendered using `html/template` .

## Adding a new rule

There are 2 ways of adding a rule.

1. Using regex: The macro tag from confluence content can be replaced with appropriate HTML tags before passing it to the html-to-markdown module.

2. Adding a md Rule: A new rule can be added filtering the macro and adding appropriate markdown syntax for a particular filter (This works based on HTML nodes)


A combination of both is also possible as done for `INFOBOX` where the confluence macro is converted to a custom html tag using regex and a new rule is added for the custom tag.

## Limitations

Since it is a fun project there are a lot of limitations and improvements required

- Not all errors are handled. (Unauthorized and Not Found for instance)

- The confluence URL parsing is not enabled, the content ID is the only parameter that can work

- Confluence link on pages do not return the right URL

- Many macros are not handled


This file was generated using the tool :-) 