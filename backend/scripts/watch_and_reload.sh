#!/usr/bin/env bash
trap cleanup SIGINT

killgo(){
    # Kill the process when the script ends
    ps ax | grep ${GO_EXECUTABLE} | grep -v grep | awk '{print $1}' | xargs kill
}

cleanup(){
    killgo
    exit
}

# Path to the Go executable
GO_EXECUTABLE="./bin/main"

# Start the application initially
$GO_EXECUTABLE &

# Monitor changes to the Go executable
while inotifywait -e attrib "$GO_EXECUTABLE";
do
    killgo
    $GO_EXECUTABLE &
done;


