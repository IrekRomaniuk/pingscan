package main

import (
	"testing"

    . "github.com/IrekRomaniuk/pingscan/targets"
    . "github.com/smartystreets/goconvey/convey"
)

func TestReadTargets(t *testing.T) {
    Convey("Read pinglist.txt from examples ", t, func() {
        target := "./examples/pinglist.txt"
        hosts, _ := ReadTargets(target)
        Convey("So pinglist.txt should contain 3 items", func() {
            So(len(hosts), ShouldEqual,3)
        })
    })
}

func TestPing(t *testing.T) {
    timeout := 5

	r := *ping(&timeout, &[]string{"127.0.0.1"})
    if r[0].RTT() <= 0 || r[0].Err != nil {
        t.Error("Ping localhost failed")
    }

    r = *ping(&timeout, &[]string{"google.com"})
    if r[0].RTT() <= 0 || r[0].Err != nil {
        t.Error("Ping google.com failed")
    }

    r = *ping(&timeout, &[]string{"yahoo.com"})
    if r[0].RTT() <= 0 || r[0].Err != nil {
        t.Error("Ping yahoo.com failed")
    }

    r = *ping(&timeout, &[]string{"owjdfiojsfdjfsoijeifojweoifjdsiojciosdjc.czs"})
    if r[0].Err == nil {
        t.Error("Ping nonsense not failed")
    }

    r = *ping(&timeout, &[]string{"266.266.266.266"})
    if r[0].Err == nil {
        t.Error("Ping nonsense not failed")
    }
}
