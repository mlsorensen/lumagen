# Lumagen Serial Module

This is a quick golang module to establish a serial session to a Lumagen Radiance Pro,
and perform simple interactions. It currently only reads and parses messages from the 
Lumagen.

The monitor code will read serial input line by line and then pass these lines along
to parsers, which can do what they want with the message or ignore it.

Currently, the only parser implemented is the handling of an I22 message per the Lumagen
documentation. This parser identifies a line as an I22 message and parses it into a 
`ZQI22Message` struct. The developer can provide this parser a handler function that the 
`ZQI22Message` will be passed along to for further processing.

See [main.go](main.go) for an example.