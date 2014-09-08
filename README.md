# Forninet Go Client
The idea is to proxy the syslog messages received by a fortinet to
different computers so they can analyze the traffic.

## TODO

- Analyze the traffic on each client node
- OpenGL charts

# Questions

On multiplexing messages:

- 1 channel per forward address with one always opened connection per
  address vs one connection per message one connection per message per
	forward address ???
