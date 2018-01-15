
# Installation

*rss2* is a collection of command line programs run from a shell like Bash. 

## Compiled version

This is generalized instructions for a release. 

Compiled versions are available for Mac OS X (amd64 processor, macosx-amd64), Linux (amd64 process, linux-amd64), 
Windows (amd64 processor, windows-amd64) and Rapsberry Pi (arm7 processor, raspbian-arm7)

VERSION_NUMBER is a [symantic version number](http://semver.org/) (e.g. v0.1.2)


For all the released version go to the project page on Github and click latest release

>    https://github.com/caltechlibrary/rss2/releases/latest


| Platform    | Zip Filename                               |
|-------------|--------------------------------------------|
| Windows     | rss2-VERSION_NUMBER-windows-amd64.zip |
| Mac OS X    | rss2-VERSION_NUMBER-macosx-amd64.zip  |
| Linux/Intel | rss2-VERSION_NUMBER-linux-amd64.zip   |
| Raspbery Pi | rss2-VERSION_NUMBER-raspbian-arm7.zip |


## The basic recipe

+ Find the Zip file listed matching the architecture you're running and download it
    + (e.g. if you're on a Windows 10 laptop/Surface with a amd64 style CPU you'd choose the Zip file with "windows-amd64" in the name).
+ Download the zip file and unzip the file.  
+ Copy the contents of the folder named "bin" to a folder that is in your path 
    + (e.g. "$HOME/bin" is common).
+ Adjust your PATH if needed
    + (e.g. `export PATH="$HOME/bin:$PATH"`)
+ Test


### Mac OS X

1. Download the zip file
2. Unzip the zip file
3. Copy the executables to $HOME/bin (or a folder in your path)
4. Make sure the new location in in our path
5. Test

Here's an example of the commands run in the Terminal App after downloading the 
zip file.

```shell
    cd Downloads/
    unzip rss2-*-macosx-amd64.zip
    mkdir -p $HOME/bin
    cp -v bin/* $HOME/bin/
    export PATH=$HOME/bin:$PATH
    csvfind -version
```

### Windows

1. Download the zip file
2. Unzip the zip file
3. Copy the executables to $HOME/bin (or a folder in your path)
4. Test

Here's an example of the commands run in from the Bash shell on Windows 10 after
downloading the zip file.

```shell
    cd Downloads/
    unzip rss2-*-windows-amd64.zip
    mkdir -p $HOME/bin
    cp -v bin/* $HOME/bin/
    export PATH=$HOME/bin:$PATH
    csvfind -version
```


### Linux 

1. Download the zip file
2. Unzip the zip file
3. Copy the executables to $HOME/bin (or a folder in your path)
4. Test

Here's an example of the commands run in from the Bash shell after
downloading the zip file.

```shell
    cd Downloads/
    unzip rss2-*-linux-amd64.zip
    mkdir -p $HOME/bin
    cp -v bin/* $HOME/bin/
    export PATH=$HOME/bin:$PATH
    csvfind -version
```


### Raspberry Pi

Released version is for a Raspberry Pi 2 or later use (i.e. requires ARM 7 support).

1. Download the zip file
2. Unzip the zip file
3. Copy the executables to $HOME/bin (or a folder in your path)
4. Test

Here's an example of the commands run in from the Bash shell after
downloading the zip file.

```shell
    cd Downloads/
    unzip rss2-*-raspbian-arm7.zip
    mkdir -p $HOME/bin
    cp -v bin/* $HOME/bin/
    export PATH=$HOME/bin:$PATH
    csvfind -version
```


## Compiling from source

_rss2_ is "go gettable".  Use the "go get" command to download the dependant packages
as well as _rss2_'s source code.

```shell
    go get -u github.com/caltechlibrary/rss2/...
```

Or clone the repstory and then compile

```shell
    cd
    git clone https://github.com/caltechlibrary/rss2 src/github.com/caltechlibrary/rss2
    cd src/github.com/caltechlibrary/rss2
    make
    make test
    make install
```


