#!/bin/bash

# base64(hello, world!) = aGVsbG8sIHdvcmxkIQ==
PUSH_FORMAT='{
    "receiver":"service-node",
    "req_data":{
        "method":"POST",
        "host":"hidden-echo-service",
        "path":"/echo",
        "body":"aGVsbG8sIHdvcmxkIQ=="
    }
}';