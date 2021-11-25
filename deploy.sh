#!/bin/bash

git pull http://mri_administrator:mamangsekayu97@180.211.92.131:7990/scm/bo/client-email-sender.git master
echo Yono@MR1 | sudo docker-compose up -d --build