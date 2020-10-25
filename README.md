Quote Of The Day Service
===

* [Usage](#usage)

Implementation of [RFC-865](https://tools.ietf.org/html/rfc865).

```



                       Quote of the Day Protocol




This RFC specifies a standard for the ARPA Internet community.  Hosts on
the ARPA Internet that choose to implement a Quote of the Day Protocol
are expected to adopt and implement this standard.

A useful debugging and measurement tool is a quote of the day service.
A quote of the day service simply sends a short message without regard
to the input.

TCP Based Character Generator Service

   One quote of the day service is defined as a connection based
   application on TCP.  A server listens for TCP connections on TCP port
   17.  Once a connection is established a short message is sent out the
   connection (and any data received is thrown away).  The service
   closes the connection after sending the quote.

UDP Based Character Generator Service

   Another quote of the day service is defined as a datagram based
   application on UDP.  A server listens for UDP datagrams on UDP port
   17.  When a datagram is received, an answering datagram is sent
   containing a quote (the data in the received datagram is ignored).

Quote Syntax

   There is no specific syntax for the quote.  It is recommended that it
   be limited to the ASCII printing characters, space, carriage return,
   and line feed.  The quote may be just one or up to several lines, but
   it should be less than 512 characters.
```

# Usage

Complete the application with your quotes in `internal/quotes.go`.
That's all it needs. Either run the application locally or build the Docker image afterwards.

```bash
$ go run cmd.go
```

```bash
$ curl -v telnet://localhost:17
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0*   Trying ::1:17...
* Connected to localhost (::1) port 17 (#0)
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0She isn't human; She is art, with a heart.
* Closing connection 0
```
