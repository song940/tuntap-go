# tuntap.go

## Setup

Linux

```sh
sudo ip addr add 10.0.0.1/24 dev utun9
sudo ip link set dev utun9 up
```

macOS

```sh
sudo ifconfig utun9 10.0.0.1 10.0.0.255 up
```

Sending some ICMP broadcast message:

```sh
ping 10.0.0.255
```

You'll see output containing the IPv4 ICMP frame:
```
2023/02/14 16:14:01 Ethertype: 0a 00
2023/02/14 16:14:01 Src: 00:00:40:01:ce:c4
2023/02/14 16:14:01 Dst: 45:00:00:54:96:e5
2023/02/14 16:14:01 Payload: 00 01 0a 00 00 ff 08 00 7a ae be 7f 00 0d 63 eb 42 c9 00 08 2d 05 08 09 0a 0b 0c 0d 0e 0f 10 11 12 13 14 15 16 17 18 19 1a 1b 1c 1d 1e 1f 20 21 22 23 24 25 26 27 28 29 2a 2b 2c 2d 2e 2f 30 31 32 33 34 35 36 37
```