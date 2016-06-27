# GO-GETSTREAM

[![godoc](https://godoc.org/github.com/hyperworks/go-getstream?status.svg)](https://godoc.org/github.com/hyperworks/go-getstream)
[![travis build](https://api.travis-ci.org/hyperworks/go-getstream.svg)](https://travis-ci.org/hyperworks/go-getstream)

WIP [getstream.io](getstream.io) client in pure GO.

# WORKING

* Authentication
* Add, list all, and remove an activity.
* Follows/unfollows. 

# YET TO WORK
* multiple adding, follow
* Listing/paging options. 
* Subscription. ??
* Aggregated feeds support. == related to listing, paging 
* Notification feeds support. == get read seen ??

# NICE TO HAVE

* Ability to use pass and retreive a custom structs (w/ embedded `*getstream.Activity`)
  ala `sqlx` style.
* Real-time clients w/ channels.

# LICENSE

MIT

