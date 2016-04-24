# Goblin <img src="https://github.com/ykumards/goblin/blob/master/goblin.png" alt="goblin image" width="10%" height="10%"/>
Goblin paints your github calender green. It modifies your repository's date environmental variable to make commits in the past. 

Goblin's main purpose was to help me learn the basics of running external commands (mostly git) in Go, so it requires [Go](https://github.com/golang/go/wiki/Ubuntu) compiler to function.

## Usage
Make sure you clone the repo using ssh and configure github with your public rsa key (details [here](https://help.github.com/articles/generating-an-ssh-key/)). In case you've cloned the repo using https, [this](https://help.github.com/articles/changing-a-remote-s-url/) might be useful.

For running the script, place goblin.go inside the repo and use:
    <pre><code>go run goblin.go *days*</pre></code>
*days*, is the number of many days you'd like goblin to go back. 

This will take a while.

This has been tested on Ubuntu 14.04 (trusty).
