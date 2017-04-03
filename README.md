Prometheus mDNS service discovery
=================================

(Under development - doesn't quite work yet)

Build:

    go build ./... && ./prometheus-mdns-sd

It listens for mDNS/ZeroConf/Bonjour service announcements under
`_prometheus-http._tcp` and `_prometheus-https._tcp`, captures port-number and
hostname for later re-labeling.

Manually create service announcement (OS X):

    dns-sd -R "My test server with metrics-endpoint" _prometheus-http._tcp. . 9000 path=/metrics
