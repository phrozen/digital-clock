

build-darwin:
    mkdir -p build/darwin/digital-clock.app/Contents/MacOS
    mkdir -p tmp/Icon.iconset
    sips -z 16 16 icon.png --out tmp/Icon.iconset/icon_16x16.png
    sips -z 32 32 icon.png --out tmp/Icon.iconset/icon_16x16@2x.png
    sips -z 32 32 icon.png --out tmp/Icon.iconset/icon_32x32.png
    sips -z 64 64 icon.png --out tmp/Icon.iconset/icon_32x32@2x.png
    sips -z 128 128 icon.png --out tmp/Icon.iconset/icon_128x128.png
    sips -z 256 256 icon.png --out tmp/Icon.iconset/icon_128x128@2x.png
    sips -z 256 256 icon.png --out tmp/Icon.iconset/icon_256x256.png
    sips -z 512 512 icon.png --out tmp/Icon.iconset/icon_256x256@2x.png
    sips -z 512 512 icon.png --out tmp/Icon.iconset/icon_512x512.png
    iconutil -c icns tmp/Icon.iconset
    mkdir build/darwin/digital-clock.app/Contents/Resources
    mv tmp/Icon.icns build/darwin/digital-clock.app/Contents/Resources
    rm -rf tmp
    touch build/darwin/digital-clock.app/Contents/Info.plist
    go build -o build/darwin/digital-clock.app/Contents/MacOS
    

