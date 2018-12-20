#!/bin/sh
curl http://files.grouplens.org/papers/ml-100k.zip -o "./data/ml-100k.zip"
unzip ./data/ml-100k.zip -d data