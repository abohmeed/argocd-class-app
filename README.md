# argocd-class-app

The **app repository** for *ArgoCD 3 in Production: GitOps at Scale on Kubernetes*.

Source and CI live here. Manifests live in
[`argocd-class-resources`](https://github.com/abohmeed/argocd-class-resources), and
per-environment values in
[`argocd-class-values`](https://github.com/abohmeed/argocd-class-values). Three repositories,
three different reasons to change — and that separation is the subject of **S11 L02**.

## Why this is not in the config repository

If the pipeline that builds this image also committed the new tag back into *this* repo, that
commit would trigger the pipeline again, which would build again, which would commit again.
The loop is not hypothetical; it is the most common way a first GitOps pipeline eats itself.
The manifest bump goes into the **config** repository, which has no build workflow watching it,
so nothing retriggers.

Argo CD never looks at this repository at all. It watches the config repo.

## The service

About forty lines: it serves one line of text, read from `BANNER`, on `:5678`. The port matches
`hashicorp/http-echo` deliberately, so this image drops into the existing manifests with no other
change. `BANNER` is read **once at process start**, which is the behaviour S01 L04 is built on —
edit the ConfigMap and nothing moves until the pod restarts.

It is deliberately boring. Its job is to have a test that can fail and an image that can ship.

## The pipeline

`.github/workflows/ci.yml` runs `go vet` and `go test`, and **only then** builds and pushes to
`ghcr.io/<owner>/argocd-class-app:<sha>`.

The line that matters is `needs: test` on the image job. Without it the two jobs race, and a red
test ships anyway. S11 L03 breaks a test on camera to show the push not happening — verified: a
failing test stops the run before the build step is ever reached.

Images are tagged by commit SHA, never `latest`. Argo CD Image Updater in S11 L05 needs an
immutable tag it can order, and `latest` is neither.

## Running the tests

```bash
go vet ./...
go test -v ./...
```
