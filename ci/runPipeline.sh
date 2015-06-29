echo "Run the tasks first to test against local resources before running pipeline"
echo ""
echo "Edit the ci/config/default.yml and use another file to store the secret credentials so it does not get checked into git"
echo ""
echo "Sample config in ~/private.yml can set values for"
echo ""
echo "release-access-key: ..."
echo "release-secret-key: ..."

fly configure -c scripts/pipeline.yml \
       --vars-from config/default.yml \
       --vars-from ~/private.yml        \
       wildcard_plugin
