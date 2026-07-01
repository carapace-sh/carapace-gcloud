ARG VERSION=575.0.0
FROM google/cloud-sdk:${VERSION}

RUN gcloud alpha interactive; true

CMD ["cat", "/root/.config/gcloud/cli/gcloud.json"]
