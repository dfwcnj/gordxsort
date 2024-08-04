// package goranddatagen generates random string, uint64, and datetime data
package main

//package goranddatagen

import (
	"encoding/binary"
	"fmt"
	"math/rand"
	"time"
)

var alphanum = []rune("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")

// randSeq(n int)
// generate a random string length n with lower, upper case letters and digits
func randSeq(n int, rlen bool) string {
	b := make([]rune, n)
	ll := len(alphanum)
	for i := range b {
		b[i] = alphanum[rand.Intn(ll)]
	}
	if rlen == true {
		rl := rand.Intn(n)
		if rl != 0 {
			b = b[:rl]
		}
	}
	return string(b)
}

// randomstrings(n int, slen int)
// generate n random strings with length slen
// return a slice containing the strings
func randomstrings(n int, slen int, rlen bool) []string {
	ssl := make([]string, 0)
	for _ = range n {
		ssl = append(ssl, randSeq(slen, rlen))
	}
	return ssl
}

// randomints(ņ int)
// generate n random int64 values
// return a slice containing the int64 values
func randomuintb(n int) lines {
	rubsl := make(lines, 0)
	for _ = range n {
		//fmt.Println(rand.Uint64())
		ru := rand.Uint64()
		rub := make(line, 8)
		binary.LittleEndian.PutUint64(rub, ru)
		rubsl = append(rubsl, rub)
	}
	return rubsl
}

// randomints(ņ int)
// generate n random int64 values
// return a slice containing the int64 values
func randomuints(n int) []uint64 {
	usl := make([]uint64, 0)
	for _ = range n {
		//fmt.Println(rand.Uint64())
		usl = append(usl, rand.Uint64())
	}
	return usl
}

// randomdates(n int, format string)
// generate n random dates with format
// return a slice containing the random date strings
func randomdates(n int, format string) []string {
	now := time.Now().Unix()
	var mod = int64(now)
	var s string
	dsl := make([]string, 0)
	for _ = range n {
		ri := rand.Int63() % mod
		tm := time.Unix(int64(ri), int64(0))

		switch format {
		case "DateTime":
			s = fmt.Sprint(tm.Format(time.DateTime))
		case "Layout":
			s = fmt.Sprint(tm.Format(time.Layout))
		case "RubyDate":
			s = fmt.Sprint(tm.Format(time.RubyDate))
		case "UnixDate":
			s = fmt.Sprint(tm.Format(time.UnixDate))
		case "RFC3339":
			s = fmt.Sprint(tm.Format(time.RFC3339))
		default:
			s = fmt.Sprint(tm)
		}
		dsl = append(dsl, s)
	}
	return dsl
}
