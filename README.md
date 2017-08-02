go-j2s
======
A tool for generating Go source code from json data file.

It saves your time for writing nested corresponding struct/slice according to in-coming json data. The generated struct/slice is also has field tag added. It also can work on multiple json files at once.

Contents:
---------
1. [Installation](#installation)
2. [Usage](#usage)
3. [Go Code](#go-code)
4. [Scalability](#scalability)
5. [Go Generate](#go-generate)

Installation
------------

    $ go get -u github.com/liangrog/go-j2s/...    

Uage
----
To specify the base directory to search for all the json files, use `-in` option. It defaults to current directory.

    $ go-j2s -in /tmp

By default, go-j2s only searchs default or the directory specified in `-in`. To search recursively, use `-r`.

    $ go-j2s -in /tmp -r 

For excluding directory from the search, use `-excl`. go-j2s will check if searching directory contains the exluding directory string. 

    $ go-j2s -excl /tmp/test       

You can specify multiple.

    $ go-j2s -excl /tmp/test -excl /tmp/data
        
To specify a source json file to convert, use `-from` option. It will make `-in` option obsolete.

    $ go-j2s -from /home/user/data.json

You can also have multiple source json files.

    $ go-j2s -from /home/user/data.json -from /tmp/test.json

By default, the generated Go file will be saved to the current directory. To change it, use `-out`.

    $ go-j2s -out /home/user/project

`j2s` is the default generated Go file name. You can change it by using `-name`.

    $ go-j2s -name model

The generated Go code uses `main` as package. You can specify the package by using `-pkg`.

    $ go-j2s -pkg config

Go Code
-------
The generated Go code follows the Go "encoding/json" [unmarshal rules](https://golang.org/pkg/encoding/json/). Those are:

    bool, for JSON booleans
    float64, for JSON numbers
    string, for JSON strings
    []interface{}, for JSON arrays
    map[string]interface{}, for JSON objects
    nil for JSON null


Scalability
-----------
go-j2s unmarshals all found json file concurrently, providing quick conversion time.

Go Generate
-----------
You can also use go-j2s in conjuntion with `go generate`. Drop below line in any of your existing Go file for example:

    //go:generate go-j2s -in /tmp/ -r 

Then simply run:

    $ go generate

