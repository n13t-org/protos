#!/usr/bin/env bash

#
# DEPRECATED
#

# Save all params for later use.
PARAMS="$@"

# Capture current proto dir
while test $# -gt 0; do
    case "$1" in
        -d)
            shift
            if test $# -gt 0; then
                PROTO_DIR=$1
            else
                echo "no directory specified"
                exit 1
            fi
            shift
            ;;
        *)
            break;
            ;;
    esac
done

# Copy all protos dir to /opt/include to support relative import. Ex: import "gitlab/v4/event.proto"
#mkdir -p /opt/include/n13t/
#ln -s /defs /opt/include/n13t
cp -R /defs /opt/include/n13t
# Remove the current proto dir to avoid conflict errors.
#rm /opt/include/$PROTO_DIR

# Call the original entrypoint to process
/usr/local/bin/entrypoint.sh $PARAMS
