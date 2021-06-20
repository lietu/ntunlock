# Nuclear Throne unlocker

Cross-platform (Windows, Linux, Mac OS X) Nuclear Throne save-game editor with the options to unlock all characters, all crowns on all characters, all B-skins, and hard mode.

Why? Well, maybe you deleted your save, maybe the Steam Cloud integration deleted your save, or maybe you just want to see what Hard Mode is like with Crown of Blood.


## Running it (for normal people)

Download the [release](https://github.com/lietu/ntunlock/releases) for your platform, extract the .ZIP and run the `ntunlocker` -file in it. It will look something like this:

![NTUnlock Screenshot](screenshot.jpg?raw=true)


## Building it (for nerds)

You'll need Go installed. Get it from [golang.org](https://golang.org). I built this on 1.5.x, that might be required.

On Windows:
```
go get github.com/lietu/ntunlock/ntunlock
cd %GOPATH%\src\github.com\lietu\ntunlock\cmd\ntunlock
go build
```

On Linux/MacOSX:
```
go get github.com/lietu/ntunlock/ntunlock
cd $GOPATH/src/github.com/lietu/ntunlock/cmd/ntunlock
go build
```


## Licensing

Short answer: MIT and New BSD

Long answer: Read LICENSE.md


# Financial support

This project has been made possible thanks to [Cocreators](https://cocreators.ee) and [Lietu](https://lietu.net). You can help us continue our open source work by supporting us on [Buy me a coffee](https://www.buymeacoffee.com/cocreators).

[!["Buy Me A Coffee"](https://www.buymeacoffee.com/assets/img/custom_images/orange_img.png)](https://www.buymeacoffee.com/cocreators)
