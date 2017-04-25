# go-base

## What kind of project are you proposing?
go-base, a Go version of hack4impact/flask-base, “A simple Flask boilerplate app with SQLAlchemy, Redis, User Authentication, and more.”

## What do you hope to learn or gain from the project?
* Practice in building web apps in Go
* Understanding the differences between building boilerplate app in Flask/Python and Go 
* Creating a boilerplate app in Go, in case future Hack4Impact projects need the concurrent functionality of Go

## What are your concrete goals or milestones? How will we judge the completeness and success of your project?
### “What’s included?” in flask-base (bolding priority features for go-base)
* Blueprints
* User and permissions management
* Flask-SQLAlchemy for databases
* Flask-WTF for forms
* Flask-Assets for asset management and SCSS compilation
* Flask-Mail for sending emails
* gzip compression
* Redis Queue for handling asynchronous tasks
* ZXCVBN password strength checker
* CKEditor for editing pages

### “Demos” for flask-base (first milestone for completeness)
* Home Page
* Registering User
* Admin Editing Page
* Admin Editing Users

## If you’re contributing to an open source project, what issues are you addressing? What features are you implementing?
* Project will be open source

##	What libraries will you use?
* net/http
HTTP client and server implementations
* net/url
parses URLs and implements query escaping
* text/template
data-driven templates for generating textual output
* gorilla/handlers
collection of handlers (aka "HTTP middleware") for use with Go's net/http package (or any framework supporting http.Handler)
*	gorilla/mux
request router and dispatcher for matching incoming requests to their respective handler
*	justinas/alice
Convenient way to chain your HTTP middleware functions and the app handler

##	What will the interface look like (if applicable)? This includes a GUI, CLI, or programmatic (API) interface depending on your type of project.
* GUI: see “Demos” gifs 

