rm -f ImageVideoResolutionPickerTool.exe
rm -f ImageVideoResolutionPickerToolRelease.exe
# build
go build -o ImageVideoResolutionPickerTool.exe main.go
# upx compress
./upx -o ImageVideoResolutionPickerToolRelease.exe ImageVideoResolutionPickerTool.exe
rm -f ImageVideoResolutionPickerTool.exe
