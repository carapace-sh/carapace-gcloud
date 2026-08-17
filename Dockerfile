FROM google/cloud-sdk:580.0.0

RUN gcloud alpha interactive; true

CMD ["cat", "/root/.config/gcloud/cli/gcloud.json"]
