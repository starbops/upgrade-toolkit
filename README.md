# Harvester Upgrade Toolkit

Upgrade Toolkit is the central part of the Harvester upgrade solution. The entire Upgrade V2 enhancement includes

- Upgrade Shim: embedded in the Harvester controller manager
- Upgrade Toolkit
  - Upgrade Repo
  - Upgrade Manager
  - Helper scripts

## How to Initiate an Upgrade

1. When the time comes, a yellow "Upgrade" button will show up on the top right corner of the Harvester UI dashboard
1. Click the "Upgrade" button
1. Tweak upgrade specific options on the pop-up upgrade dialog and then click the "Upgrade" button to start upgrading the cluster
1. Check the upgrade progress on the upgrade modal by clicking the green circle on the top right corner of the Harvester UI dashboard
1. Upgrade finishes

## Overall Workflow

1. The user creates a Version CR
1. The user creates the corresponding UpgradePlan CR or click the "Upgrade" button on the Harvester UI dashboard
1. Upgrade Shim downloads the ISO image from remote
1. Upgrade Shim deploys Upgrade Toolkit which includes Upgrade Repo and Manager
1. Upgrade Repo downloads the ISO image from internal
1. Upgrade Repo preloads all the container images
1. Upgrade Repo transitions to ready
1. Upgrade Manager transitions to ready
1. Upgrade Manager upgrades cluster components
1. Upgrade Manager upgrades node components
1. Upgrade Manager cleans up resources
1. Upgrade Manager marks the Upgrade CR as complete
1. Upgrade Shim tears down Upgrade Toolkit

## Customized Upgrades

Upgrade Toolkit supports upgrading a Harvester cluster using other container images for Upgrade Repo and Manager that are not packaged in the ISO image. To do so, please see below.

Create a Version CR. This is the same as before.

```bash
cat <<EOF | kubectl create -f -
apiVersion: management.harvesterhci.io/v1beta1
kind: Version
metadata:
  namespace: harvester-system
  name: v1.6.0
spec:
  isoURL: http://10.115.49.5/iso/harvester/v1.6.0/harvester-v1.6.0-amd64.iso
EOF
```

When creating the UpgradePlan CR, specifying a different container image to use for Upgrade Repo and Manager:

```bash
cat <<EOF | kubectl create -f -
apiVersion: management.harvesterhci.io/v1beta1
kind: UpgradePlan
metadata:
  namespace: harvester-system
  generateName: hvst-upgrade-
spec:
  version: v1.6.0
  upgrade: dev
EOF
```

A successfully executed UpgradePlan looks like the following:

```yaml
apiVersion: management.harvesterhci.io/v1beta1
kind: UpgradePlan
metadata:
  creationTimestamp: "2025-09-23T13:41:31Z"
  generateName: hvst-upgrade-
  generation: 1
  name: hvst-upgrade-jq9n5
  resourceVersion: "1656954"
  uid: 1b0200e0-ecd4-4768-8c7c-f19814fbcefa
spec:
  version: v1.6.0
status:
  conditions:
  - lastTransitionTime: "2025-09-23T13:44:44Z"
    message: ""
    observedGeneration: 1
    reason: Executed
    status: "False"
    type: Available
  - lastTransitionTime: "2025-09-23T13:44:44Z"
    message: UpgradePlan has completed
    observedGeneration: 1
    reason: CleanedUp
    status: "False"
    type: Progressing
  - lastTransitionTime: "2025-09-23T13:44:44Z"
    message: ""
    observedGeneration: 1
    reason: ReconcileSuccess
    status: "False"
    type: Degraded
  phase: Succeeded
  phaseTransitionTimestamps:
  - phase: Init
    phaseTransitionTimestamp: "2025-09-23T13:41:31Z"
  - phase: ISODownloading
    phaseTransitionTimestamp: "2025-09-23T13:41:31Z"
  - phase: ISODownloaded
    phaseTransitionTimestamp: "2025-09-23T13:41:31Z"
  - phase: RepoCreating
    phaseTransitionTimestamp: "2025-09-23T13:41:31Z"
  - phase: RepoCreated
    phaseTransitionTimestamp: "2025-09-23T13:41:31Z"
  - phase: MetadataPopulated
    phaseTransitionTimestamp: "2025-09-23T13:41:32Z"
  - phase: ImagePreloading
    phaseTransitionTimestamp: "2025-09-23T13:41:32Z"
  - phase: ImagePreloaded
    phaseTransitionTimestamp: "2025-09-23T13:41:32Z"
  - phase: ClusterUpgrading
    phaseTransitionTimestamp: "2025-09-23T13:41:33Z"
  - phase: ClusterUpgraded
    phaseTransitionTimestamp: "2025-09-23T13:42:06Z"
  - phase: NodeUpgrading
    phaseTransitionTimestamp: "2025-09-23T13:42:06Z"
  - phase: NodeUpgraded
    phaseTransitionTimestamp: "2025-09-23T13:44:43Z"
  - phase: CleaningUp
    phaseTransitionTimestamp: "2025-09-23T13:44:43Z"
  - phase: CleanedUp
    phaseTransitionTimestamp: "2025-09-23T13:44:43Z"
  - phase: Succeeded
    phaseTransitionTimestamp: "2025-09-23T13:44:44Z"
  releaseMetadata:
    harvester: v1.6.0
    harvesterChart: 1.6.0
    kubernetes: v1.33.3+rke2r1
    minUpgradableVersion: v1.5.0
    monitoringChart: 105.1.2+up61.3.2
    os: Harvester v1.6.0
    rancher: v2.12.0
```
