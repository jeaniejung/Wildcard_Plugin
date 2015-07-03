# Wildcard Plugin
This CF CLI Plugin allows users to search through and delete their applications using wildcards. It is useful for users with multiple apps in their spaces. 
This CF CLI Plugin to rolls traffic from one app to another over a specified time interval. It is useful for blue green deployments or other situations where start / stop is not enough or too abrupt. It was created by Guido Westenberg and Josh Kruck after being asked by too many people if CF had this functionality and having to say no. 

#Requirements
Users must disable their wildcard expansion functionality in order for the plugin to run correctly. This can be done by running the command 'set -f'. The wildcard expansion functionality can be re-enabled by the command 'set +f' 

#Assumptions
This plugin makes no assumptions besides the disablement of the wildcard expansion functionality mentioned under Reguirements
