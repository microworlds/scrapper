# Scrapper
![Scrapper Demo Pic](https://www.github.com/microworlds/example/scrapper.png)

## Overview
Scrapper is a little software that lets you scrape all the links in a website and provide mock `json` data for your application. You give it a url to scrape with the `-u` flag and a filename to save the data using the `-f` flag. The app stores the file in the same directory where you run it. It uses `http://www.bbc.com` as default url if not provided and `data.json` as default output destination

## Usage 
`./scrapper -u "http://www.example.com" -f data`

Run `./scrapper --help` or `./scrapper -h` to see the commandline flags and their uses


## Platforms
This software is compatible with Windows, Linux and Mac OS. The binaries are in the `bin` folder of this repository. You can also build from source using the `make compile` command (on Unix systems)

## Todo
Make it a rock solid web scrapper