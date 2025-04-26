package util

import (
	"github.com/gabriel-vasile/mimetype"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

// CreateDir 创建目录
func CreateDir(dirPath string) bool {
	err := os.Mkdir(dirPath, 0755)
	return err == nil
}

// GetFileDirectory 获取文件所在文件夹
func GetFileDirectory(filePath string) string {
	directoryPath, _ := filepath.Split(filePath)
	return directoryPath
}

// ReadFileMimeInfo 获取文件mime信息
func ReadFileMimeInfo(filepath string) *mimetype.MIME {
	mt, err := mimetype.DetectFile(filepath)
	if err != nil {
		log.Fatal(err)
	}
	return mt
}

// CheckFileIsExist 判断文件是否存在，存在返回 true，不存在返回false
func CheckFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

// WriteByteArraysToFile 写字节数组到文件
func WriteByteArraysToFile(content []byte, filename string) error {
	return ioutil.WriteFile(filename, content, 0777)
}

//// 移动文件到根目录
//func renameFile(rootDir string, modelType string, videoList []string, pathSeparator string) {
//	total := len(videoList)
//	var count = 0
//	bar := goPrint.NewBar(100)
//	bar.SetNotice("=== 重命名文件：")
//	bar.SetGraph(">")
//	for _, videoFilePath := range videoList {
//		wh := videoPath2WidthHeightTagMap[videoFilePath]
//		fileName := vars.GlobalFilePath2FileNameMap[videoFilePath]
//		if strings.Contains(fileName, videoTag) { // 处理过了
//			fileNames := strings.Split(fileName, videoTag)
//			if len(fileNames) == 2 {
//				fileName = fileNames[1]
//				targetFilePath := rootDir + pathSeparator + "[" + videoPath2DurationMap[videoFilePath] + "]" + modelType + wh + videoTag + fileName
//				err := os.Rename(videoFilePath, targetFilePath)
//				if err != nil {
//					fmt.Printf("=== 重命名异常: %s \n", videoFilePath)
//				}
//			}
//		} else {
//			targetFilePath := rootDir + pathSeparator + "[" + videoPath2DurationMap[videoFilePath] + "]" + modelType + wh + videoTag + " - " + fileName
//			err := os.Rename(videoFilePath, targetFilePath)
//			if err != nil {
//				fmt.Printf("=== 重命名异常: %s \n", videoFilePath)
//			}
//		}
//		count = count + 1
//		bar.PrintBar(util.CalcPercentage(count, total))
//	}
//	bar.PrintEnd("=== Finish")
//}
//
//// 移动文件到原目录
//func renameFileV2(modelType string, videoList []string) {
//	total := len(videoList)
//	var count = 0
//	bar := goPrint.NewBar(100)
//	bar.SetNotice("=== 重命名文件：")
//	bar.SetGraph(">")
//	for _, videoFilePath := range videoList {
//		wh := videoPath2WidthHeightTagMap[videoFilePath]
//		fileName := vars.GlobalFilePath2FileNameMap[videoFilePath]
//		filePath := util.GetFileDirectory(videoFilePath)
//		if strings.Contains(fileName, videoTag) { // 处理过了
//			fileNames := strings.Split(fileName, videoTag)
//			if len(fileNames) == 2 {
//				fileName = fileNames[1]
//				targetFilePath := filePath + "[" + videoPath2DurationMap[videoFilePath] + "]" + modelType + wh + videoTag + fileName
//				err := os.Rename(videoFilePath, targetFilePath)
//				if err != nil {
//					fmt.Printf("=== 重命名异常: %s \n", videoFilePath)
//				}
//			}
//		} else {
//			targetFilePath := filePath + "[" + videoPath2DurationMap[videoFilePath] + "]" + modelType + wh + videoTag + " - " + fileName
//			err := os.Rename(videoFilePath, targetFilePath)
//			if err != nil {
//				fmt.Printf("=== 重命名异常: %s \n", videoFilePath)
//			}
//		}
//		count = count + 1
//		bar.PrintBar(util.CalcPercentage(count, total))
//	}
//	bar.PrintEnd("=== Finish")
//}
