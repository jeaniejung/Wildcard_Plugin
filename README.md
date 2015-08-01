# Wildcard Plugin
This CF CLI Plugin allows users to search through and delete their applications using wildcards. It is useful for users who have multiple apps to manage in their spaces. 

#Requirements
Users must disable their wildcard expansion functionality in order for the plugin to run correctly. This can be done by running the following command. 
```
$ set -f
```
The wildcard expansion functionality can be re-enabled by the following command.
```
$ set +f
```

# Installation

#### Install from CLI (v.6.10.0 and up)
```
$ cf add-plugin-repo CF-Community http://plugins.cloudfoundry.org/
$ cf install-plugin Wildcard_Plugin -r CF-Community
```
  
#### Install from binary
- Download the appropriate plugin binary from [releases](https://github.com/swisscom/cf-statistics-plugin/releases)
- Install the plugin: `$ cf install-plugin <binary>`

#### Install from Source
```
$ go get github.com/cloudfoundry/cli
$ go get github.com/jeaniejung/Wildcard_Plugin
$ go build $GOPATH/src/github.com/jeaniejung/Wildcard_Plugin/*.go
$ cf install-plugin $GOPATH/bin/Wildcard_Plugin/wildcard_plugin
```

## Usage

```
$ cf wildcard-apps APP_NAME_WITH_WILDCARD
```
```
$ cf wildcard-delete APP_NAME_WITH_WILDCARD [-f -r]
```

## Uninstall

```
$ cf uninstall-plugin wildcard_plugin
```
## Commands for wildcard-apps, wc-a

| command/option | usage | description|
| :--------------- |:---------------| :------------|
|`wildcard-apps, wc-a`| `cf wc-a APP_NAME_WITH_WILDCARD` |Displays list of matched apps in current space|

## Commands for wildcard-delete, wc-d

| command/option | usage | description|
| :--------------- |:---------------| :------------|
|`wildcard-delete, wc-d`| `cf wc-d APP_NAME_WITH_WILDCARD` |Displays list of matched apps and prompts the user for interctive deletion or force deletion of all matched apps|
|`-f`|`cf wc-d APP_NAME_WITH_WILDCARD -f`|Force deletion of all apps that match the pattern of the wildcard without confirmation|
|`-r`|`cf wc-d APP_NAME_WITH_WILDCARD -r`|Force deletion of the routes of all apps that match the pattern of the wildcard without confirmation|
|`-f -r`|`cf wc-d APP_NAME_WITH_WILDCARD -f -r`|Force deletion of all apps and their mapped routes that match the pattern of the wildcard without confirmation|


