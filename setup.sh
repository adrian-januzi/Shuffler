#!/bin/bash
go build -o shuffler
sudo mv shuffler /usr/local/bin

SRC_DIR="configs"
DEST_DIR="$HOME/.shuffler"

mkdir -p "$DEST_DIR"
cp -r "$SRC_DIR"/* "$DEST_DIR"