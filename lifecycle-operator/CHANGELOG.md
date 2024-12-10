# Changelog

## [3.0.0](https://github.com/keptn/lifecycle-toolkit/compare/lifecycle-operator-v2.0.0...lifecycle-operator-v3.0.0) (2024-12-10)


### ⚠ BREAKING CHANGES

* **lifecycle-operator:** The Lifecycle Operator helm chart was adapted after removal of the Keptn Scheduler and many Helm values were simplified, please double check your values files and adapt them accordingly.

### Bug Fixes

* **lifecycle-operator:** remove scheduler from helm charts ([#3855](https://github.com/keptn/lifecycle-toolkit/issues/3855)) ([fd78a53](https://github.com/keptn/lifecycle-toolkit/commit/fd78a536c1131ca57b0e8e7929a6c382b34e47b8))


### Other

* bump helm chart versions ([#3857](https://github.com/keptn/lifecycle-toolkit/issues/3857)) ([1873178](https://github.com/keptn/lifecycle-toolkit/commit/1873178a28878c0a12ce00e20a8e62d105068fe5))

## [2.0.0](https://github.com/keptn/lifecycle-toolkit/compare/lifecycle-operator-v1.2.0...lifecycle-operator-v2.0.0) (2024-11-11)


### ⚠ BREAKING CHANGES

* The Keptn Scheduler was removed and therefore support for Kubernetes 1.26 and lower was dropped.
* remove Keptn scheduler ([#3821](https://github.com/keptn/lifecycle-toolkit/issues/3821))

### Other

* polish helm charts ([#3853](https://github.com/keptn/lifecycle-toolkit/issues/3853)) ([17fa47b](https://github.com/keptn/lifecycle-toolkit/commit/17fa47b16fb43bb627a466e999851bc1f225e878))
* remove Keptn scheduler ([#3821](https://github.com/keptn/lifecycle-toolkit/issues/3821)) ([de3a0e7](https://github.com/keptn/lifecycle-toolkit/commit/de3a0e7a17c8c16ab350188d5235ca6b435c6ba5))


### Docs

* remove Keptn scheduler ([#3826](https://github.com/keptn/lifecycle-toolkit/issues/3826)) ([cb01b09](https://github.com/keptn/lifecycle-toolkit/commit/cb01b09ec5ec079595e5d64de70c541b432f4de0))


### Dependency Updates

* bump umbrella chart dependencies ([#3816](https://github.com/keptn/lifecycle-toolkit/issues/3816)) ([302150a](https://github.com/keptn/lifecycle-toolkit/commit/302150a60b07c88d9d1fe6412dfb91f89fb36c7f))
* **lifecycle-operator:** bump python and deno runtime images ([#3852](https://github.com/keptn/lifecycle-toolkit/issues/3852)) ([140498b](https://github.com/keptn/lifecycle-toolkit/commit/140498ba922640ee76d8a8a7330056e721eb8193))
* update github.com/keptn/lifecycle-toolkit/keptn-cert-manager digest to 17fa47b ([#3838](https://github.com/keptn/lifecycle-toolkit/issues/3838)) ([c8d56cd](https://github.com/keptn/lifecycle-toolkit/commit/c8d56cdacacc2730b6454377cec54e04a47853cd))
* update golang docker tag to v1.23.3 ([#3847](https://github.com/keptn/lifecycle-toolkit/issues/3847)) ([a200b38](https://github.com/keptn/lifecycle-toolkit/commit/a200b38e5c4f1b4d024e50585d68410f8ceba446))
* update module github.com/onsi/ginkgo/v2 to v2.21.0 ([#3834](https://github.com/keptn/lifecycle-toolkit/issues/3834)) ([cbd00cb](https://github.com/keptn/lifecycle-toolkit/commit/cbd00cb6a02c00dffc6961663910d27b717d16d8))
* update module github.com/onsi/gomega to v1.35.1 ([#3810](https://github.com/keptn/lifecycle-toolkit/issues/3810)) ([cada033](https://github.com/keptn/lifecycle-toolkit/commit/cada033a016569688be95a29269ecfa1688504cc))
* update module github.com/prometheus/client_golang to v1.20.5 ([#3809](https://github.com/keptn/lifecycle-toolkit/issues/3809)) ([3d72aa0](https://github.com/keptn/lifecycle-toolkit/commit/3d72aa0589ca76251109110390d91dd9f44f3343))
* update module google.golang.org/grpc to v1.68.0 ([#3850](https://github.com/keptn/lifecycle-toolkit/issues/3850)) ([776ec29](https://github.com/keptn/lifecycle-toolkit/commit/776ec298d1e882d2efd9bcf0aedef340299300d8))

## [1.2.0](https://github.com/keptn/lifecycle-toolkit/compare/lifecycle-operator-v1.1.1...lifecycle-operator-v1.2.0) (2024-10-31)


### Features

* add support for Kubernetes v1.31 ([#3785](https://github.com/keptn/lifecycle-toolkit/issues/3785)) ([2c5ba22](https://github.com/keptn/lifecycle-toolkit/commit/2c5ba22537e6e583c631aff78eca9b737c6bd69f))


### Other

* bump runtime versions ([#3813](https://github.com/keptn/lifecycle-toolkit/issues/3813)) ([e03976d](https://github.com/keptn/lifecycle-toolkit/commit/e03976d1e2a7f8e28f1a107d0a8b44b615764f40))
* release scheduler 1.0.2 ([#3709](https://github.com/keptn/lifecycle-toolkit/issues/3709)) ([8bf4406](https://github.com/keptn/lifecycle-toolkit/commit/8bf4406c5ad236558781773ad7ebafcc2cc937ed))
* update helm chart versions ([#3713](https://github.com/keptn/lifecycle-toolkit/issues/3713)) ([e052c92](https://github.com/keptn/lifecycle-toolkit/commit/e052c927cdb83215dfc2724e259f8fd72988a0a2))


### Dependency Updates

* update all golang.org/x packages ([#3729](https://github.com/keptn/lifecycle-toolkit/issues/3729)) ([d0a22de](https://github.com/keptn/lifecycle-toolkit/commit/d0a22deacb3af85c95e07c9b4d2d91120b7c07ca))
* update all golang.org/x packages ([#3754](https://github.com/keptn/lifecycle-toolkit/issues/3754)) ([736a4ac](https://github.com/keptn/lifecycle-toolkit/commit/736a4ac129b8f81a17e609090efadc6b1bc10939))
* update dependency kubernetes-sigs/controller-tools to v0.16.2 ([#3700](https://github.com/keptn/lifecycle-toolkit/issues/3700)) ([a0ff0ed](https://github.com/keptn/lifecycle-toolkit/commit/a0ff0ed8eccb9ef0da275f5a1bcce010d058ad7c))
* update dependency kubernetes-sigs/controller-tools to v0.16.4 ([#3733](https://github.com/keptn/lifecycle-toolkit/issues/3733)) ([a3dca53](https://github.com/keptn/lifecycle-toolkit/commit/a3dca534108788fb47c73e21a48f791070fddb31))
* update dependency kubernetes-sigs/controller-tools to v0.16.5 ([#3781](https://github.com/keptn/lifecycle-toolkit/issues/3781)) ([9e742e8](https://github.com/keptn/lifecycle-toolkit/commit/9e742e8975a45103d7f33929e442ed70e84ac7c3))
* update dependency kubernetes-sigs/kustomize to v5.5.0 ([#3793](https://github.com/keptn/lifecycle-toolkit/issues/3793)) ([62ff5ac](https://github.com/keptn/lifecycle-toolkit/commit/62ff5ac84966ab3dcbf1a640133a47e7301479bd))
* update github.com/keptn/lifecycle-toolkit/keptn-cert-manager digest to 8c76404 ([#3685](https://github.com/keptn/lifecycle-toolkit/issues/3685)) ([053a574](https://github.com/keptn/lifecycle-toolkit/commit/053a574ce7618a9b732ed946be5bc50c7d6c5b4d))
* update golang docker tag to v1.23.0 ([#3702](https://github.com/keptn/lifecycle-toolkit/issues/3702)) ([a4d7328](https://github.com/keptn/lifecycle-toolkit/commit/a4d73287d793bdd4681d5e8ba9fe56d185705b80))
* update golang docker tag to v1.23.1 ([#3732](https://github.com/keptn/lifecycle-toolkit/issues/3732)) ([8eb878f](https://github.com/keptn/lifecycle-toolkit/commit/8eb878f908467ab9a7140694d263de0fc352f2dd))
* update golang docker tag to v1.23.2 ([#3758](https://github.com/keptn/lifecycle-toolkit/issues/3758)) ([37a3f51](https://github.com/keptn/lifecycle-toolkit/commit/37a3f512da9ecd8eb8cf1dcb339647e9872e0b7e))
* update kubebuilder and k8s packages ([#3806](https://github.com/keptn/lifecycle-toolkit/issues/3806)) ([7b92373](https://github.com/keptn/lifecycle-toolkit/commit/7b92373b2110c28dec82d576747f821d47a98bb2))
* update kubernetes packages to v0.29.9 (patch) ([#3735](https://github.com/keptn/lifecycle-toolkit/issues/3735)) ([fcb3242](https://github.com/keptn/lifecycle-toolkit/commit/fcb3242218302ee406cdb10407e020b4ee516588))
* update module github.com/onsi/ginkgo/v2 to v2.20.2 ([#3724](https://github.com/keptn/lifecycle-toolkit/issues/3724)) ([7d62fad](https://github.com/keptn/lifecycle-toolkit/commit/7d62fadadbd251debe2b2c0c46a9042538a70c71))
* update module google.golang.org/grpc to v1.67.1 ([#3796](https://github.com/keptn/lifecycle-toolkit/issues/3796)) ([2902073](https://github.com/keptn/lifecycle-toolkit/commit/2902073b1fe0bee77d8f7af2b40fc49d0e183cbb))
* update module google.golang.org/protobuf to v1.35.1 ([#3797](https://github.com/keptn/lifecycle-toolkit/issues/3797)) ([2c5f1b3](https://github.com/keptn/lifecycle-toolkit/commit/2c5f1b310f34609dddb602059cbcca35c672ca6f))
* update opentelemetry-go monorepo (minor) ([#3798](https://github.com/keptn/lifecycle-toolkit/issues/3798)) ([878fa47](https://github.com/keptn/lifecycle-toolkit/commit/878fa47ce3f3c4a2fb6863169c13c408bc266d61))

## [1.1.1](https://github.com/keptn/lifecycle-toolkit/compare/lifecycle-operator-v1.1.0...lifecycle-operator-v1.1.1) (2024-08-27)


### Bug Fixes

* **lifecycle-operator:** update ownerReference when updating KeptnWorkload ([#3683](https://github.com/keptn/lifecycle-toolkit/issues/3683)) ([de2c441](https://github.com/keptn/lifecycle-toolkit/commit/de2c4412021b0aff8f2276830205f7fab1586c41))


### Other

* backport helm chart versions ([#3662](https://github.com/keptn/lifecycle-toolkit/issues/3662)) ([81320a5](https://github.com/keptn/lifecycle-toolkit/commit/81320a534f52cf16d3e59c76b451c92b3502b35f))


### Dependency Updates

* update all golang.org/x packages ([#3708](https://github.com/keptn/lifecycle-toolkit/issues/3708)) ([1014268](https://github.com/keptn/lifecycle-toolkit/commit/1014268f72066049bd707373feea367f680d0132))
* update golang docker tag to v1.22.6 ([#3692](https://github.com/keptn/lifecycle-toolkit/issues/3692)) ([73b30ee](https://github.com/keptn/lifecycle-toolkit/commit/73b30ee6a0748535d8af1160a7c0b8e2f2c04ec2))
* update kubernetes packages to v0.29.8 (patch) ([#3693](https://github.com/keptn/lifecycle-toolkit/issues/3693)) ([7ca1250](https://github.com/keptn/lifecycle-toolkit/commit/7ca12502889b1b74431a3210039299b620d37c3f))
* update module dario.cat/mergo to v1.0.1 ([#3694](https://github.com/keptn/lifecycle-toolkit/issues/3694)) ([80d8e9d](https://github.com/keptn/lifecycle-toolkit/commit/80d8e9de846adef30c5061c4e16e82698069b890))
* update module github.com/argoproj/argo-rollouts to v1.7.2 ([#3695](https://github.com/keptn/lifecycle-toolkit/issues/3695)) ([e0445d9](https://github.com/keptn/lifecycle-toolkit/commit/e0445d9d76d5589a7ae711e8d0dbfe617a7cb20c))
* update module github.com/onsi/ginkgo/v2 to v2.20.1 ([#3704](https://github.com/keptn/lifecycle-toolkit/issues/3704)) ([df11440](https://github.com/keptn/lifecycle-toolkit/commit/df1144084e337c8eda64e2635b8ee534fd10252f))

## [1.1.0](https://github.com/keptn/lifecycle-toolkit/compare/lifecycle-operator-v1.0.0...lifecycle-operator-v1.1.0) (2024-08-06)


### Features

* **helm-chart:** ability to set hostNetwork for lifecycle operator deployment ([#3500](https://github.com/keptn/lifecycle-toolkit/issues/3500)) ([c08bb07](https://github.com/keptn/lifecycle-toolkit/commit/c08bb07f9e59c6488062e76d09ff9dcff4311102))
* **lifecycle-operator:** introduce RestApiEnabled parameter in KeptnConfig ([#3620](https://github.com/keptn/lifecycle-toolkit/issues/3620)) ([1ab1f14](https://github.com/keptn/lifecycle-toolkit/commit/1ab1f14c290b0edbb473c33620a25e5ca706cb74))


### Other

* backport release helm chart versions ([#3460](https://github.com/keptn/lifecycle-toolkit/issues/3460)) ([95d6809](https://github.com/keptn/lifecycle-toolkit/commit/95d6809272eeb2b981b63f3718c4e799fc72b743))
* introduce unit test for non-blocking lifecycle operator mode ([#3581](https://github.com/keptn/lifecycle-toolkit/issues/3581)) ([d62f4ea](https://github.com/keptn/lifecycle-toolkit/commit/d62f4eaea739c97dcd5bcc70d3078bf77af2190e))


### Dependency Updates

* bump keptn-cert-manager to latest version ([#3659](https://github.com/keptn/lifecycle-toolkit/issues/3659)) ([82bfb2b](https://github.com/keptn/lifecycle-toolkit/commit/82bfb2bfd36e61df37f91bf0565762d443a60d45))
* update dependency golangci/golangci-lint to v1.59.1 ([#3631](https://github.com/keptn/lifecycle-toolkit/issues/3631)) ([29c7c1c](https://github.com/keptn/lifecycle-toolkit/commit/29c7c1cdab7ca4d95fd89622fc07f50d557c9e9c))
* update dependency kubernetes-sigs/controller-tools to v0.15.0 ([#3473](https://github.com/keptn/lifecycle-toolkit/issues/3473)) ([8987cd1](https://github.com/keptn/lifecycle-toolkit/commit/8987cd1773ba2e90e941e87caece0020b1100a18))
* update dependency kubernetes-sigs/kustomize to v5.4.2 ([#3548](https://github.com/keptn/lifecycle-toolkit/issues/3548)) ([fdecc9b](https://github.com/keptn/lifecycle-toolkit/commit/fdecc9b44c52a69407310dd85240b355a6b56c6e))
* update dependency kubernetes-sigs/kustomize to v5.4.3 ([#3617](https://github.com/keptn/lifecycle-toolkit/issues/3617)) ([8e583ab](https://github.com/keptn/lifecycle-toolkit/commit/8e583ab22897eb8039c4038d0f85081c0f1a3a79))
* update golang docker tag to v1.21.10 ([#3508](https://github.com/keptn/lifecycle-toolkit/issues/3508)) ([ed3409f](https://github.com/keptn/lifecycle-toolkit/commit/ed3409f1fa0f309e1fb0b1d971ec56e0c8f854bb))
* update golang docker tag to v1.21.11 ([#3552](https://github.com/keptn/lifecycle-toolkit/issues/3552)) ([6fdf850](https://github.com/keptn/lifecycle-toolkit/commit/6fdf8503c179264b1428d13dc00717cb2eb9d589))
* update golang docker tag to v1.21.12 ([#3604](https://github.com/keptn/lifecycle-toolkit/issues/3604)) ([0838454](https://github.com/keptn/lifecycle-toolkit/commit/083845486d2b5a94d75b49afbd1093c68b1d4523))
* update golang docker tag to v1.22.5 ([#3315](https://github.com/keptn/lifecycle-toolkit/issues/3315)) ([0d7a613](https://github.com/keptn/lifecycle-toolkit/commit/0d7a613997697ba4db937f5e6193bc23c6b34551))
* update golang.org/x/exp digest to 7f521ea ([#3551](https://github.com/keptn/lifecycle-toolkit/issues/3551)) ([4fcf576](https://github.com/keptn/lifecycle-toolkit/commit/4fcf576ef59182b883177a4f4be1dff8822b04ac))
* update golang.org/x/exp digest to 8a7402a ([#3615](https://github.com/keptn/lifecycle-toolkit/issues/3615)) ([c3dfe86](https://github.com/keptn/lifecycle-toolkit/commit/c3dfe86c59880f6b88e63b6f9dbf0272c8a6989e))
* update google.golang.org/grpc to v1.64.1 ([#3582](https://github.com/keptn/lifecycle-toolkit/issues/3582)) ([bb1070a](https://github.com/keptn/lifecycle-toolkit/commit/bb1070a760eb364bb8527f40c0206c0ee1bf7847))
* update k8s to v1.29.7 ([#3652](https://github.com/keptn/lifecycle-toolkit/issues/3652)) ([f008022](https://github.com/keptn/lifecycle-toolkit/commit/f0080220d2afa84415f995d80b30433edcf9a9f4))
* update kubernetes packages to v0.28.10 (patch) ([#3522](https://github.com/keptn/lifecycle-toolkit/issues/3522)) ([d9c1d70](https://github.com/keptn/lifecycle-toolkit/commit/d9c1d70eef3170c62b43dc2896df5b8d630af515))
* update kubernetes packages to v0.28.11 (patch) ([#3554](https://github.com/keptn/lifecycle-toolkit/issues/3554)) ([b548057](https://github.com/keptn/lifecycle-toolkit/commit/b5480576c8eff5b00ce5d0269502e754ea6f97ec))
* update kubernetes packages to v0.28.12 (patch) ([#3607](https://github.com/keptn/lifecycle-toolkit/issues/3607)) ([7792a1b](https://github.com/keptn/lifecycle-toolkit/commit/7792a1b85211d2e61cabaff4f640e6c299fddd12))
* update module github.com/argoproj/argo-rollouts to v1.7.1 ([#3650](https://github.com/keptn/lifecycle-toolkit/issues/3650)) ([e6bd7d2](https://github.com/keptn/lifecycle-toolkit/commit/e6bd7d2648120d13774410ccf502abddb019199b))
* update module github.com/go-logr/logr to v1.4.2 ([#3534](https://github.com/keptn/lifecycle-toolkit/issues/3534)) ([b9409f5](https://github.com/keptn/lifecycle-toolkit/commit/b9409f50b61cc2e7c7032abb6cbcc27fc5ef6c99))
* update module github.com/onsi/ginkgo/v2 to v2.17.2 ([#3489](https://github.com/keptn/lifecycle-toolkit/issues/3489)) ([4a7e9cf](https://github.com/keptn/lifecycle-toolkit/commit/4a7e9cf80e083e35b39da7e8204ff1cfb5d842ae))
* update module github.com/onsi/ginkgo/v2 to v2.17.3 ([#3512](https://github.com/keptn/lifecycle-toolkit/issues/3512)) ([c0d2afe](https://github.com/keptn/lifecycle-toolkit/commit/c0d2afe9eda7d73654538cef636439fb5168beef))
* update module github.com/onsi/ginkgo/v2 to v2.19.1 ([#3636](https://github.com/keptn/lifecycle-toolkit/issues/3636)) ([0018af4](https://github.com/keptn/lifecycle-toolkit/commit/0018af46771d1d8a26356206d1a8e471a90613b8))
* update module github.com/onsi/gomega to v1.33.0 ([#3468](https://github.com/keptn/lifecycle-toolkit/issues/3468)) ([e8ddfde](https://github.com/keptn/lifecycle-toolkit/commit/e8ddfdecd1a5fef73013d968a0913ac4aa9d744f))
* update module github.com/onsi/gomega to v1.33.1 ([#3495](https://github.com/keptn/lifecycle-toolkit/issues/3495)) ([80caa18](https://github.com/keptn/lifecycle-toolkit/commit/80caa185c2909089c746c392f32dda54edd42dc3))
* update module github.com/prometheus/client_golang to v1.19.1 ([#3517](https://github.com/keptn/lifecycle-toolkit/issues/3517)) ([63c593e](https://github.com/keptn/lifecycle-toolkit/commit/63c593eb93d609475711888518d757b34a07a64b))
* update module golang.org/x/exp to v0.0.0-20240506185415-9bf2ced13842 ([#3507](https://github.com/keptn/lifecycle-toolkit/issues/3507)) ([0d2f645](https://github.com/keptn/lifecycle-toolkit/commit/0d2f6459f366e191f4a2792d97496d1c1d7e3965))
* update module google.golang.org/grpc to v1.64.0 ([#3526](https://github.com/keptn/lifecycle-toolkit/issues/3526)) ([19a0f7c](https://github.com/keptn/lifecycle-toolkit/commit/19a0f7ce421ca204f7d2eb168a2487426c8220e4))
* update module google.golang.org/grpc to v1.65.0 ([#3639](https://github.com/keptn/lifecycle-toolkit/issues/3639)) ([5fabe3b](https://github.com/keptn/lifecycle-toolkit/commit/5fabe3b0b627fb915ea4b7d9a670a7dae4360daf))
* update module google.golang.org/protobuf to v1.34.0 ([#3496](https://github.com/keptn/lifecycle-toolkit/issues/3496)) ([5862111](https://github.com/keptn/lifecycle-toolkit/commit/586211185649f9542dd7782d551d1ad04e3c809f))
* update module google.golang.org/protobuf to v1.34.1 ([#3518](https://github.com/keptn/lifecycle-toolkit/issues/3518)) ([6146e6d](https://github.com/keptn/lifecycle-toolkit/commit/6146e6d142b3e75e52730ef66ae76e6d60fab3bc))
* update module google.golang.org/protobuf to v1.34.2 ([#3568](https://github.com/keptn/lifecycle-toolkit/issues/3568)) ([08b5684](https://github.com/keptn/lifecycle-toolkit/commit/08b568411f444461a30a286be5cc63f27f0bb8c2))
* update module sigs.k8s.io/controller-runtime to v0.16.6 ([#3523](https://github.com/keptn/lifecycle-toolkit/issues/3523)) ([e044b51](https://github.com/keptn/lifecycle-toolkit/commit/e044b51c794b601edaa815855a8e4fb195f0842c))
* update opentelemetry-go monorepo (minor) ([#3480](https://github.com/keptn/lifecycle-toolkit/issues/3480)) ([11fa477](https://github.com/keptn/lifecycle-toolkit/commit/11fa4771e478972f09a8c63eb240215dca23bcef))
* update opentelemetry-go monorepo (minor) ([#3645](https://github.com/keptn/lifecycle-toolkit/issues/3645)) ([ebc0c29](https://github.com/keptn/lifecycle-toolkit/commit/ebc0c29fdea99acdc37bb4aaa6ea79be56d13cc6))
* update python and deno runtimes ([#3655](https://github.com/keptn/lifecycle-toolkit/issues/3655)) ([41f1cbe](https://github.com/keptn/lifecycle-toolkit/commit/41f1cbebb8fd4ef85130bf7771bb599c17b88f5d))

## [1.0.0](https://github.com/keptn/lifecycle-toolkit/compare/lifecycle-operator-v0.9.2...lifecycle-operator-v1.0.0) (2024-04-24)


### Features

* **helm-chart:** make charts Openshift compliant ([#3415](https://github.com/keptn/lifecycle-toolkit/issues/3415)) ([32f077a](https://github.com/keptn/lifecycle-toolkit/commit/32f077aa875d591d5b17eca01b2e75cafeaae44d))
* **lifecycle-operator:** introduce v1 API version ([#3344](https://github.com/keptn/lifecycle-toolkit/issues/3344)) ([1d851c5](https://github.com/keptn/lifecycle-toolkit/commit/1d851c5d72d6a1cf4678b838bd5176c2de89ece7))
* **lifecycle-operator:** move API HUB version to v1 ([#3350](https://github.com/keptn/lifecycle-toolkit/issues/3350)) ([eed393b](https://github.com/keptn/lifecycle-toolkit/commit/eed393b859000866ecc271d117b9971c3693cfda))
* **metrics-operator:** use v1 API in operator logic ([#3269](https://github.com/keptn/lifecycle-toolkit/issues/3269)) ([e9a584b](https://github.com/keptn/lifecycle-toolkit/commit/e9a584bc28ce6306362c722fed8849f5d5be0bda))


### Bug Fixes

* introduce missing Role into keptn-cert-manager helm charts ([#3435](https://github.com/keptn/lifecycle-toolkit/issues/3435)) ([16afdaa](https://github.com/keptn/lifecycle-toolkit/commit/16afdaaf4ae56179d0f725ae9f9e9ae96709f042))


### Other

* bump helm charts versions ([#3303](https://github.com/keptn/lifecycle-toolkit/issues/3303)) ([19cbe9f](https://github.com/keptn/lifecycle-toolkit/commit/19cbe9fda082015d4a61d23c1276d599f6370cec))
* bump runtime versions to latest released ([#3455](https://github.com/keptn/lifecycle-toolkit/issues/3455)) ([4034df7](https://github.com/keptn/lifecycle-toolkit/commit/4034df773a97d59cb758e82ccb5ab12ddab97fad))
* **lifecycle-operator:** bump release version ([#3458](https://github.com/keptn/lifecycle-toolkit/issues/3458)) ([ec765bc](https://github.com/keptn/lifecycle-toolkit/commit/ec765bcdd14ab82b8ecd1d70f6ab0332ccd3bb71))
* **lifecycle-operator:** clean up unused API logic ([#3351](https://github.com/keptn/lifecycle-toolkit/issues/3351)) ([016dc07](https://github.com/keptn/lifecycle-toolkit/commit/016dc07236785d7bcfbccfe2a806b70e1ec44421))
* **lifecycle-operator:** remove deprecated fields in KeptnTaskDefinition ([#3345](https://github.com/keptn/lifecycle-toolkit/issues/3345)) ([15a8ae3](https://github.com/keptn/lifecycle-toolkit/commit/15a8ae3394fe1347434dda87764b86c8cefc1637))


### Dependency Updates

* bump golang.org/x/net to v0.23.0 ([#3388](https://github.com/keptn/lifecycle-toolkit/issues/3388)) ([e9c1dda](https://github.com/keptn/lifecycle-toolkit/commit/e9c1dda3489117422160d53467d2155b1ca2bad3))
* update dependency kubernetes-sigs/kustomize to v5.4.1 ([#3394](https://github.com/keptn/lifecycle-toolkit/issues/3394)) ([2dda172](https://github.com/keptn/lifecycle-toolkit/commit/2dda17232aab5542929a5fa73378cd2399a2f5e5))
* update golang docker tag to v1.21.9 ([#3384](https://github.com/keptn/lifecycle-toolkit/issues/3384)) ([e4f1a6a](https://github.com/keptn/lifecycle-toolkit/commit/e4f1a6adefc2670a6c18efbaf416aee80eb2584a))
* update golang.org/x/exp digest to 93d18d7 ([#3400](https://github.com/keptn/lifecycle-toolkit/issues/3400)) ([5a9e73b](https://github.com/keptn/lifecycle-toolkit/commit/5a9e73b6296b4aea435979aa2f2a2b70e1241628))
* update golang.org/x/exp digest to a685a6e ([#3346](https://github.com/keptn/lifecycle-toolkit/issues/3346)) ([f0e7571](https://github.com/keptn/lifecycle-toolkit/commit/f0e7571054f9bf5ecb9d5c01f471da76d75ea488))
* update golang.org/x/exp digest to c0f41cb ([#3389](https://github.com/keptn/lifecycle-toolkit/issues/3389)) ([4407d07](https://github.com/keptn/lifecycle-toolkit/commit/4407d077084e968184d9c1b3d96746019bb6db4c))
* update golang.org/x/exp digest to fe59bbe ([#3427](https://github.com/keptn/lifecycle-toolkit/issues/3427)) ([f1a17ff](https://github.com/keptn/lifecycle-toolkit/commit/f1a17ffd5163a6aa048724eaa017c51fa42b35b2))
* update kubernetes packages to v0.28.8 (patch) ([#3300](https://github.com/keptn/lifecycle-toolkit/issues/3300)) ([435e722](https://github.com/keptn/lifecycle-toolkit/commit/435e722776b69c6e7acbf3631d81cdeafc9815ec))
* update module github.com/keptn/lifecycle-toolkit/keptn-cert-manager ([#3318](https://github.com/keptn/lifecycle-toolkit/issues/3318)) ([f765187](https://github.com/keptn/lifecycle-toolkit/commit/f765187aa304850ea6f4cf0014e2de4bb1fecafa))
* update module github.com/onsi/ginkgo/v2 to v2.16.0 ([#3319](https://github.com/keptn/lifecycle-toolkit/issues/3319)) ([10bc1c0](https://github.com/keptn/lifecycle-toolkit/commit/10bc1c02759f6eebe4f30812f868a9e6465c5e3d))
* update module github.com/onsi/ginkgo/v2 to v2.17.0 ([#3339](https://github.com/keptn/lifecycle-toolkit/issues/3339)) ([997a63c](https://github.com/keptn/lifecycle-toolkit/commit/997a63cbeb8e7707a9c7c6fb21a1f6feabb75e1d))
* update module github.com/onsi/ginkgo/v2 to v2.17.1 ([#3363](https://github.com/keptn/lifecycle-toolkit/issues/3363)) ([a34b8e5](https://github.com/keptn/lifecycle-toolkit/commit/a34b8e5959f775fe632cad1b7c74f6de46ff9aa0))
* update module google.golang.org/grpc to v1.62.2 ([#3391](https://github.com/keptn/lifecycle-toolkit/issues/3391)) ([213204b](https://github.com/keptn/lifecycle-toolkit/commit/213204b9685ac355f42701ece024c7df18bf4308))
* update module google.golang.org/grpc to v1.63.2 ([#3422](https://github.com/keptn/lifecycle-toolkit/issues/3422)) ([7da06b7](https://github.com/keptn/lifecycle-toolkit/commit/7da06b74b2cd28fb5a092d4b2028c1bed99b01a3))
* update module k8s.io/apimachinery to v0.28.9 ([#3433](https://github.com/keptn/lifecycle-toolkit/issues/3433)) ([a75d65e](https://github.com/keptn/lifecycle-toolkit/commit/a75d65e6528509276af4060aea6f85a02d03ad30))
* update opentelemetry-go monorepo (minor) ([#3408](https://github.com/keptn/lifecycle-toolkit/issues/3408)) ([15ebf45](https://github.com/keptn/lifecycle-toolkit/commit/15ebf45f382f8661abc15d7ae71feeea57126431))

## [0.9.2](https://github.com/keptn/lifecycle-toolkit/compare/lifecycle-operator-v0.9.1...lifecycle-operator-v0.9.2) (2024-03-20)


### Features

* **lifecycle-operator:** add namespace to `deploymentduration` metrics ([#3292](https://github.com/keptn/lifecycle-toolkit/issues/3292)) ([0735e31](https://github.com/keptn/lifecycle-toolkit/commit/0735e31db1967da85e346f9f028a67f178611606))


### Bug Fixes

* **helm-chart:** introduce cert volumes to metrics and lifecycle operators ([#3247](https://github.com/keptn/lifecycle-toolkit/issues/3247)) ([b7744dd](https://github.com/keptn/lifecycle-toolkit/commit/b7744dd36289b9d7c843f1679481830a843f90ac))
* **lifecycle-operator:** remove noops tracer for evaluations ([#3290](https://github.com/keptn/lifecycle-toolkit/issues/3290)) ([4948dc5](https://github.com/keptn/lifecycle-toolkit/commit/4948dc5f20424bbe9e21c31abbe4c4147b729410))
* security vulnerabilities ([#3230](https://github.com/keptn/lifecycle-toolkit/issues/3230)) ([1d099d7](https://github.com/keptn/lifecycle-toolkit/commit/1d099d7a4c9b5e856de52932693b97c29bea3122))


### Other

* backport helm release versions ([#3241](https://github.com/keptn/lifecycle-toolkit/issues/3241)) ([074bb16](https://github.com/keptn/lifecycle-toolkit/commit/074bb165a9a70c8daa187f215f2dd74f3159b95d))
* bump Go base images and pipelines version to 1.21 ([#3218](https://github.com/keptn/lifecycle-toolkit/issues/3218)) ([de01ca4](https://github.com/keptn/lifecycle-toolkit/commit/de01ca493b307d8c27701552549b982e22281a2e))
* **lifecycle-operator:** remove failAction parameter from KeptnEvaluation helm charts ([#3275](https://github.com/keptn/lifecycle-toolkit/issues/3275)) ([fffc75b](https://github.com/keptn/lifecycle-toolkit/commit/fffc75baf6d665d9de25a437177f5866d0040d63))
* release scheduler 0.9.2 ([#3228](https://github.com/keptn/lifecycle-toolkit/issues/3228)) ([998c6a9](https://github.com/keptn/lifecycle-toolkit/commit/998c6a9c0e6f11713b99113420276436be694159))
* update chart dependencies ([#3179](https://github.com/keptn/lifecycle-toolkit/issues/3179)) ([b8efdd5](https://github.com/keptn/lifecycle-toolkit/commit/b8efdd50002231a06bac9c5ab02fcdbadea4c60d))


### Dependency Updates

* bump python and deno runtimes to latest version ([#3295](https://github.com/keptn/lifecycle-toolkit/issues/3295)) ([65616cd](https://github.com/keptn/lifecycle-toolkit/commit/65616cd2ac9da98c755e28d3f045750e582172f4))
* update golang.org/x/exp digest to a85f2c6 ([#3288](https://github.com/keptn/lifecycle-toolkit/issues/3288)) ([62a8c14](https://github.com/keptn/lifecycle-toolkit/commit/62a8c14a06ec81b6a42450195d9ff341f7aaff41))
* update golang.org/x/exp digest to c7f7c64 ([#3272](https://github.com/keptn/lifecycle-toolkit/issues/3272)) ([a2f0f00](https://github.com/keptn/lifecycle-toolkit/commit/a2f0f00172e379d64c47b99b4b9ef7181fac321c))
* update module github.com/keptn/lifecycle-toolkit/keptn-cert-manager to v0.8.0 ([#3167](https://github.com/keptn/lifecycle-toolkit/issues/3167)) ([7ad3344](https://github.com/keptn/lifecycle-toolkit/commit/7ad3344e555e848fb38ac55d7e521700a9a33f9f))
* update module google.golang.org/grpc to v1.62.1 ([#3281](https://github.com/keptn/lifecycle-toolkit/issues/3281)) ([f86c49a](https://github.com/keptn/lifecycle-toolkit/commit/f86c49a8e4a72ceccab95f15d0dcde2a4e7dbfb0))

## [0.9.1](https://github.com/keptn/lifecycle-toolkit/compare/lifecycle-operator-v0.9.0...lifecycle-operator-v0.9.1) (2024-03-04)


### Features

* add global value for imagePullPolicy ([#2807](https://github.com/keptn/lifecycle-toolkit/issues/2807)) ([5596d12](https://github.com/keptn/lifecycle-toolkit/commit/5596d1252b164e469aa122c0ebda8526ccbca888))
* **lifecycle-operator:** adapt KeptnConfig reconciler to set up blockedDeployment parameter ([#3112](https://github.com/keptn/lifecycle-toolkit/issues/3112)) ([c8ad8b1](https://github.com/keptn/lifecycle-toolkit/commit/c8ad8b1c5157539746d176f8361ca8f1a2f071d8))
* **lifecycle-operator:** adapt KeptnConfig reconciler to set up observabilityTimeout parameter ([#3154](https://github.com/keptn/lifecycle-toolkit/issues/3154)) ([f14a1ff](https://github.com/keptn/lifecycle-toolkit/commit/f14a1ff586cde3b0ace20d8b89fc6b4a94768630))
* **lifecycle-operator:** adapt WorkloadVersionReconciler logic to use ObservabilityTimeout for workload deployment ([#3160](https://github.com/keptn/lifecycle-toolkit/issues/3160)) ([e98d10e](https://github.com/keptn/lifecycle-toolkit/commit/e98d10eb8f038f3cfd8bf373a8731417c1811f45))
* **lifecycle-operator:** add Counter meter for promotion phase ([#3105](https://github.com/keptn/lifecycle-toolkit/issues/3105)) ([fa146fa](https://github.com/keptn/lifecycle-toolkit/commit/fa146face9f02ad6843bac8ba20d1503c2affa03))
* **lifecycle-operator:** add feature flag for enabling promotion tasks ([#3055](https://github.com/keptn/lifecycle-toolkit/issues/3055)) ([d4044c1](https://github.com/keptn/lifecycle-toolkit/commit/d4044c1c1a6fc9126aac456ba6e3bca05a5d541e))
* **lifecycle-operator:** implement promotion task ([#3057](https://github.com/keptn/lifecycle-toolkit/issues/3057)) ([e165600](https://github.com/keptn/lifecycle-toolkit/commit/e165600ac59c018e115915bebbcce50fbd5a7e5b))
* **lifecycle-operator:** introduce a possibility to configure number of retries and interval for KeptnEvaluationDefinition ([#3141](https://github.com/keptn/lifecycle-toolkit/issues/3141)) ([65f7327](https://github.com/keptn/lifecycle-toolkit/commit/65f73275d9b6112aba0844fd42c773ed26de2867))
* **lifecycle-operator:** introduce blockDeployment parameter into KeptnConfig ([#3111](https://github.com/keptn/lifecycle-toolkit/issues/3111)) ([ab5b89d](https://github.com/keptn/lifecycle-toolkit/commit/ab5b89d963fe78b15c8951cecda1a6c25a190a8f))
* **lifecycle-operator:** introduce non-blocking deployment functionality for application lifecycle ([#3113](https://github.com/keptn/lifecycle-toolkit/issues/3113)) ([bf78974](https://github.com/keptn/lifecycle-toolkit/commit/bf78974ba9ac11ecb3a21585193822671cd7c325))
* **lifecycle-operator:** introduce ObservabilityTimeout parameter in KeptnConfig ([#3149](https://github.com/keptn/lifecycle-toolkit/issues/3149)) ([79de15e](https://github.com/keptn/lifecycle-toolkit/commit/79de15e94c1e006db970a4bd3ac5def72a1f82c4))
* **lifecycle-operator:** introduce ObservabilityTimeout parameter in KeptnWorkload ([#3153](https://github.com/keptn/lifecycle-toolkit/issues/3153)) ([0e88438](https://github.com/keptn/lifecycle-toolkit/commit/0e8843828a7d0f495e19c545a698f54ecb5ec8cc))
* **lifecycle-operator:** introduce promotionTask parameters in KeptnAppContext ([#3056](https://github.com/keptn/lifecycle-toolkit/issues/3056)) ([c2c3af3](https://github.com/keptn/lifecycle-toolkit/commit/c2c3af3ee3f7576a4a6e9e79c8f02c9e93eea6b4))


### Bug Fixes

* **lifecycle-operator:** close root spans of failed AppVersions/WorkloadVersions ([#3174](https://github.com/keptn/lifecycle-toolkit/issues/3174)) ([120005b](https://github.com/keptn/lifecycle-toolkit/commit/120005b48597b286782721d18be8f3605eb59210))
* **lifecycle-operator:** retrieve KeptnEvaluationDefinition before creating KeptnEvaluation ([#3144](https://github.com/keptn/lifecycle-toolkit/issues/3144)) ([54a9b8b](https://github.com/keptn/lifecycle-toolkit/commit/54a9b8b85e8ee2fc02cc3cc375104d174fef8eeb))


### Other

* bump go version to 1.21 ([#3006](https://github.com/keptn/lifecycle-toolkit/issues/3006)) ([8236c25](https://github.com/keptn/lifecycle-toolkit/commit/8236c25da7ec3768e76d12eb2e8f5765a005ecfa))
* bump helm chart dependencies ([#2991](https://github.com/keptn/lifecycle-toolkit/issues/2991)) ([49ee351](https://github.com/keptn/lifecycle-toolkit/commit/49ee3511fd6e425ac095bd7f16ecd1dae6258eb0))
* **lifecycle-operator:** clean up leftover logic for supporting standalone Pods as Workloads ([#3140](https://github.com/keptn/lifecycle-toolkit/issues/3140)) ([17321bc](https://github.com/keptn/lifecycle-toolkit/commit/17321bcd18627479259c963ad2c96c5d0562ac8d))
* **lifecycle-operator:** remove unused FailAction parameter from KeptnEvaluation ([#3138](https://github.com/keptn/lifecycle-toolkit/issues/3138)) ([4febd99](https://github.com/keptn/lifecycle-toolkit/commit/4febd992682290473823d6cb8d826533e8dcef76))
* **lifecycle-operator:** revert unused ObservabilityTimeout parameter from KeptnWorkload ([#3163](https://github.com/keptn/lifecycle-toolkit/issues/3163)) ([7b68ac8](https://github.com/keptn/lifecycle-toolkit/commit/7b68ac8df2fb317e2099a498aa995369f547f5d1))


### Docs

* fix generation of underlying types ([#3150](https://github.com/keptn/lifecycle-toolkit/issues/3150)) ([a387a88](https://github.com/keptn/lifecycle-toolkit/commit/a387a88d3ad249e9eee34c43e3e391bc3709dab4))
* review keptntaskdefinition examples ([#3085](https://github.com/keptn/lifecycle-toolkit/issues/3085)) ([d0a0c43](https://github.com/keptn/lifecycle-toolkit/commit/d0a0c4348459624f0659db5d1d5484db3335f314))
* update keptn state descriptions in our CRDs ([#3124](https://github.com/keptn/lifecycle-toolkit/issues/3124)) ([d87b288](https://github.com/keptn/lifecycle-toolkit/commit/d87b288b8e88a34908228a2e3bae8686857f680c))


### Dependency Updates

* update ghcr.io/keptn/deno-runtime docker tag to v2.0.2 ([#3156](https://github.com/keptn/lifecycle-toolkit/issues/3156)) ([4452584](https://github.com/keptn/lifecycle-toolkit/commit/445258414a093646c5eadf893220cfcbc953dd5b))
* update ghcr.io/keptn/python-runtime docker tag to v1.0.3 ([#3152](https://github.com/keptn/lifecycle-toolkit/issues/3152)) ([85d8fd0](https://github.com/keptn/lifecycle-toolkit/commit/85d8fd0b12cf05a9b73bb54b4904ad80f3cc4214))
* update golang.org/x/exp digest to 814bf88 ([#3109](https://github.com/keptn/lifecycle-toolkit/issues/3109)) ([8610295](https://github.com/keptn/lifecycle-toolkit/commit/86102953785511b8ae73e56820aa5d796c357a2d))
* update golang.org/x/exp digest to ec58324 ([#3043](https://github.com/keptn/lifecycle-toolkit/issues/3043)) ([d736aef](https://github.com/keptn/lifecycle-toolkit/commit/d736aefcd323b144bd2771ffd7677c03aa57be0a))
* update helm release common to v0.1.4 ([#3114](https://github.com/keptn/lifecycle-toolkit/issues/3114)) ([12b2e58](https://github.com/keptn/lifecycle-toolkit/commit/12b2e58e085fd40cf5c04ca0e5eb071823777701))
* update kubernetes packages to v0.28.7 (patch) ([#3062](https://github.com/keptn/lifecycle-toolkit/issues/3062)) ([8698803](https://github.com/keptn/lifecycle-toolkit/commit/8698803ff60b71d658d60bfc0c6b8b3d4282798d))
* update module github.com/argoproj/argo-rollouts to v1.6.6 ([#3061](https://github.com/keptn/lifecycle-toolkit/issues/3061)) ([9c4297b](https://github.com/keptn/lifecycle-toolkit/commit/9c4297b077b67d921306db6f824839aa425754e9))
* update module github.com/cloudevents/sdk-go/v2 to v2.15.1 ([#3118](https://github.com/keptn/lifecycle-toolkit/issues/3118)) ([73c2a31](https://github.com/keptn/lifecycle-toolkit/commit/73c2a31ae535ece58f4869aa6fc85e3a0c1a6ae0))
* update module github.com/keptn/lifecycle-toolkit/keptn-cert-manager to v0.8.0 ([#2974](https://github.com/keptn/lifecycle-toolkit/issues/2974)) ([cd36e8d](https://github.com/keptn/lifecycle-toolkit/commit/cd36e8df8a7fabfbbe443200f4659c0b0a8be937))
* update module github.com/keptn/lifecycle-toolkit/keptn-cert-manager to v0.8.0 ([#3047](https://github.com/keptn/lifecycle-toolkit/issues/3047)) ([d6b4a64](https://github.com/keptn/lifecycle-toolkit/commit/d6b4a642298586dccab464486de45906364a7898))
* update module github.com/keptn/lifecycle-toolkit/keptn-cert-manager to v0.8.0 ([#3158](https://github.com/keptn/lifecycle-toolkit/issues/3158)) ([d775416](https://github.com/keptn/lifecycle-toolkit/commit/d775416edcc5519a7134c2b52a13b469d883890f))
* update module github.com/stretchr/testify to v1.9.0 ([#3171](https://github.com/keptn/lifecycle-toolkit/issues/3171)) ([d334790](https://github.com/keptn/lifecycle-toolkit/commit/d3347903ad91c33ba4bf664277c53024eb02825a))
* update module google.golang.org/grpc to v1.61.1 ([#3072](https://github.com/keptn/lifecycle-toolkit/issues/3072)) ([3c9d1f3](https://github.com/keptn/lifecycle-toolkit/commit/3c9d1f3bb7dd7ebfda56563a235ff8c8ce6c61f6))
* update module google.golang.org/grpc to v1.62.0 ([#3119](https://github.com/keptn/lifecycle-toolkit/issues/3119)) ([ea061db](https://github.com/keptn/lifecycle-toolkit/commit/ea061dbb272f3fa3bf0ce99bd33617bc1dc98a18))
* update module sigs.k8s.io/controller-runtime to v0.16.4 ([#3033](https://github.com/keptn/lifecycle-toolkit/issues/3033)) ([f576707](https://github.com/keptn/lifecycle-toolkit/commit/f57670729a18cfdb391c3af5ffdd92de6a330ee5))
* update module sigs.k8s.io/controller-runtime to v0.16.5 ([#3073](https://github.com/keptn/lifecycle-toolkit/issues/3073)) ([599e2d8](https://github.com/keptn/lifecycle-toolkit/commit/599e2d8712ed7d7b614026a0038d238ed0833b37))
* update module sigs.k8s.io/yaml to v1.4.0 ([#2984](https://github.com/keptn/lifecycle-toolkit/issues/2984)) ([584aff6](https://github.com/keptn/lifecycle-toolkit/commit/584aff65411cca24b69c4efa84428eb8188f05b1))
* update opentelemetry-go monorepo (minor) ([#3129](https://github.com/keptn/lifecycle-toolkit/issues/3129)) ([513986d](https://github.com/keptn/lifecycle-toolkit/commit/513986d4e6bb481906ecba33b19da85ffe5b7e5d))
* update opentelemetry-go monorepo (patch) ([#3010](https://github.com/keptn/lifecycle-toolkit/issues/3010)) ([a6d1724](https://github.com/keptn/lifecycle-toolkit/commit/a6d172444765dbe8e34ae2fd92d390b66afe69f1))
* update opentelemetry-go monorepo to v1.23.1 (minor) ([#3092](https://github.com/keptn/lifecycle-toolkit/issues/3092)) ([ac71144](https://github.com/keptn/lifecycle-toolkit/commit/ac711443311ee241c58125944bee4a7ffc10d026))

## [0.9.0](https://github.com/keptn/lifecycle-toolkit/compare/lifecycle-operator-v0.8.3...lifecycle-operator-v0.9.0) (2024-02-08)


### ⚠ BREAKING CHANGES

* **lifecycle-operator:** Pre/Post evaluations and tasks for an application are now defined in the newly introduced `KeptnAppContext` instead of the `KeptnApp` CRD. `KeptnApps` are now fully managed by the operator and are not intended to be created by the user. The version of a `KeptnApp` will be automatically derived as a function of all workloads that belong to the same application.
* **lifecycle-operator:** move API HUB version to v1beta1 ([#2772](https://github.com/keptn/lifecycle-toolkit/issues/2772))
* **lifecycle-operator:** The environment variable `OTEL_COLLECTOR_URL` is not supported in the lifecycle-operator anymore, and the OTel collector URL is now only set via the `spec.OTelCollectorUrl` property of the `KeptnConfig` CRD. This means that, in order to use Keptn's OpenTelemetry capabilities, the `spec.OtelCollectorUrl` needs to be specified in the `KeptnConfig` resource.
* rename KLT to Keptn ([#2554](https://github.com/keptn/lifecycle-toolkit/issues/2554))
* **lifecycle-operator:** The environment variable giving deno and python runtime access to context information has been renamed from `CONTEXT` to `KEPTN_CONTEXT`

### Features

* add annotation to select container for version extraction ([#2471](https://github.com/keptn/lifecycle-toolkit/issues/2471)) ([d093860](https://github.com/keptn/lifecycle-toolkit/commit/d093860732798b0edb58abedf567558a2c07ad21))
* add configurable service account to KeptnTasks ([#2254](https://github.com/keptn/lifecycle-toolkit/issues/2254)) ([e7db66f](https://github.com/keptn/lifecycle-toolkit/commit/e7db66f91a638759d9d95ef34fa22f59a8a37f9d))
* introduce configurable support of cert-manager.io CA injection ([#2811](https://github.com/keptn/lifecycle-toolkit/issues/2811)) ([d6d83c7](https://github.com/keptn/lifecycle-toolkit/commit/d6d83c7f67a18a4b30aabe774a8fa2c93399f301))
* introduce configurable TTLSecondsAfterFinished for tasks ([#2404](https://github.com/keptn/lifecycle-toolkit/issues/2404)) ([8341dbf](https://github.com/keptn/lifecycle-toolkit/commit/8341dbf256b23d342226b9c44a2057e4fd775854))
* **lifecycle-operator:** add `KEPTN_CONTEXT` to task container env vars ([#2516](https://github.com/keptn/lifecycle-toolkit/issues/2516)) ([a18a833](https://github.com/keptn/lifecycle-toolkit/commit/a18a83306fed5636a971565e12b2e71d315b75b4))
* **lifecycle-operator:** add context metadata and traceParent of current phase to tasks ([#2858](https://github.com/keptn/lifecycle-toolkit/issues/2858)) ([0798406](https://github.com/keptn/lifecycle-toolkit/commit/0798406108b545e8f7debceae5dc1cb28f0a8d11))
* **lifecycle-operator:** add Helm value for DORA metrics port ([#2571](https://github.com/keptn/lifecycle-toolkit/issues/2571)) ([bf472a3](https://github.com/keptn/lifecycle-toolkit/commit/bf472a34efcda14ccb78869aa141a8cd981f4839))
* **lifecycle-operator:** add option to exclude additional namespaces ([#2536](https://github.com/keptn/lifecycle-toolkit/issues/2536)) ([fd42ac7](https://github.com/keptn/lifecycle-toolkit/commit/fd42ac7325927fa6f2f0cfe6875f055fd2cd1be0))
* **lifecycle-operator:** introduce keptnappcontext crd ([#2769](https://github.com/keptn/lifecycle-toolkit/issues/2769)) ([4e7751a](https://github.com/keptn/lifecycle-toolkit/commit/4e7751ae7344d8334db5bd8e6e4463e87eb3314b))
* **lifecycle-operator:** move API HUB version to v1beta1 ([#2772](https://github.com/keptn/lifecycle-toolkit/issues/2772)) ([5d7ebbd](https://github.com/keptn/lifecycle-toolkit/commit/5d7ebbdc2ef55714e62dd8ad8b600a1098f9adef))
* **lifecycle-operator:** propagate KeptnAppVersion Context Metadata to KeptnWorkloadVersion span ([#2859](https://github.com/keptn/lifecycle-toolkit/issues/2859)) ([5c14bf5](https://github.com/keptn/lifecycle-toolkit/commit/5c14bf59e813db10f953ea019c8d61d7ec2e8f6d))
* **lifecycle-operator:** propagate metadata from deployment annotations ([#2832](https://github.com/keptn/lifecycle-toolkit/issues/2832)) ([6f700ce](https://github.com/keptn/lifecycle-toolkit/commit/6f700ce453ff1c26f353bc5e109c8b3e1840b283))
* **lifecycle-operator:** rename CONTEXT to KEPTN_CONTEXT in task runtimes ([#2521](https://github.com/keptn/lifecycle-toolkit/issues/2521)) ([a7322bd](https://github.com/keptn/lifecycle-toolkit/commit/a7322bd9266fa1589d77b06675d70d1a9e6c29ac))
* **lifecycle-operator:** support imagePullSecrets in KeptnTaskDefinitions ([#2549](https://github.com/keptn/lifecycle-toolkit/issues/2549)) ([c71d868](https://github.com/keptn/lifecycle-toolkit/commit/c71d86864ba48a82d9f66d57e93521d99c426970))
* **lifecycle-operator:** support linked spans in KeptnAppVersion ([#2833](https://github.com/keptn/lifecycle-toolkit/issues/2833)) ([36e19b2](https://github.com/keptn/lifecycle-toolkit/commit/36e19b2a9f9706722a05bd13e46340bd68922265))


### Bug Fixes

* **helm-chart:** remove double templating of annotations ([#2770](https://github.com/keptn/lifecycle-toolkit/issues/2770)) ([b7a1d29](https://github.com/keptn/lifecycle-toolkit/commit/b7a1d291223eddd9ac83425c71c8c1a515f25f58))
* **lifecycle-operator:** adopt KeptnApp name from either Keptn or k8s label ([#2440](https://github.com/keptn/lifecycle-toolkit/issues/2440)) ([3185943](https://github.com/keptn/lifecycle-toolkit/commit/318594309af9653253f84b35f86e9b6675c572ca))
* **lifecycle-operator:** duplicate version in project file ([#2767](https://github.com/keptn/lifecycle-toolkit/issues/2767)) ([c7ed8a6](https://github.com/keptn/lifecycle-toolkit/commit/c7ed8a69c9af658606761261216e6c00bae5ffa8))
* **lifecycle-operator:** fix app deployment span structure ([#2352](https://github.com/keptn/lifecycle-toolkit/issues/2352)) ([64c1919](https://github.com/keptn/lifecycle-toolkit/commit/64c1919f43378650a6677b3b5baa91776e96bef9))
* **lifecycle-operator:** introduce separate controller for removing scheduling gates from pods ([#2946](https://github.com/keptn/lifecycle-toolkit/issues/2946)) ([9fa3770](https://github.com/keptn/lifecycle-toolkit/commit/9fa3770bbf3a2a2374993144df4fa469837aa7a0))
* **lifecycle-operator:** make sure spec of KeptnWorkloadVersion is consistent with KeptnWorkload ([#2926](https://github.com/keptn/lifecycle-toolkit/issues/2926)) ([f2f8c29](https://github.com/keptn/lifecycle-toolkit/commit/f2f8c296a1b7f9746c55c2781c22727c62a2bab3))


### Other

* adapt examples to use v1beta1 API resources ([#2868](https://github.com/keptn/lifecycle-toolkit/issues/2868)) ([587773f](https://github.com/keptn/lifecycle-toolkit/commit/587773fbea63dbf575879bd3bec447fe55ac4311))
* adapt helm charts to the new Keptn naming ([#2564](https://github.com/keptn/lifecycle-toolkit/issues/2564)) ([9ee4583](https://github.com/keptn/lifecycle-toolkit/commit/9ee45834bfa4dcedcbe99362d5d58b9febe3caae))
* add config for spell checker action, fix typos ([#2443](https://github.com/keptn/lifecycle-toolkit/issues/2443)) ([eac178f](https://github.com/keptn/lifecycle-toolkit/commit/eac178f650962208449553086d54d26d27fa4da3))
* add KeptnApp migration script ([#2959](https://github.com/keptn/lifecycle-toolkit/issues/2959)) ([7311422](https://github.com/keptn/lifecycle-toolkit/commit/7311422791f5429fa77ac18da857e4f14b502eba))
* clean up deprecated API resources from helm charts ([#2800](https://github.com/keptn/lifecycle-toolkit/issues/2800)) ([43d092d](https://github.com/keptn/lifecycle-toolkit/commit/43d092d17f852d60f4e29a2887128b33a3fd2764))
* clean up unused volumes ([#2638](https://github.com/keptn/lifecycle-toolkit/issues/2638)) ([32be4db](https://github.com/keptn/lifecycle-toolkit/commit/32be4db7ed35676967148fdc93cbe1a378220afa))
* **helm-chart:** generate umbrella chart lock ([#2391](https://github.com/keptn/lifecycle-toolkit/issues/2391)) ([55e12d4](https://github.com/keptn/lifecycle-toolkit/commit/55e12d4a6c3b5cd0fbb2cd6b8b8d29f2b7c8c500))
* **lifecycle-operator:** adapt KeptnAppVersionReconciler to make use of PhaseHandler interface ([#2463](https://github.com/keptn/lifecycle-toolkit/issues/2463)) ([2511e05](https://github.com/keptn/lifecycle-toolkit/commit/2511e05cefe8876c0bb67fbf9763213ef81a81a0))
* **lifecycle-operator:** introduce PhaseHandler interface to be used in KeptnWorkloadVersion reconciler ([#2450](https://github.com/keptn/lifecycle-toolkit/issues/2450)) ([7d4b431](https://github.com/keptn/lifecycle-toolkit/commit/7d4b431af5a6e9deec03784b04267d9711c93f17))
* **lifecycle-operator:** introduce v1beta1 lifecycle API ([#2640](https://github.com/keptn/lifecycle-toolkit/issues/2640)) ([11b7ea2](https://github.com/keptn/lifecycle-toolkit/commit/11b7ea2bbf6fc22dc781fdf1e7afdde1b6b54035))
* **lifecycle-operator:** make evaluationhandler injectable in `KeptnWorkloadVersionController` ([#2299](https://github.com/keptn/lifecycle-toolkit/issues/2299)) ([211b272](https://github.com/keptn/lifecycle-toolkit/commit/211b2727cdce51378a33ce92775f231e2b025117))
* **lifecycle-operator:** make evaluationhandler injectable in KeptnAppVersionController ([#2402](https://github.com/keptn/lifecycle-toolkit/issues/2402)) ([a060859](https://github.com/keptn/lifecycle-toolkit/commit/a06085954ff3fd508f6c0ebec806df78babd8dc4))
* **lifecycle-operator:** propagate Context Metadata to KeptnAppVersion ([#2848](https://github.com/keptn/lifecycle-toolkit/issues/2848)) ([5fac158](https://github.com/keptn/lifecycle-toolkit/commit/5fac158a7ffed67f7502fe03683138d717ea1acd))
* **lifecycle-operator:** refactor `WorkloadVersionReconciler` ([#2417](https://github.com/keptn/lifecycle-toolkit/issues/2417)) ([c41f909](https://github.com/keptn/lifecycle-toolkit/commit/c41f909044a40485bee07ddcaa59a0d9924a1bf1))
* **lifecycle-operator:** remove `OTEL_COLLECTOR_URL` env var in favour of related option in `KeptnConfig` CRD ([#2593](https://github.com/keptn/lifecycle-toolkit/issues/2593)) ([df0a5b4](https://github.com/keptn/lifecycle-toolkit/commit/df0a5b4a9ec04326a044bc5a79a6babf54a13363))
* **lifecycle-operator:** remove pre post deploy task evaluation v1beta1 ([#2782](https://github.com/keptn/lifecycle-toolkit/issues/2782)) ([6e992d7](https://github.com/keptn/lifecycle-toolkit/commit/6e992d72313792d7e3024fd99599ca8658c98737))
* **lifecycle-operator:** renamed TracerFactory to tracerFactory workloadversion ([#2428](https://github.com/keptn/lifecycle-toolkit/issues/2428)) ([8c10e38](https://github.com/keptn/lifecycle-toolkit/commit/8c10e38435fd41079b9853b4e7f2039549ff80b9))
* **lifecycle-operator:** split controllers/common package into multiple ([#2386](https://github.com/keptn/lifecycle-toolkit/issues/2386)) ([cbda641](https://github.com/keptn/lifecycle-toolkit/commit/cbda6410e12e24cb8af3754a6f396e4b99164e14))
* re-generate CRD manifests ([#2830](https://github.com/keptn/lifecycle-toolkit/issues/2830)) ([c0b1942](https://github.com/keptn/lifecycle-toolkit/commit/c0b1942e8f2ddd177776ed681432016d81805724))
* remove performance-test workflow and relative makefile entry ([#2706](https://github.com/keptn/lifecycle-toolkit/issues/2706)) ([8599276](https://github.com/keptn/lifecycle-toolkit/commit/859927698453bbd1f718b347c73f70da6596713f))
* rename Keptn default namespace to 'keptn-system' ([#2565](https://github.com/keptn/lifecycle-toolkit/issues/2565)) ([aec1148](https://github.com/keptn/lifecycle-toolkit/commit/aec11489451ab1b0bcd69a6b90b0d45f69c5df7c))
* rename KLT to Keptn ([#2554](https://github.com/keptn/lifecycle-toolkit/issues/2554)) ([15b0ac0](https://github.com/keptn/lifecycle-toolkit/commit/15b0ac0b36b8081b85b63f36e94b00065bcc8b22))
* revert helm charts bump ([#2806](https://github.com/keptn/lifecycle-toolkit/issues/2806)) ([2e85214](https://github.com/keptn/lifecycle-toolkit/commit/2e85214ecd6112e9f9af750d9bde2d491dc8ae73))
* update to crd generator to v0.0.10 ([#2329](https://github.com/keptn/lifecycle-toolkit/issues/2329)) ([525ae03](https://github.com/keptn/lifecycle-toolkit/commit/525ae03725f374d0b056c6da2fd7af3e4062f7a2))
* upgrade helm chart versions ([#2801](https://github.com/keptn/lifecycle-toolkit/issues/2801)) ([ad26093](https://github.com/keptn/lifecycle-toolkit/commit/ad2609373c4819fc560766e64bc032fcfd801889))


### Docs

* remove Scarf transparent pixels ([#2590](https://github.com/keptn/lifecycle-toolkit/issues/2590)) ([95851fa](https://github.com/keptn/lifecycle-toolkit/commit/95851fa52cb3a6565a4b52ae0e8b00dcc9861a3b))


### Dependency Updates

* update dependency kubernetes-sigs/controller-tools to v0.14.0 ([#2797](https://github.com/keptn/lifecycle-toolkit/issues/2797)) ([71f20a6](https://github.com/keptn/lifecycle-toolkit/commit/71f20a63f8e307d6e94c9c2df79a1258ab147ede))
* update dependency kubernetes-sigs/kustomize to v5.3.0 ([#2659](https://github.com/keptn/lifecycle-toolkit/issues/2659)) ([8877921](https://github.com/keptn/lifecycle-toolkit/commit/8877921b8be3052ce61a4f8decd96537c93df27a))
* update ghcr.io/keptn/deno-runtime docker tag to v2 ([#2969](https://github.com/keptn/lifecycle-toolkit/issues/2969)) ([ea3e77d](https://github.com/keptn/lifecycle-toolkit/commit/ea3e77da83cb1d170e10329ecafcc837a03ee095))
* update ghcr.io/keptn/python-runtime docker tag to v1.0.2 ([#2968](https://github.com/keptn/lifecycle-toolkit/issues/2968)) ([ae7d394](https://github.com/keptn/lifecycle-toolkit/commit/ae7d3943c8aee315273eda0c13909a1cc8cb4b52))
* update github.com/keptn/lifecycle-toolkit/klt-cert-manager digest to 0677987 ([#2429](https://github.com/keptn/lifecycle-toolkit/issues/2429)) ([f718913](https://github.com/keptn/lifecycle-toolkit/commit/f7189131cefcc6fe9a42a560d696ca019afc541f))
* update github.com/keptn/lifecycle-toolkit/klt-cert-manager digest to 964fd25 ([#2485](https://github.com/keptn/lifecycle-toolkit/issues/2485)) ([f7124d0](https://github.com/keptn/lifecycle-toolkit/commit/f7124d034dd6e1558581de35f449bf08b2c73bab))
* update github.com/keptn/lifecycle-toolkit/klt-cert-manager digest to d2c3e14 ([#2375](https://github.com/keptn/lifecycle-toolkit/issues/2375)) ([b945bf8](https://github.com/keptn/lifecycle-toolkit/commit/b945bf875e435ab713d5b37cf8c0415948942bf1))
* update golang.org/x/exp digest to 1b97071 ([#2875](https://github.com/keptn/lifecycle-toolkit/issues/2875)) ([20f5705](https://github.com/keptn/lifecycle-toolkit/commit/20f5705141e252afbe76834be739f305ac3b273a))
* update golang.org/x/exp digest to 2c58cdc ([#2971](https://github.com/keptn/lifecycle-toolkit/issues/2971)) ([fddbce7](https://github.com/keptn/lifecycle-toolkit/commit/fddbce72ea68e3f507adf61d76f259eab4303cdb))
* update keptn/common helm chart to 0.1.3 ([#2831](https://github.com/keptn/lifecycle-toolkit/issues/2831)) ([29187fa](https://github.com/keptn/lifecycle-toolkit/commit/29187fa7eeab148b7188b4c3f05317cc291c15e4))
* update kubernetes packages to v0.28.5 (patch) ([#2714](https://github.com/keptn/lifecycle-toolkit/issues/2714)) ([192c0b1](https://github.com/keptn/lifecycle-toolkit/commit/192c0b16fc0852dca572448d8caeb113b0e21d40))
* update kubernetes packages to v0.28.6 (patch) ([#2827](https://github.com/keptn/lifecycle-toolkit/issues/2827)) ([da080fa](https://github.com/keptn/lifecycle-toolkit/commit/da080fafadef25028f9e4b1a78d8a862e58b47e7))
* update module github.com/argoproj/argo-rollouts to v1.6.2 ([#2411](https://github.com/keptn/lifecycle-toolkit/issues/2411)) ([9e9d731](https://github.com/keptn/lifecycle-toolkit/commit/9e9d731084ee453c26a133f32cf82d58b275b4da))
* update module github.com/argoproj/argo-rollouts to v1.6.3 ([#2652](https://github.com/keptn/lifecycle-toolkit/issues/2652)) ([e386ec6](https://github.com/keptn/lifecycle-toolkit/commit/e386ec643fa7a202fda32d5f1126581b7c084109))
* update module github.com/argoproj/argo-rollouts to v1.6.4 ([#2679](https://github.com/keptn/lifecycle-toolkit/issues/2679)) ([95380bb](https://github.com/keptn/lifecycle-toolkit/commit/95380bb523e71f63b3f7d0769934b85931b5fec8))
* update module github.com/argoproj/argo-rollouts to v1.6.5 ([#2892](https://github.com/keptn/lifecycle-toolkit/issues/2892)) ([7c8b14f](https://github.com/keptn/lifecycle-toolkit/commit/7c8b14f8c09be6995eb341582177dfed8038b7cd))
* update module github.com/cloudevents/sdk-go/v2 to v2.15.0 ([#2845](https://github.com/keptn/lifecycle-toolkit/issues/2845)) ([22dd509](https://github.com/keptn/lifecycle-toolkit/commit/22dd5093e263979f466b08f20f72a8763528c957))
* update module github.com/go-logr/logr to v1.4.1 ([#2726](https://github.com/keptn/lifecycle-toolkit/issues/2726)) ([3598999](https://github.com/keptn/lifecycle-toolkit/commit/3598999e1cfce6ee528fb5fb777c0b7b7c21678a))
* update module github.com/keptn/lifecycle-toolkit/keptn-cert-manager to v0.8.0 ([#2534](https://github.com/keptn/lifecycle-toolkit/issues/2534)) ([94007a0](https://github.com/keptn/lifecycle-toolkit/commit/94007a03cd9bd7e09bad79feb12b27b615a75151))
* update module github.com/keptn/lifecycle-toolkit/keptn-cert-manager to v2.0.0 ([#2668](https://github.com/keptn/lifecycle-toolkit/issues/2668)) ([be6523b](https://github.com/keptn/lifecycle-toolkit/commit/be6523b39b431e9c1cfac51ac553c4c71e0ad4a1))
* update module github.com/onsi/ginkgo/v2 to v2.13.1 ([#2486](https://github.com/keptn/lifecycle-toolkit/issues/2486)) ([14dcd27](https://github.com/keptn/lifecycle-toolkit/commit/14dcd27f4b1e67803332a8dc53b42b67c7bb2030))
* update module github.com/onsi/ginkgo/v2 to v2.13.2 ([#2624](https://github.com/keptn/lifecycle-toolkit/issues/2624)) ([197c7db](https://github.com/keptn/lifecycle-toolkit/commit/197c7db78a5baf754e773ab79c5cd6a5ab9c5591))
* update module github.com/onsi/ginkgo/v2 to v2.14.0 ([#2808](https://github.com/keptn/lifecycle-toolkit/issues/2808)) ([17b0cb1](https://github.com/keptn/lifecycle-toolkit/commit/17b0cb1314778f5f1b65f4d1029ecca41bb50d3a))
* update module github.com/onsi/ginkgo/v2 to v2.15.0 ([#2855](https://github.com/keptn/lifecycle-toolkit/issues/2855)) ([1c4f410](https://github.com/keptn/lifecycle-toolkit/commit/1c4f410f5571f02254eda4c5027c8a5e3822b28e))
* update module github.com/onsi/gomega to v1.29.0 ([#2379](https://github.com/keptn/lifecycle-toolkit/issues/2379)) ([98e420a](https://github.com/keptn/lifecycle-toolkit/commit/98e420a4b2138e90e2f87c399139bd8e5a90cef5))
* update module github.com/onsi/gomega to v1.30.0 ([#2478](https://github.com/keptn/lifecycle-toolkit/issues/2478)) ([398b949](https://github.com/keptn/lifecycle-toolkit/commit/398b9493414ab5d70bd76d94b038456e58813e70))
* update module github.com/onsi/gomega to v1.31.1 ([#2856](https://github.com/keptn/lifecycle-toolkit/issues/2856)) ([d0817a7](https://github.com/keptn/lifecycle-toolkit/commit/d0817a7118e58af5326a43f1a059f2eddfa36215))
* update module github.com/prometheus/client_golang to v1.18.0 ([#2764](https://github.com/keptn/lifecycle-toolkit/issues/2764)) ([67fa60b](https://github.com/keptn/lifecycle-toolkit/commit/67fa60b8581fee0b6200f8f877b396a39df32d58))
* update module golang.org/x/net to v0.18.0 ([#2479](https://github.com/keptn/lifecycle-toolkit/issues/2479)) ([6ddd8ee](https://github.com/keptn/lifecycle-toolkit/commit/6ddd8eeec5eabb0c67b5a7b9965a34368f62c8d5))
* update module golang.org/x/net to v0.19.0 ([#2619](https://github.com/keptn/lifecycle-toolkit/issues/2619)) ([af2d0a5](https://github.com/keptn/lifecycle-toolkit/commit/af2d0a509b670792e06e2d05ab4be261d3bb54f4))
* update module golang.org/x/net to v0.20.0 ([#2786](https://github.com/keptn/lifecycle-toolkit/issues/2786)) ([8294c7b](https://github.com/keptn/lifecycle-toolkit/commit/8294c7b471d7f4d33961513e056c36ba14c940c7))
* update module google.golang.org/grpc to v1.60.0 ([#2681](https://github.com/keptn/lifecycle-toolkit/issues/2681)) ([7dd45a3](https://github.com/keptn/lifecycle-toolkit/commit/7dd45a33fba8fd3235e40202ece9057cef429bb6))
* update module google.golang.org/grpc to v1.60.1 ([#2724](https://github.com/keptn/lifecycle-toolkit/issues/2724)) ([31d69dd](https://github.com/keptn/lifecycle-toolkit/commit/31d69dd33df76f0a5f9b2d46af822e5f43e681a5))
* update module google.golang.org/grpc to v1.61.0 ([#2888](https://github.com/keptn/lifecycle-toolkit/issues/2888)) ([7a56cbd](https://github.com/keptn/lifecycle-toolkit/commit/7a56cbd1f528bb73c1070611d6b28005c875fe36))
* update module k8s.io/apimachinery to v0.28.4 ([#2514](https://github.com/keptn/lifecycle-toolkit/issues/2514)) ([c25c236](https://github.com/keptn/lifecycle-toolkit/commit/c25c236ecc37dc1f33b75a172cee2422bdb416ba))
* update opentelemetry-go monorepo (minor) ([#2487](https://github.com/keptn/lifecycle-toolkit/issues/2487)) ([a5d492a](https://github.com/keptn/lifecycle-toolkit/commit/a5d492abe1757bcac0259ae059d137d8afa6d57a))
* update opentelemetry-go monorepo (minor) ([#2535](https://github.com/keptn/lifecycle-toolkit/issues/2535)) ([7e3f5e6](https://github.com/keptn/lifecycle-toolkit/commit/7e3f5e6a14edeb1063765c3122f90e4c7659c943))
* update opentelemetry-go monorepo (minor) ([#2865](https://github.com/keptn/lifecycle-toolkit/issues/2865)) ([be0ecde](https://github.com/keptn/lifecycle-toolkit/commit/be0ecde8088af5e4a43d01951f6b7f354267308d))

## [0.8.3](https://github.com/keptn/lifecycle-toolkit/compare/lifecycle-operator-v0.8.2...lifecycle-operator-v0.8.3) (2023-10-31)


### Features

* adapt code to use KeptnWorkloadVersion instead of KeptnWorkloadInstance ([#2255](https://github.com/keptn/lifecycle-toolkit/issues/2255)) ([c06fae1](https://github.com/keptn/lifecycle-toolkit/commit/c06fae13daa2aa98a3daf71abafe0e8ce4e5f4a3))
* add test and lint cmd to makefiles ([#2176](https://github.com/keptn/lifecycle-toolkit/issues/2176)) ([c55e0a9](https://github.com/keptn/lifecycle-toolkit/commit/c55e0a9f368c82ad3032eb676edd59e68b29fad6))
* create new Keptn umbrella Helm chart ([#2214](https://github.com/keptn/lifecycle-toolkit/issues/2214)) ([41bd47b](https://github.com/keptn/lifecycle-toolkit/commit/41bd47b7748c4d645243a4dae165651bbfd3533f))
* generalize helm chart ([#2282](https://github.com/keptn/lifecycle-toolkit/issues/2282)) ([81334eb](https://github.com/keptn/lifecycle-toolkit/commit/81334ebec4d8afda27902b6e854c4c637a3daa87))
* **lifecycle-operator:** add helm chart for lifecycle operator ([#2200](https://github.com/keptn/lifecycle-toolkit/issues/2200)) ([9f0853f](https://github.com/keptn/lifecycle-toolkit/commit/9f0853fca2b92c9636e76dc77666148d86078af7))
* **lifecycle-operator:** automatically decide for scheduler installation based on k8s version ([#2212](https://github.com/keptn/lifecycle-toolkit/issues/2212)) ([25976ea](https://github.com/keptn/lifecycle-toolkit/commit/25976ead3fb1d95634ee3a00a7d37b3e98b2ec06))
* **lifecycle-operator:** introduce functions for SchedulingGates functionality ([#2140](https://github.com/keptn/lifecycle-toolkit/issues/2140)) ([b40503e](https://github.com/keptn/lifecycle-toolkit/commit/b40503ef6c867033994383767ad5149eb08ab8da))
* **lifecycle-operator:** introduce metric showing readiness of operator ([#2152](https://github.com/keptn/lifecycle-toolkit/issues/2152)) ([c0e3f48](https://github.com/keptn/lifecycle-toolkit/commit/c0e3f48dd0e34084c7d2d8e469e73c3f2865ea48))
* **lifecycle-operator:** introduce option to enable lifecycle orchestration only for specific namespaces ([#2244](https://github.com/keptn/lifecycle-toolkit/issues/2244)) ([12caf03](https://github.com/keptn/lifecycle-toolkit/commit/12caf031d336c7a34e495b36daccb5ec3524ae49))
* **lifecycle-operator:** introduce v1alpha4 API version for KeptnWorkloadInstance ([#2250](https://github.com/keptn/lifecycle-toolkit/issues/2250)) ([d95dc10](https://github.com/keptn/lifecycle-toolkit/commit/d95dc1037ce22296aff65d6ad6fa420e96172d5d))
* **metrics-operator:** add support for user-friendly duration string for specifying time frame ([#2147](https://github.com/keptn/lifecycle-toolkit/issues/2147)) ([34e5384](https://github.com/keptn/lifecycle-toolkit/commit/34e5384bb434836658a7bf375c4fc8de765023b6))
* move helm docs into values files ([#2281](https://github.com/keptn/lifecycle-toolkit/issues/2281)) ([bd1a37b](https://github.com/keptn/lifecycle-toolkit/commit/bd1a37b324e25d07e88e7c4d1ad8150a7b3d4dac))
* support scheduling gates in integration tests ([#2149](https://github.com/keptn/lifecycle-toolkit/issues/2149)) ([3ff67d5](https://github.com/keptn/lifecycle-toolkit/commit/3ff67d5632f287613f337c7418aa5e28e616d536))


### Bug Fixes

* change klt to keptn for annotations and certs ([#2229](https://github.com/keptn/lifecycle-toolkit/issues/2229)) ([608a75e](https://github.com/keptn/lifecycle-toolkit/commit/608a75ebb73006b82b370b40e86b83ee874764e8))
* helm charts image registry, image pull policy and install action ([#2361](https://github.com/keptn/lifecycle-toolkit/issues/2361)) ([76ed884](https://github.com/keptn/lifecycle-toolkit/commit/76ed884498971c87c48cdab6fea822dfcf3e6e2f))
* **lifecycle-operator:** make sure the CloudEvents endpoint from the KeptnConfig is applied ([#2289](https://github.com/keptn/lifecycle-toolkit/issues/2289)) ([b5d9fc0](https://github.com/keptn/lifecycle-toolkit/commit/b5d9fc0b182ff3d1a777dabec74314df3157edbb))
* **lifecycle-operator:** remove hardcoded keptn namespace ([#2141](https://github.com/keptn/lifecycle-toolkit/issues/2141)) ([f10b447](https://github.com/keptn/lifecycle-toolkit/commit/f10b4470bdc4346e6ccd17fecc92c8bd5675c7e5))
* update kustomization.yaml to avoid usage of deprecated patches/configs ([#2004](https://github.com/keptn/lifecycle-toolkit/issues/2004)) ([8d70fac](https://github.com/keptn/lifecycle-toolkit/commit/8d70fac1f9469107257976659fb8b7b414d0455b))
* update outdated CRDs in helm chart templates ([#2123](https://github.com/keptn/lifecycle-toolkit/issues/2123)) ([34c9d11](https://github.com/keptn/lifecycle-toolkit/commit/34c9d11a1dd34b181d2d1a9e5c61fd75638aaebf))


### Other

* adapt Makefile command to run unit tests ([#2072](https://github.com/keptn/lifecycle-toolkit/issues/2072)) ([2db2569](https://github.com/keptn/lifecycle-toolkit/commit/2db25691748beedbb02ed92806d327067c422285))
* **lifecycle-operator:** improve logging ([#2253](https://github.com/keptn/lifecycle-toolkit/issues/2253)) ([8dd3394](https://github.com/keptn/lifecycle-toolkit/commit/8dd3394087cf0d445ec0b3bad0a54242ad9f4f26))
* **lifecycle-operator:** refactor pod mutating webhook ([#2233](https://github.com/keptn/lifecycle-toolkit/issues/2233)) ([c2cc89a](https://github.com/keptn/lifecycle-toolkit/commit/c2cc89a3ad3ac0fef3410adb1c0b24aa10e8dc66))
* **lifecycle-operator:** remove direct dependency on jsonpatch ([#2187](https://github.com/keptn/lifecycle-toolkit/issues/2187)) ([d7fce2a](https://github.com/keptn/lifecycle-toolkit/commit/d7fce2a320bd34cd41d564e0d528675e5a1cd93e))
* **lifecycle-operator:** remove spans created by webhook ([#2331](https://github.com/keptn/lifecycle-toolkit/issues/2331)) ([9f21fb6](https://github.com/keptn/lifecycle-toolkit/commit/9f21fb62284e806f6356341315873f98e0c4fd29))
* **lifecycle-operator:** remove spans for reconciliation loops, adjust log levels ([#2310](https://github.com/keptn/lifecycle-toolkit/issues/2310)) ([d73008c](https://github.com/keptn/lifecycle-toolkit/commit/d73008ccaa5e028f1551392b9c68a4ea0315350e))
* regenerate CRDs ([#2074](https://github.com/keptn/lifecycle-toolkit/issues/2074)) ([63f5dc1](https://github.com/keptn/lifecycle-toolkit/commit/63f5dc1bc3dfd696de3730ed3949c0f99abdecc0))
* update k8s version ([#1701](https://github.com/keptn/lifecycle-toolkit/issues/1701)) ([010d7cd](https://github.com/keptn/lifecycle-toolkit/commit/010d7cd48c2e26993e25de607f30b40513c9cd61))
* update release please config to work with umbrella chart ([#2357](https://github.com/keptn/lifecycle-toolkit/issues/2357)) ([6ff3a5f](https://github.com/keptn/lifecycle-toolkit/commit/6ff3a5f64e394504fd5e7b67f0ac0a608428c1be))


### Docs

* adapt KeptnTask example to changes in API  ([#2124](https://github.com/keptn/lifecycle-toolkit/issues/2124)) ([bcc64e8](https://github.com/keptn/lifecycle-toolkit/commit/bcc64e814d7735bc330d2d0b3b52eccf7a51dbbe))


### Dependency Updates

* update dependency kubernetes-sigs/kustomize to v5.2.1 ([#2308](https://github.com/keptn/lifecycle-toolkit/issues/2308)) ([6653a47](https://github.com/keptn/lifecycle-toolkit/commit/6653a47d4156c0e60aa471f11a643a2664669023))
* update ghcr.io/keptn/deno-runtime docker tag to v1.0.2 ([#2367](https://github.com/keptn/lifecycle-toolkit/issues/2367)) ([6c17203](https://github.com/keptn/lifecycle-toolkit/commit/6c1720356fab6b4a9d1c0dae30e76e6d5c135c70))
* update ghcr.io/keptn/python-runtime docker tag to v1.0.1 ([#2368](https://github.com/keptn/lifecycle-toolkit/issues/2368)) ([134191a](https://github.com/keptn/lifecycle-toolkit/commit/134191a523c6d278771ad1f3421e4ae68dad4de9))
* update ghcr.io/keptn/scheduler docker tag to v0.8.3 ([#2374](https://github.com/keptn/lifecycle-toolkit/issues/2374)) ([16a4a14](https://github.com/keptn/lifecycle-toolkit/commit/16a4a147905fe19b319010e880730ee46e6c5965))
* update github.com/keptn/lifecycle-toolkit/klt-cert-manager digest to 010d7cd ([#2106](https://github.com/keptn/lifecycle-toolkit/issues/2106)) ([2ef614a](https://github.com/keptn/lifecycle-toolkit/commit/2ef614ad08dbeed1301889ed560375a2bb9e737c))
* update github.com/keptn/lifecycle-toolkit/klt-cert-manager digest to 066be3e ([#2274](https://github.com/keptn/lifecycle-toolkit/issues/2274)) ([c6d9c52](https://github.com/keptn/lifecycle-toolkit/commit/c6d9c524afa61e7c87553b89ebab1a2b8daa1438))
* update github.com/keptn/lifecycle-toolkit/klt-cert-manager digest to 099a457 ([#2169](https://github.com/keptn/lifecycle-toolkit/issues/2169)) ([643ae4e](https://github.com/keptn/lifecycle-toolkit/commit/643ae4e69ac527a342eed20c0e078c6b0e8cdd97))
* update github.com/keptn/lifecycle-toolkit/klt-cert-manager digest to 3077e31 ([#2313](https://github.com/keptn/lifecycle-toolkit/issues/2313)) ([cf52945](https://github.com/keptn/lifecycle-toolkit/commit/cf529455f2e99119e48ba433d28e8aecc31cad00))
* update github.com/keptn/lifecycle-toolkit/klt-cert-manager digest to 4342d33 ([#2177](https://github.com/keptn/lifecycle-toolkit/issues/2177)) ([2b5267c](https://github.com/keptn/lifecycle-toolkit/commit/2b5267c0a88b7f68167314d624c54453a326b5ce))
* update github.com/keptn/lifecycle-toolkit/klt-cert-manager digest to 469578e ([#2038](https://github.com/keptn/lifecycle-toolkit/issues/2038)) ([d240e56](https://github.com/keptn/lifecycle-toolkit/commit/d240e56fbc0b42caad04d8393ec59e55f1013efa))
* update github.com/keptn/lifecycle-toolkit/klt-cert-manager digest to 5efa650 ([#2155](https://github.com/keptn/lifecycle-toolkit/issues/2155)) ([fa8c891](https://github.com/keptn/lifecycle-toolkit/commit/fa8c8912825ad0bbc3f75b7a037e856bac6dad93))
* update github.com/keptn/lifecycle-toolkit/klt-cert-manager digest to 608a75e ([#2231](https://github.com/keptn/lifecycle-toolkit/issues/2231)) ([26ff714](https://github.com/keptn/lifecycle-toolkit/commit/26ff714800bb605bfe58b61da432237132edf072))
* update github.com/keptn/lifecycle-toolkit/klt-cert-manager digest to 6566e7d ([#2143](https://github.com/keptn/lifecycle-toolkit/issues/2143)) ([9e7fe83](https://github.com/keptn/lifecycle-toolkit/commit/9e7fe8353dd8c84fae96081c2bb7522ab7ff7f5a))
* update github.com/keptn/lifecycle-toolkit/klt-cert-manager digest to 8dd3394 ([#2271](https://github.com/keptn/lifecycle-toolkit/issues/2271)) ([b29fc99](https://github.com/keptn/lifecycle-toolkit/commit/b29fc999aef2c214b7b45a1161b226e85e3eaffe))
* update github.com/keptn/lifecycle-toolkit/klt-cert-manager digest to a15b038 ([#2205](https://github.com/keptn/lifecycle-toolkit/issues/2205)) ([1592926](https://github.com/keptn/lifecycle-toolkit/commit/1592926d6c70062cd632969f991531aa9b00f0de))
* update github.com/keptn/lifecycle-toolkit/klt-cert-manager digest to a656512 ([#2230](https://github.com/keptn/lifecycle-toolkit/issues/2230)) ([f11fdb9](https://github.com/keptn/lifecycle-toolkit/commit/f11fdb959b4e207d3704361870d515f61ad92360))
* update github.com/keptn/lifecycle-toolkit/klt-cert-manager digest to b2853f9 ([#2094](https://github.com/keptn/lifecycle-toolkit/issues/2094)) ([b9019cd](https://github.com/keptn/lifecycle-toolkit/commit/b9019cd96a161c4e0c4dd08e3ddbabd152ea921c))
* update github.com/keptn/lifecycle-toolkit/klt-cert-manager digest to c1166ff ([#2242](https://github.com/keptn/lifecycle-toolkit/issues/2242)) ([aa53137](https://github.com/keptn/lifecycle-toolkit/commit/aa531375032468a6e0d1b3a9f6eb3e6e9b0c998b))
* update github.com/keptn/lifecycle-toolkit/klt-cert-manager digest to f2f3a0c ([#2132](https://github.com/keptn/lifecycle-toolkit/issues/2132)) ([2039d36](https://github.com/keptn/lifecycle-toolkit/commit/2039d36f427e22bfe692fade207319747ee15083))
* update github.com/keptn/lifecycle-toolkit/klt-cert-manager digest to f2f8dfe ([#2297](https://github.com/keptn/lifecycle-toolkit/issues/2297)) ([e13b9be](https://github.com/keptn/lifecycle-toolkit/commit/e13b9be3217fdd5bf3a646dd3d6ba49438cbb9e6))
* update github.com/keptn/lifecycle-toolkit/klt-cert-manager digest to f3bbb96 ([#2342](https://github.com/keptn/lifecycle-toolkit/issues/2342)) ([89ddb2f](https://github.com/keptn/lifecycle-toolkit/commit/89ddb2f427561bbc41ea2e4b762ac3a14aab3bc5))
* update github.com/keptn/lifecycle-toolkit/klt-cert-manager digest to fda2315 ([#2300](https://github.com/keptn/lifecycle-toolkit/issues/2300)) ([bffbaf2](https://github.com/keptn/lifecycle-toolkit/commit/bffbaf2edf4710e085db1ee956b3ccc1b6599275))
* update kubernetes packages (patch) ([#2102](https://github.com/keptn/lifecycle-toolkit/issues/2102)) ([b2853f9](https://github.com/keptn/lifecycle-toolkit/commit/b2853f9ecdfb4b7b81d0b88cf782b82c9958c5cb))
* update module github.com/argoproj/argo-rollouts to v1.6.0 ([#2064](https://github.com/keptn/lifecycle-toolkit/issues/2064)) ([d5c428a](https://github.com/keptn/lifecycle-toolkit/commit/d5c428a7e31c00362b8280da5acd91ade89c1fa8))
* update module github.com/go-logr/logr to v1.3.0 ([#2346](https://github.com/keptn/lifecycle-toolkit/issues/2346)) ([bc06204](https://github.com/keptn/lifecycle-toolkit/commit/bc06204b97c765d0f5664fd66f441af86f21e191))
* update module github.com/onsi/ginkgo/v2 to v2.12.1 ([#2156](https://github.com/keptn/lifecycle-toolkit/issues/2156)) ([dbf2867](https://github.com/keptn/lifecycle-toolkit/commit/dbf2867133067b162e82b71b6547c3dfac95d0af))
* update module github.com/onsi/ginkgo/v2 to v2.13.0 ([#2272](https://github.com/keptn/lifecycle-toolkit/issues/2272)) ([0df464d](https://github.com/keptn/lifecycle-toolkit/commit/0df464dd8e4fc7729deeb5bae4938b236902d661))
* update module github.com/onsi/gomega to v1.28.0 ([#2209](https://github.com/keptn/lifecycle-toolkit/issues/2209)) ([c0726d0](https://github.com/keptn/lifecycle-toolkit/commit/c0726d0b0e9d9732123aaf8b1ad012bc24672b84))
* update module github.com/onsi/gomega to v1.28.1 ([#2343](https://github.com/keptn/lifecycle-toolkit/issues/2343)) ([64b1508](https://github.com/keptn/lifecycle-toolkit/commit/64b1508f0e383aa7fbc406e17e2cc66546601e53))
* update module github.com/prometheus/client_golang to v1.17.0 ([#2207](https://github.com/keptn/lifecycle-toolkit/issues/2207)) ([de8b958](https://github.com/keptn/lifecycle-toolkit/commit/de8b9587cb95b1ee9c2be7a66320d284491102d9))
* update module golang.org/x/net to v0.15.0 ([#2065](https://github.com/keptn/lifecycle-toolkit/issues/2065)) ([50ce9c0](https://github.com/keptn/lifecycle-toolkit/commit/50ce9c09914f505ffaf33eee41564afa65661215))
* update module golang.org/x/net to v0.16.0 ([#2249](https://github.com/keptn/lifecycle-toolkit/issues/2249)) ([e89ea71](https://github.com/keptn/lifecycle-toolkit/commit/e89ea71bc1a2d69828179c64ffe3c34ce359dd94))
* update module golang.org/x/net to v0.17.0 ([#2267](https://github.com/keptn/lifecycle-toolkit/issues/2267)) ([8443874](https://github.com/keptn/lifecycle-toolkit/commit/8443874254cda9e5f4c662cab1a3e5e3b3277435))
* update module google.golang.org/grpc to v1.58.0 ([#2066](https://github.com/keptn/lifecycle-toolkit/issues/2066)) ([6fae5a7](https://github.com/keptn/lifecycle-toolkit/commit/6fae5a7ebf356625b4754b7890f7c71dbb4ac0a6))
* update module google.golang.org/grpc to v1.58.1 ([#2115](https://github.com/keptn/lifecycle-toolkit/issues/2115)) ([d08df40](https://github.com/keptn/lifecycle-toolkit/commit/d08df40188bc633037c49a1468a70eefc960a4a1))
* update module google.golang.org/grpc to v1.58.2 ([#2163](https://github.com/keptn/lifecycle-toolkit/issues/2163)) ([5efa650](https://github.com/keptn/lifecycle-toolkit/commit/5efa6502403daa37bdfc51fa8600da6b1f845ac2))
* update module google.golang.org/grpc to v1.58.3 ([#2275](https://github.com/keptn/lifecycle-toolkit/issues/2275)) ([66e86c0](https://github.com/keptn/lifecycle-toolkit/commit/66e86c03272d75207bd3b42014d88b1b912b9198))
* update module google.golang.org/grpc to v1.59.0 ([#2302](https://github.com/keptn/lifecycle-toolkit/issues/2302)) ([fda2315](https://github.com/keptn/lifecycle-toolkit/commit/fda231552475eaf0f60457ad42a26c4ed3473008))
* update module k8s.io/apimachinery to v0.28.3 ([#2298](https://github.com/keptn/lifecycle-toolkit/issues/2298)) ([f2f8dfe](https://github.com/keptn/lifecycle-toolkit/commit/f2f8dfec6e47517f2c476d6425c22db875f9bd3c))
* update module sigs.k8s.io/controller-runtime to v0.16.3 ([#2306](https://github.com/keptn/lifecycle-toolkit/issues/2306)) ([3d634a7](https://github.com/keptn/lifecycle-toolkit/commit/3d634a79996be6cb50805c745c51309c2f091a61))
* update opentelemetry-go monorepo (minor) ([#2108](https://github.com/keptn/lifecycle-toolkit/issues/2108)) ([4e5d29e](https://github.com/keptn/lifecycle-toolkit/commit/4e5d29e681f78590b4406ba7b74cc46ca6107e4b))
* update opentelemetry-go monorepo (minor) ([#2210](https://github.com/keptn/lifecycle-toolkit/issues/2210)) ([d577311](https://github.com/keptn/lifecycle-toolkit/commit/d5773111c327f5d30ec24437d16cf5d4454dd69e))

## [0.8.2](https://github.com/keptn/lifecycle-toolkit/compare/lifecycle-operator-v0.8.1...lifecycle-operator-v0.8.2) (2023-09-06)


### Features

* add cloud events support ([#1843](https://github.com/keptn/lifecycle-toolkit/issues/1843)) ([5b47120](https://github.com/keptn/lifecycle-toolkit/commit/5b471203e412a919903876212ac45c04f180e482))
* **lifecycle-operator:** clean up KeptnTask API by removing duplicated attributes ([#1965](https://github.com/keptn/lifecycle-toolkit/issues/1965)) ([257b220](https://github.com/keptn/lifecycle-toolkit/commit/257b220a6171ccc82d1b471002b6cf773ec9bd09))
* **metrics-operator:** add analysis controller ([#1875](https://github.com/keptn/lifecycle-toolkit/issues/1875)) ([017e08b](https://github.com/keptn/lifecycle-toolkit/commit/017e08b0a65679ca417e363f2223b7f4fef3bc55))
* **metrics-operator:** add Analysis CRD ([#1839](https://github.com/keptn/lifecycle-toolkit/issues/1839)) ([9521a16](https://github.com/keptn/lifecycle-toolkit/commit/9521a16ce4946d3169993780f2d2a4f3a75d0445))
* monorepo setup for lifecycle-operator, scheduler and runtimes ([#1857](https://github.com/keptn/lifecycle-toolkit/issues/1857)) ([84e243a](https://github.com/keptn/lifecycle-toolkit/commit/84e243a213ffba86eddd51ccc4bf4dbd61140069))


### Bug Fixes

* **lifecycle-operator:** avoid setting the overall state of an App or WorkloadInstance between state transitions ([#1871](https://github.com/keptn/lifecycle-toolkit/issues/1871)) ([ee0b085](https://github.com/keptn/lifecycle-toolkit/commit/ee0b085b05b2b9781457eba34d5d1050b3c7a604))


### Other

* **main:** release lifecycle-operator-and-scheduler libraries ([#1979](https://github.com/keptn/lifecycle-toolkit/issues/1979)) ([12d0f40](https://github.com/keptn/lifecycle-toolkit/commit/12d0f40725e466825c4a0d483fa344e5888b03ae))
* rename operator folder to lifecycle-operator ([#1819](https://github.com/keptn/lifecycle-toolkit/issues/1819)) ([97a2d25](https://github.com/keptn/lifecycle-toolkit/commit/97a2d25919c0a02165dd0dc6c7c82d57ad200139))


### Docs

* fix typos and grammar issues ([#1925](https://github.com/keptn/lifecycle-toolkit/issues/1925)) ([5570d55](https://github.com/keptn/lifecycle-toolkit/commit/5570d555bfc4bbdcbfc66b2725d5352090e5b937))
* implement KLT -&gt; Keptn name change ([#2001](https://github.com/keptn/lifecycle-toolkit/issues/2001)) ([440c308](https://github.com/keptn/lifecycle-toolkit/commit/440c3082e5400f89d791724651984ba2bc0a4724))


### Dependency Updates

* update dependency kubernetes-sigs/controller-tools to v0.13.0 ([#1930](https://github.com/keptn/lifecycle-toolkit/issues/1930)) ([8b34b63](https://github.com/keptn/lifecycle-toolkit/commit/8b34b63404d0339633ef41ff1cf2005deae8d2b7))
* update dependency kubernetes-sigs/kustomize to v5.1.1 ([#1853](https://github.com/keptn/lifecycle-toolkit/issues/1853)) ([354ab3f](https://github.com/keptn/lifecycle-toolkit/commit/354ab3f980c2569e17a0354ece417df40317d120))
* update github.com/keptn/lifecycle-toolkit/klt-cert-manager digest to 440c308 ([#2017](https://github.com/keptn/lifecycle-toolkit/issues/2017)) ([c365734](https://github.com/keptn/lifecycle-toolkit/commit/c365734fa7e3e40b2ae4c97c61628892d040dacc))
* update github.com/keptn/lifecycle-toolkit/klt-cert-manager digest to 88a54f9 ([#1794](https://github.com/keptn/lifecycle-toolkit/issues/1794)) ([fc976eb](https://github.com/keptn/lifecycle-toolkit/commit/fc976eb07ed9a5e49ed7d4ab1dbf187cee583e64))
* update github.com/keptn/lifecycle-toolkit/klt-cert-manager digest to 8dbec2d ([#1995](https://github.com/keptn/lifecycle-toolkit/issues/1995)) ([2f51445](https://github.com/keptn/lifecycle-toolkit/commit/2f5144540c4b3876e800bff29c30bfded334be40))
* update github.com/keptn/lifecycle-toolkit/klt-cert-manager digest to bb133cf ([#1963](https://github.com/keptn/lifecycle-toolkit/issues/1963)) ([c7697bf](https://github.com/keptn/lifecycle-toolkit/commit/c7697bf54d5fe18b7c62c5b11801c6c83079b0a3))
* update kubernetes packages to v0.26.8 (patch) ([#1945](https://github.com/keptn/lifecycle-toolkit/issues/1945)) ([6ce03d6](https://github.com/keptn/lifecycle-toolkit/commit/6ce03d600cbb3d3d3988573c616ec7f3830ba324))
* update module github.com/onsi/ginkgo/v2 to v2.12.0 ([#2019](https://github.com/keptn/lifecycle-toolkit/issues/2019)) ([41e878f](https://github.com/keptn/lifecycle-toolkit/commit/41e878ff8bbb438efa4b221470a571687dd392e9))
* update module github.com/onsi/gomega to v1.27.10 ([#1796](https://github.com/keptn/lifecycle-toolkit/issues/1796)) ([8f14bff](https://github.com/keptn/lifecycle-toolkit/commit/8f14bffe27485a36e0b05b770a01e357402d92f7))
* update module golang.org/x/net to v0.14.0 ([#1855](https://github.com/keptn/lifecycle-toolkit/issues/1855)) ([3186188](https://github.com/keptn/lifecycle-toolkit/commit/31861889bf7b227f489b941ac4a52db86551fcc2))
* update module google.golang.org/grpc to v1.57.0 ([#1861](https://github.com/keptn/lifecycle-toolkit/issues/1861)) ([fdcbdf5](https://github.com/keptn/lifecycle-toolkit/commit/fdcbdf50365dfd69d16c679c6814e89570a8a0e2))
* update opentelemetry-go monorepo (minor) ([#1931](https://github.com/keptn/lifecycle-toolkit/issues/1931)) ([a0a1a7e](https://github.com/keptn/lifecycle-toolkit/commit/a0a1a7e97906ab56ed85da7ab9b6d1e13c902397))


### Refactoring

* **lifecycle-operator:** eventing and telemetry ([#1844](https://github.com/keptn/lifecycle-toolkit/issues/1844)) ([0130576](https://github.com/keptn/lifecycle-toolkit/commit/0130576a17a78453019c150af849c06553d799a6))
* **lifecycle-operator:** refactor event emitter unit tests ([#1867](https://github.com/keptn/lifecycle-toolkit/issues/1867)) ([2558f74](https://github.com/keptn/lifecycle-toolkit/commit/2558f742031e4d38a8006ce9894f41bddac7cd3f))
