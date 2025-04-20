# healthchecks

This is a simple app to test Kubernetes container health check configurations.

Force the liveness probe to fail:

```bash
oc rsh $(oc get pod -o name) rm -f /tmp/healthy
```
