===================
Go HTTP Server
===================

Simple HTTP server written in Go accepting POST request in "http://127.0.0.1:8080/".

It includes a fron-end application that sends POST requests with JSON data when the user resizes the screen, copy and paste a field and submit the form. To make it simple, tbe APP doesn't check if the user has changed the clipboard from a different source.

How to run
----------

    cd src/web/

    go build web.go

    go run web.go

