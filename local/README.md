**WORK IN PROGRESS**

Contains a [magefile](https://magefile.org/) to help ease apply kubeflow resources. Currently performs a full install but in future it would be expanded to have targets for component wise installation.

To get the list of all targets run the following command

```
mage
```

If you don't have or don't want to install mage then run the following command(Requires GO)

```
go run mage.go
```

# Targets

To make a target run

```
go run mage.go <target>
```

Working Targets:

- `cluster`: Creates a single node k8s cluster using [kind](https://kind.sigs.k8s.io/)
- `manifests`: Clones the [kubeflow manifests repository](https://github.com/kubeflow/manifests).
- `full`: Applies a full deployment of Kubeflow
- `clean`: Cleans everything by just deleting the kind cluster