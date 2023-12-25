#!/bin/bash

VERSION="${1:-2.1}"

cleanup() {
  echo "Cleaning up..."
  rm -f xlink.xsd
  rm -f MTConnectAssets.xsd
  rm -f MTConnectDevices.xsd
  rm -f MTConnectError.xsd
  rm -f MTConnectStreams.xsd
  rm -rf schema/v$VERSION/mtassets/mtassets
  rm -rf schema/v$VERSION/mtdevices/mtdevices
  echo "Cleanup done."
}

trap cleanup EXIT

if type xgen >/dev/null 2>&1; then
    echo "xgen is available."
else
    echo "xgen is not available. install curl."
    go install github.com/xuri/xgen/cmd/xgen@latest
fi


curl https://raw.githubusercontent.com/mtconnect/schema/master/xlink.xsd -o xlink.xsd

curl https://raw.githubusercontent.com/mtconnect/schema/master/MTConnectAssets_$VERSION.xsd -o MTConnectAssets.xsd
xgen -i MTConnectAssets.xsd -l Go -o schema/v$VERSION/mtassets/mtassets -p mtassets

curl https://raw.githubusercontent.com/mtconnect/schema/master/MTConnectDevices_$VERSION.xsd -o MTConnectDevices.xsd
xgen -i MTConnectDevices.xsd -l Go -o schema/v$VERSION/mtdevices/mtdevices -p mtdevices

curl https://raw.githubusercontent.com/mtconnect/schema/master/MTConnectError_$VERSION.xsd -o MTConnectError.xsd
xgen -i MTConnectError.xsd -l Go -o schema/v$VERSION/mterror/mterror -p mterror

curl https://raw.githubusercontent.com/mtconnect/schema/master/MTConnectStreams_$VERSION.xsd -o MTConnectStreams.xsd
xgen -i MTConnectStreams.xsd -l Go -o schema/v$VERSION/mtstreams/mtstreams -p mtstreams

echo "Script finished."
