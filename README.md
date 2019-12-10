# ytt for multi components apps

!!! NOTE !!!

use custom built ytt:

```
git clone https://github.com/k14s/ytt ytt-go/src/github.com/k14s/ytt
export GOPATH=$PWD/ytt-go
cd ytt-go/src/github.com/k14s/ytt
git checkout ytt-library
./hack/build.sh
./ytt version
```

```
$ ./ytt -f config/ -f user-provided/values-with-db.yml
$ ./ytt -f config/ -f user-provided/values-with-exterbal-db.yml
$ ./ytt -f config/_ytt_lib/cc/config -f config/_ytt_lib/cc/user-provided/values.yml
```
