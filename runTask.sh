fly execute -c ./ci/scripts/build/task.yml --i repo=./ --privileged
fly execute -c ./ci/scripts/test/task.yml --i repo=./ --privileged --i binaries=./binaries cf_password=____ binary-name=wildcard binary-version=0.0.9 binary-bucket="cf-wildcard-plugin" cf_api_endpoint=api.run.pivotal.io cf_user=ejung@pivotal.io cf_org=platform-eng cf_space=Jeanie cf_password=____ binary-name=wildcard binary-version=0.0.9 binary-bucket="cf-wildcard-plugin" cf_api_endpoint=api.run.pivotal.io cf_user=ejung@pivotal.io cf_org=platform-eng cf_space=Jeanie 

