#!/bin/bash

# build BE.
cd be
./build.sh
cd ..

# build FE
cd fe/quiz
./build.sh