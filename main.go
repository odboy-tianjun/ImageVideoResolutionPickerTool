package main

import (
	"ImageVideoResolutionPickerTool/core"
	"ImageVideoResolutionPickerTool/vars"
	_ "embed"
	"fmt"
	_ "image/gif"  // 导入gif支持
	_ "image/jpeg" // 导入jpeg支持
	_ "image/png"  // 导入png支持
	"os"
	"time"
)

func main() {
	fmt.Println("=== 警告：该程序会改变目录结构，将在5秒后开始执行")
	time.Sleep(time.Second * 5)
	rootDir, err := os.Getwd()
	if err != nil {
		fmt.Println("=== 获取当前路径异常", err)
		return
	}
	rootDir = "E:\\DGD\\d20250424\\video"
	scanner := core.FileScanner{}
	scanner.DoScan(rootDir)
	scanner.DoFilter()
	if len(vars.GlobalImagePathList) > 0 {
		core.DoHandleImage(rootDir)
	}
	if len(vars.GlobalVideoPathList) > 0 {
		core.DoHandleVideo(rootDir)
	}
}
