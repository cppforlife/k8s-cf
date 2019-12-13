# ytt for multi components apps

!!! NOTE !!!

Must use ytt 0.23.0+

```
$ ytt -f config/ -f user-provided/values-with-db.yml
$ ytt -f config/ -f user-provided/values-with-exterbal-db.yml
$ ytt -f config/_ytt_lib/cc/config -f config/_ytt_lib/cc/user-provided/values.yml
```

Mock CF consists of:

- metacontroller (config/_ytt_lib/github.com/GoogleCloudPlatform/metacontroller)
  - contains unmodified upstream installation YAML files
- cf-k8s-networking (config/_ytt_lib/github.com/cloudfoundry/cf-k8s-networking)
  - contains unmodified upstream ytt templates
  - configured with cc and uaa urls
- cc (config/_ytt_lib/github.com/cloudfoundry/cc)
  - dummy app that needs info about uaa and db
- uaa (config/_ytt_lib/github.com/cloudfoundry/uaa)
  - dummy app that needs info about db
  - configured with client secrets for cc and cf-k8s-networking
- db (config/_ytt_lib/github.com/cloudfoundry/db)
  - configured with db creds and db names for cc and uaa

Idea is that this repository vendors in particular component version. Each components has its own data values files that specifies inputs.

Components are assembled together in config/all.yml, and config/metacontroller.yml.

There is a top level "cf" data values (config/values.yml) that ask for data that's necessary to install entire system.

To see what this mock cf system will install, run following:

```
ytt -f config/ -f user-provided/values-with-db.yml -f ./hack/ignore-kbld.yml | kapp deploy -a test-cf -f- -c --diff-run
```
