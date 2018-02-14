
## Minikube Environment Variables
Minikube supports passing environment variables instead of flags for every value listed in `minikube config list`.  This is done by passing an environment variable with the prefix `MINIKUBE_`. For example the `minikube start --iso-url="$ISO_URL"` flag can also be set by setting the `MINIKUBE_ISO_URL="$ISO_URL"` environment variable.

Some features can only be accessed by environment variables, here is a list of these features:

* **MINIKUBE_HOME** - (string) sets the path for the .minikube directory that minikube uses for state/configuration

* **MINIKUBE_WANTUPDATENOTIFICATION** - (bool) sets whether the user wants an update notification for new minikube versions

* **MINIKUBE_REMINDERWAITPERIODINHOURS** - (int) sets the number of hours to check for an update notification
* **MINIKUBE_WANTREPORTERROR** - (bool) sets whether the user wants to send anonymous errors reports to help improve minikube

* **MINIKUBE_WANTREPORTERRORPROMPT** - (bool) sets whether the user wants to be prompted on an error that they can report them to help improve minikube

* **MINIKUBE_WANTKUBECTLDOWNLOADMSG** - (bool) sets whether minikube should tell a user that `kubectl` cannot be found on there path

* **MINIKUBE_ENABLE_PROFILING** - (int, `1` enables it) enables trace profiling to be generated for minikube which can be analyzed via:

```shell
# set env var and then run minikube
$ MINIKUBE_ENABLE_PROFILING=1 ./out/minikube start
2017/01/09 13:18:00 profile: cpu profiling enabled, /tmp/profile933201292/cpu.pprof
Starting local Kubernetes cluster...
Kubectl is now configured to use the cluster.
2017/01/09 13:19:06 profile: cpu profiling disabled, /tmp/profile933201292/cpu.pprof

# Then you can examine the profile with:
$ go tool pprof  /tmp/profile933201292/cpu.pprof
```
