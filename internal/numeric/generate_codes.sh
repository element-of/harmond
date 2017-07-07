#!/bin/sh

echo 'package numeric

// The list of all numerics the ircd uses.
//
// This is scraped from the ircd source code directly.

type Response string

const (
' > replies.go

cat ./messages.tab | grep -v NULL | grep '"' | cut -d'/' -f2 | cut -d'*' -f2 | awk '{ printf $2" Response = \""$1"\"\n" }' | sed 's/,//' >> modes_temp.go

cat modes_temp.go | sed -r 's/([A-Z]+)_([A-Z]+)/\L\1\u\2/' | sed -e 's/\b\(.\)/\u\1/g' | grep -v LAST > modes_temp2.go
rm modes_temp.go

moon generate_check.moon

cat modes_temp2.go >> replies.go
rm modes_temp2.go

echo ')' >> replies.go

cat ./messages.tab | grep -v NULL | grep '"'  | go run generate_strtable.go

gofmt -w .
