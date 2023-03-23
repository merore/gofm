# gofm
Gofm is a robot for missevan that supports chat powered by ChatGPT.

## Features
- [x] Greet
- [x] Chat with ChatGPT

## Install
### compile from source
```
make -j`nproc`
```

## Usage
### run gofm
```
./bin/gofm --config=config.yaml
```
### Chat with robot
```
// chat
@robot 给我讲个笑话

// reset chat context
@robot reset
```
### Options
### get missevan token
```
./bin/gofm token PHONE PASSWORD
```
Replace `PHONE` and `PASSWORD` with your priviate account.
### help
You can run `gofm -h` to get more detailed information.

## Contributing
PRs accepted.