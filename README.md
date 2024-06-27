
# Service Dependencies Operator

## Overview
The Service Dependencies Operator is a Kubernetes operator designed to ensure that all specified dependencies are present and correctly set up before marking the deployment pods of a service as ready. This operator uses a Custom Resource Definition (CRD) to manage service dependencies across your Kubernetes cluster.

## Prerequisites
- Kubernetes 1.16+
- Kubectl installed and configured
- Operator Lifecycle Manager (OLM) for deploying the operator (optional)

## Installation

### Install the CRD
First, you need to install the Custom Resource Definition (CRD) in your Kubernetes cluster:

```bash
kubectl apply -f crd.yaml
```

### Deploy the Operator
To deploy the Service Dependencies Operator in your cluster, apply the deployment YAML:

```bash
kubectl apply -f operator-deployment.yaml
```

## Usage

### Creating a ServiceDependency
To create a ServiceDependency, you need to define it in a YAML file. Below is an example:

```yaml
apiVersion: io.com/v1
kind: ServiceDependency
metadata:
  name: example-servicedependency
spec:
  service:
    deploymentName: my-service
    namespace: default
  dependencies:
    - kind: Deployment
      name: dependency1
      namespace: default
    - kind: ConfigMap
      name: config1
      namespace: default
```

Apply this configuration using kubectl:

```bash
kubectl apply -f example-servicedependency.yaml
```

## Viewing the Status of Dependencies

Check the status of your service dependencies:

```bash
kubectl get sd example-servicedependency -o yaml
```

This will display the current status of the dependencies, indicating whether they are ready and any additional messages.

## Contributing

Contributions are welcome! Please read our contributing guidelines located in `CONTRIBUTING.md` to learn about our development process, how to propose bugfixes and improvements, and how to build and test your changes to the project.


## Support

If you have any issues or need help, please submit an issue to the project's issue tracker.
