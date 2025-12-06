ARG VERSION=latest
FROM google/cloud-sdk:${VERSION}

RUN gcloud alpha interactive

CMD ["cat", "/root/.config/gcloud/cli/gcloud.json"]
