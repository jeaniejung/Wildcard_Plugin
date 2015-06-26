fly execute -c ./ci/scripts/build/task.yml --i repo=./ --privileged
fly execute -c ./ci/scripts/test/task.yml --i repo=./ --privileged --i binaries=./binaries
