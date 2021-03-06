Inspiration
------------

After reading https://bennycwong.github.io/post/speeding-up-oh-my-zsh/ I was inspired to look at my own zsh for any python scripts that I could replace to improve the latency of my shell.

Purpose
--------

This program will be a binary that can be used to replace the gitstatus.py script in the git-prompt plugin of oh-my-zsh.

Status
-------

The gitstatus program produced does not yet correctly set all the variables necessary for it to be a replacement, nor is it faster than the python version of the script in some cases (yet). 

Installation
-------------

    go get
    go install
    which gitstatus

Then modify the ~/.oh-my-zsh/plugins/git-prompt/git-prompt.plugin.zsh file and comment out 

    local gitstatus="$__GIT_PROMPT_DIR/gitstatus.py"
    _GIT_STATUS=$(python ${gitstatus} 2>/dev/null)
    
and add in 

    _GIT_STATUS=$(/path/to/gitstatus)
    
obviously, replacing it with the path to your compiled and installed version of the gitstatus binary
