# Envoyproxybeat

Envoyproxybeat is an [Elastic Beat](https://www.elastic.co/products/beats) that collects metrics from Envoy Proxy and indexes them into Elasticsearch or Logstash.

## Description

> [Envoy Proxy](https://www.envoyproxy.io/docs/envoy/latest/intro/what_is_envoy) is an L7 proxy and communication bus designed for large modern service-oriented architectures.

## Configuration

Adjust the `envoyproxybeat.yml` configuration file to your needs.

### `period`
Defines how often to take traceroute output. Default to `30` s.

### `port`
Defines the envoy proxy port serviced. Default to `:9901`

### `host`
Host name of ElasticSearch. Default to `localhost` 

## Document Example

<pre>

  "server": [
    "filesystem": {
      "reopen_failed": 0,
      "write_buffered": 24,
      "write_completed": 24,
      "write_total_buffered": 0,
      "flushed_by_timer": 447
    },
    "runtime": {
      "num_keys": 0,
      "override_dir_exists": 0,
      "override_dir_not_exists": 0,
      "admin_overrides_active": 0,
      "load_error": 0,
      "load_success": 0
    },
    "listener_manager": {
      "listener_added": 1,
      "listener_create_failure": 0,
      "listener_create_success": 4,
      "listener_modified": 0,
      "listener_removed": 0,
      "total_listeners_active": 1,
      "total_listeners_draining": 0,
      "total_listeners_warming": 0
    },
    "stats": {
      "overflow": 0
    },
    "server": {
      "hot_restart_epoch": 0,
      "memory_allocated": 3168120,
      "version": 4151803,
      "live": 1,
      "watchdog_miss": 0,
      "days_until_first_cert_expiring": 2147483647,
      "parent_connections": 0,
      "watchdog_mega_miss": 0,
      "memory_heap_size": 4194304,
      "total_connections": 0,
      "uptime": 4506
    },
    "http2": {},
    "cluster_manager": {
      "active_clusters": 1,
      "cluster_added": 1,
      "cluster_modified": 0,
      "cluster_removed": 0,
      "warming_clusters": 0
    }
  ],
  "type": "envoyproxybeat"
}
</pre>

## Getting Started with Envoyproxybeat

Ensure that this folder is at the following location:
`${GOPATH}/src/github.com/berfinsari/envoyproxybeat`

### Requirements

* [Golang](https://golang.org/dl/) 1.11

### Init Project
To get running with Envoyproxybeat and also install the
dependencies, run the following command:

```
make setup
```

It will create a clean git history for each major step. Note that you can always rewrite the history if you wish before pushing your changes.

To push Envoyproxybeat in the git repository, run the following commands:

```
git remote set-url origin https://github.com/berfinsari/envoyproxybeat
git push origin master
```

For further development, check out the [beat developer guide](https://www.elastic.co/guide/en/beats/libbeat/current/new-beat.html).

### Build

To build the binary for Envoyproxybeat run the command below. This will generate a binary
in the same directory with the name envoyproxybeat.

```
make
```


### Run

To run Envoyproxybeat with debugging output enabled, run:

```
./envoyproxybeat -c envoyproxybeat.yml -e -d "*"
```


### Test

To test Envoyproxybeat, run the following command:

```
make testsuite
```

alternatively:
```
make unit-tests
make system-tests
make integration-tests
make coverage-report
```

The test coverage is reported in the folder `./build/coverage/`

### Update

Each beat has a template for the mapping in elasticsearch and a documentation for the fields
which is automatically generated based on `fields.yml` by running the following command.

```
make update
```


### Cleanup

To clean  Envoyproxybeat source code, run the following commands:

```
make fmt
make simplify
```

To clean up the build directory and generated artifacts, run:

```
make clean
```


### Clone

To clone Envoyproxybeat from the git repository, run the following commands:

```
mkdir -p ${GOPATH}/src/github.com/berfinsari/envoyproxybeat
git clone https://github.com/berfinsari/envoyproxybeat ${GOPATH}/src/github.com/berfinsari/envoyproxybeat
```


For further development, check out the [beat developer guide](https://www.elastic.co/guide/en/beats/libbeat/current/new-beat.html).


## Packaging

The beat frameworks provides tools to crosscompile and package your beat for different platforms. This requires [docker](https://www.docker.com/) and vendoring as described above. To build packages of your beat, run the following command:

```
make release
```

This will fetch and create all images required for the build process. The whole process to finish can take several minutes.

## License
Covered under the Apache License, Version 2.0
Copyright (c) 2018 Berfin Sarı
