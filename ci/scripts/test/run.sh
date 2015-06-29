echo $PWD



TASK_ROOT_DIR=$PWD
SCRIPT_DIR=$(dirname $0)

CF_API_ENDPOINT=${CF_API_ENDPOINT:-api.10.244.0.34.xip.io}
CF_USER=${CF_USER:-admin}
CF_PASSWD=${CF_PASSWD:-admin}
CF_ORG=${CF_ORG:-dev}
CF_SPACE=${CF_SPACE:-dev}

# Run as privileged
apt-get install -y wget unzip gzip

echo "Install CF binary"
cd /tmp
wget "https://cli.run.pivotal.io/stable?release=linux64-binary&version=6.11.3&source=github-rel" -O cf-linux-amd64.tgz

tar zxvf cf-linux-amd64.tgz
ls
mkdir -p /usr/local/bin
cp ./cf /usr/local/bin
export PATH=$PATH:/usr/local/bin

cf -version

echo "Finished installing CF binary!!"


echo "Installing Plugin"
ls $TASK_ROOT_DIR/*
cd $TASK_ROOT_DIR/binaries

tar zxvf *.tgz
tar tvf *.tgz

# Going to only test the linux-64 bit version...
plugin_binary=`echo $PWD/bin/linux/amd64/*`
cf install-plugin $plugin_binary

echo "Logging into CF Api endpoint: $CF_API_ENDPOINT"
cf api $CF_API_ENDPOINT --skip-ssl-validation
cf login -u $CF_USER -p $CF_PASSWD -o $CF_ORG -s $CF_SPACE
#cf target -o $CF_ORG -s $CF_SPACE

echo "Logged into CF Api endpoint: $CF_API_ENDPOINT"

echo "Pushing a Test App"

#cf push ...# Push the app

echo "Testing the Plugin"
# Add code to run the test...
