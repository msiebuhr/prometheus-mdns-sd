Prometheus mDNS service discovery
=================================


Discovers mDNS (a.k.a. ZeroConf, a.k.a. Bonjour) service announcements under
`_prometheus-http._tcp` and `_prometheus-https._tcp` for ad-hoc discovery of
devices on LAN networks.

Install & running

    go install github.com/msiebuhr/prometheus-mdns-sd

Run it

    prometheus-mdns-sd -out /etc/prometheus/mdns-sd.json

And in `prometheus.yml` something along these lines:

    - job_name: mdns-sd
      scrape_interval: 30s
      scrape_timeout: 10s
      metrics_path: /metrics
      scheme: http
      file_sd_configs:
      - files:
        - /etc/prometheus/mdns-sd.json
        refresh_interval: 5m

It resolves the raw IP's (the Go DNS resolver doesn't always understand
RFC6762/RFC6763's `.local` names) and captures the port-number and hostname for
later re-labeling.

## Clients

Manually create service announcement (OS X):

    dns-sd -R "My test server with metrics-endpoint" _prometheus-http._tcp. . 9000 path=/metrics

And there's some code for
[arduino/esp8266](https://github.com/msiebuhr/esp8266-promstuff/blob/e7e6a353db1b74450efbfb711f51077ba0df06be/esp8266-thermometer-ds18b20.ino#L60-L65).


## Related reading

 * Initial PR to get mDNS SD into Prometheus: https://github.com/prometheus/prometheus/pull/1903
