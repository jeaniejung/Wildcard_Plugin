# Run Tasks on Concourse

## Run the individual tasks using local resources

Note: Install fly or download it from concourse ui page and configure it against a running concourse instance before proceeding with next steps

* Running tasks using fly would use local resources as opposed to use real (git or s3) for pipeline.
* Run the build task first to build and save the wildcard plugin
* To check the status of the run, access the Concourse page (if running a local instance on Vagrant, the address is http://192.168.100.4:8080) and click on the Folders/drawers icon on the right hand top left corner and click on one-off tasks recently kicked off.


## Running the build task

'build' task requires as input the repository path in form of repo variable

```
fly execute -c ./scripts/build/task.yml --i repo=../ --privileged
```

Save the tar ball (created inside the docker image if necessary) under <repo>/binaries folder either in s3 or locally from the build task

## Running the test task

Run the 'test' task once build task is successful and ensure there is a tar ball containing the plugin executable for each platform under 'binaries'

```
# Execute this from the <Repository>/ci folder so the repo is the parent folder of current working directory.
fly execute -c ./scripts/test/task.yml --i repo=../ --privileged \\
                                          --i binaries=../binaries

```

There can be additional overrides of the cf and other parameters
Example: 

```
# Execute this from the <Repository>/ci folder so the repo is the parent folder of current working directory.
fly execute -c ./scripts/test/task.yml                   \\
                      --privileged                       \\
                      --i repo=../                       \\
                      --i binaries=../binaries           \\
                      binary-name=wildcard               \\
                      binary-version=0.0.9               \\
                      binary-bucket=cf-wildcard-plugin   \\
                      cf_api_endpoint=api.run.pivotal.io \\
                      cf_user=testuser@xyz.com           \\
                      cf_password=samplexyz              \\
                      cf_org=test-org                    \\
                      cf_space=test-space 
```

