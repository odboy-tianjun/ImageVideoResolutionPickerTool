rm -f kenaito-media-picker.exe
rm -f kenaito-media-picker-release.exe
# build
go build -o kenaito-media-picker-release.exe main.go
# upx compress
./upx -o kenaito-media-picker.exe kenaito-media-picker-release.exe
rm -f kenaito-media-picker-release.exe
