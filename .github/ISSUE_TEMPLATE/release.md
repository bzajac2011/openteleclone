---
name: Release
about: Create an issue for releasing new versions
title: 'Release vX.X.X'
labels: release
assignees: ''

---

Like #4522, but for vX.X.X

**Performed by collector release manager**

- [ ] Prepare core release vX.X.X
- [ ] Tag and release core vX.X.X
- [ ] Prepare contrib release vX.X.X
- [ ] Tag and release contrib vX.X.X
- [ ] Prepare otelcol-releases vX.X.X
- [ ] Release binaries and container images vX.X.X

**Performed by operator maintainers**

- [ ] Release the operator vX.X.X

**Performed by helm chart maintainers**

- [ ] Update the opentelemetry-collector helm chart to use vX.X.X docker image
