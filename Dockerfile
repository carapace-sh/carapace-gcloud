FROM google/cloud-sdk:583.0.0

RUN gcloud alpha interactive; true

CMD ["cat", "/root/.config/gcloud/cli/gcloud.json"]
