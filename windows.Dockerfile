ARG OSVERSION
ARG BUILDERIMAGE="golang:1.16"
FROM --platform=linux/amd64 gcr.io/k8s-staging-e2e-test-images/windows-servercore-cache:1.0-linux-amd64-${OSVERSION} as core

FROM mcr.microsoft.com/windows/nanoserver:${OSVERSION}
COPY ./bin/nmi.exe /nmi.exe
COPY --from=core /Windows/System32/netapi32.dll /Windows/System32/netapi32.dll

USER ContainerAdministrator
LABEL description="AAD Pod Identity"

ENTRYPOINT ["/nmi.exe"]
