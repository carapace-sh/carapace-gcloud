FROM google/cloud-sdk:590.0.0

RUN gcloud alpha interactive; true

CMD ["cat", "/root/.config/gcloud/cli/gcloud.json"]
