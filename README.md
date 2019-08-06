# rosgo

## Package Summary

**rosgo** is pure Go implementation of [ROS](http://www.ros.org/) client library.

- Author: Akio Ochiai
- License: Apache License 2.0
- Source: git [https://github.com/akio/rosgo](https://github.com/akio/rosgo)

## Status

**rosgo** is under development to implement all features of [ROS Client Library Requiements](http://www.ros.org/wiki/Implementing%20Client%20Libraries).

At present, following basic functions are provided.

- Parameter API (get/set/search....)
- ROS Slave API (with some exceptions)
- Publisher/Subscriber API (with TCPROS)
- Remapping
- Message Generation

## Fork

To use this fork, add the following to your `go.mod` file (you should use Go modules):

```
replace github.com/akio/rosgo => github.com/gelin/rosgo v0.0.4

require github.com/gelin/rosgo v0.0.4
```

## See also

- [rosgo in ROS Wiki](http://www.ros.org/wiki/rosgo)
