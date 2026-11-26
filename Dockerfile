FROM google/cloud-sdk:594.0.1

RUN gcloud alpha interactive; true

CMD ["cat", "/root/.config/gcloud/cli/gcloud.json"]
