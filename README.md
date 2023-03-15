# Go Bot

```
> target,       where bot is running
> attackbox,    from where I send commands


# example 1

[target]$ go run .
2023/03/14 19:02:26 Bot is ready to handle requests at port 8080

[attackbox]$ nc <ip> 8080
ping -c2 google.com

[target]$ go run .
2023/03/14 19:02:26 Bot is ready to handle requests at port 8080
PING google.com (216.58.209.14): 56 data bytes
64 bytes from 216.58.209.14: icmp_seq=0 ttl=118 time=12.802 ms
64 bytes from 216.58.209.14: icmp_seq=1 ttl=118 time=11.810 ms

--- google.com ping statistics ---
2 packets transmitted, 2 packets received, 0.0% packet loss
round-trip min/avg/max/stddev = 11.810/12.306/12.802/0.496 ms


# example 2

[attackbox]$ nc <ip> 8080
cat << EOF > /tmp/test.log
test1
test2
EOF
cat /tmp/test.log

[target]$ go run .
2023/03/14 19:02:26 Bot is ready to handle requests at port 8080
test1
test2
```