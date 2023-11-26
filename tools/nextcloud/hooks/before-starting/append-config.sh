#!/bin/bash
#
# Remove existing custom config
sed -i '/^$CUSTOM_CONFIG/,$d' /var/www/html/config/config.php
# Appends config.php
cat /var/www/html/config/config.custom.php >> /var/www/html/config/config.php
