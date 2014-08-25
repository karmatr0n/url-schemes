Small web application to take advantage of vulnerable URL Schemes in iOS 
========================================================================

It is not a 0day in the traditional meaning, it's a simple _did not read the manual_ and _bad defaults_ case from Apple.

Thanks to Guillaume K. Ross by his research and his presentation at BSidesLV 2014.

Dependencies
============

1. [Go Lang](http://golang.org/dl/)
2. [Traffic web framework](https://github.com/pilu/traffic)

Installation
============

1. git clone https://github.com/juarlex/url-schemes.git
2. Edit the traffic.conf file and set your settings
3. go clean; go build; ./url-schemes
4. Open the page in your iOS device or your web browser: http://127.0.0.1:3000/

References
==========

1. [iOS URL Schemes: omg://, Guillaume K. Ross, BSidesLV, 2004](http://www.irongeek.com/i.php?page=videos/bsideslasvegas2014/pg10-ios-url-schemes-omg-guillaume-k-ross)
2. [RTFM 0day in iOS apps: G+, Gmail, FB Messenger, etc.](http://algorithm.dk/posts/rtfm-0day-in-ios-apps-g-gmail-fb-messenger-etc)
