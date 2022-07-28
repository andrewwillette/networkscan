## Challenge One
Scan local network for all devices.

`ipconfig getsummary en0` returns information about local subnet mask. It is `255.255.255.0`, or `/24` meaning 24 bits of 1. IE `11111111 11111111 11111111 00000000`. A common subnet for residential wifi.
