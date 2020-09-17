<div align="center"><img src="ward.gif" alt="Photograph of a woman lying on a diagonal plane."></div>
<div align="center"><small><sup>Fred Ward as Hoke Moseley in <i>Miami Blues (1990)</i></sup></small></div>
<h1 align="center">
  <b>Hoke</b>
</h1>

<h4 align="center">A health check app for k8s init containers.</h4>

<p align="center">
  <a href="#status">Status</a> •
  <a href="#usage">Usage</a> •
  <a href="#contributing">Contributing</a> •
  <a href="#license">License</a>
</p>

<p align="center">
  <a href="https://github.com/liampulles/hoke/releases">
    <img src="https://img.shields.io/github/release/liampulles/hoke.svg" alt="[GitHub release]">
  </a>
  <a href="https://travis-ci.com/liampulles/hoke">
    <img src="https://travis-ci.com/liampulles/hoke.svg?branch=master" alt="[Build Status]">
  </a>
    <img alt="GitHub go.mod Go version" src="https://img.shields.io/github/go-mod/go-version/liampulles/hoke">
  <a href="https://goreportcard.com/report/github.com/liampulles/hoke">
    <img src="https://goreportcard.com/badge/github.com/liampulles/hoke" alt="[Go Report Card]">
  </a>
  <a href="https://codecov.io/gh/liampulles/hoke">
    <img src="https://codecov.io/gh/liampulles/hoke/branch/master/graph/badge.svg" />
  </a>
  <a href="https://microbadger.com/images/lpulles/hoke">
    <img src="https://images.microbadger.com/badges/image/lpulles/hoke.svg">
  </a>
  <a href="https://github.com/liampulles/hoke/blob/master/LICENSE.md">
    <img src="https://img.shields.io/github/license/liampulles/hoke.svg" alt="[License]">
  </a>
</p>

## Status

Hoke is currently in heavy development.

## Usage

Either download a release from the releases page, or clone and run `make install`, and execute:

```bash
hoke https://github.com
```

Or, run the docker image:

```bash
docker run lpulles/hoke https://github.com
```

Or, use as a Kubernetes init container:

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: myapp-pod
  labels:
    app: myapp
spec:
  containers:
  - name: myapp-container
    image: busybox:1.28
    command: ['sh', '-c', 'echo The app is running! && sleep 3600']
  initContainers:
  - name: init-github
    image: lpulles/hoke
    command: ['hoke', 'https://github.com']
```

Run `hoke` to see more options.

## Contributing

Please submit an issue with your proposal.

## License

See [LICENSE](LICENSE)