#!/bin/bash

function install_dependencies {
	echo Installing dependencies...
	go get -v github.com/spf13/cobra/cobra
	echo cobra
}

function create_configs {
	mkdir $HOME/.tt
	touch $HOME/.tt/tasks.json
	touch $HOME/.tt/completed.json
}

install_dependencies
create_configs
