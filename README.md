Prometheus mDNS service discovery
=================================

Build:

    go build ./... && ./prometheus-mdns-sd

It listens for mDNS/ZeroConf/Bonjour service announcements under
`_prometheus-http._tcp` and `_prometheus-https._tcp`, captures port-number and
hostname for later re-labeling.

Manually create service announcement (OS X):

    dns-sd -R "My test server with metrics-endpoint" _prometheus-http._tcp. . 9000 path=/metrics

Bugs/todo at time of writing:

 - Doesn't actually write out a file for Prometheus to consume
 - Doesn't parse the TXT-segment
 - Should merge `_prometheus-http._tcp` and `_prometheus-https._tcp` segments properly.
 - Doesn't abort right away if context is canceled.
