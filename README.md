# Wildcard Plugin
This CF CLI Plugin allows users to search through and delete their applications using wildcards. It is useful for users who have multiple apps to manage in their spaces. 

#Requirements
To prevent your shell from expanding the wildcard before the plugin sees it, wildcards should be escaped using a preceding '\\'.
```
$ cf wc-d app\* 
$ cf wc-a app\?
```
# Installation

#### Install from CLI (v.6.10.0 and up)
```
$ cf add-plugin-repo CF-Community http://plugins.cloudfoundry.org/
$ cf install-plugin wildcard_plugin -r CF-Community
```

[//]: # (#### Install from binary)
[//]: # (- Download the appropriate plugin binary from [releases](https://github.com/swisscom/cf-statistics-plugin/releases))
[//]: # (- Install the plugin: `$ cf install-plugin <binary>`)

#### Install from Source
```
$ go get github.com/cloudfoundry/cli
$ go get github.com/jeaniejung/Wildcard_Plugin
$ cd $GOPATH/src/github.com/jeaniejung/Wildcard_Plugin
$ go build *.go
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
$ cf uninstall-plugin wildcard
```
## Commands for wildcard-apps, wc-a

| command/option | usage | description|
| :--------------- |:---------------| :------------|
|`wildcard-apps, wc-a`| `cf wc-a APP_NAME_WITH_WILDCARD` |List all apps in the target space matching the wildcard pattern|

## Commands for wildcard-delete, wc-d

| command/option | usage | description|
| :--------------- |:---------------| :------------|
|`wildcard-delete, wc-d`| `cf wc-d APP_NAME_WITH_WILDCARD` |Displays list of matched apps and prompts the user for interactive deletion or force deletion of all matched apps|
|`-r`|`cf wc-d APP_NAME_WITH_WILDCARD -r`|Displays list of matched apps and prompts the user for interactive deletion or force deletion of all matched apps and their routes|
|`-f`|`cf wc-d APP_NAME_WITH_WILDCARD -f`|Force deletion of all apps in the target space matching the wildcard pattern without confirmation|
|`-f -r`|`cf wc-d APP_NAME_WITH_WILDCARD -f -r`|Force deletion of all apps and their mapped routes in the target space matching the wildcard pattern without confirmation|


