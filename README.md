# ytt for multi components apps

!!! NOTE !!!

Must use ytt 0.23.0+

```
$ ytt -f config/ -f user-provided/values-with-db.yml
$ ytt -f config/ -f user-provided/values-with-exterbal-db.yml
$ ytt -f config/_ytt_lib/cc/config -f config/_ytt_lib/cc/user-provided/values.yml
```
