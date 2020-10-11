#!/bin/bash

function install_dependencies {
	echo Installing dependencies...
	go get -v github.com/spf13/cobra/cobra
	echo cobra
	go get -v github.com/jedib0t/go-pretty/table
	echo go-pretty
	go get -v github.com/rocketlaunchr/google-search
	echo google-search
	go get -v github.com/pkg/browser
	echo browser

}

function create_configs {
	mkdir $HOME/.tt
	touch $HOME/.tt/tasks.json
	touch $HOME/.tt/completed.json
}

install_dependencies
create_configs