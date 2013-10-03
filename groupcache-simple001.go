// see https://code.google.com/p/go/source/browse/2013/oscon-dl/groupcache.go?spec=svn.talks.4b257899b8d1f1e4df60abe47ec540b1d294333d&repo=talks&r=4b257899b8d1f1e4df60abe47ec540b1d294333d
package main

import (
    "github.com/golang/groupcache"
)

func main() {
    me := "http://10.0.0.1"
    peers := groupcache.NewHTTPPool(me)
}
