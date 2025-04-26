package util

import (
	"bytes"
	"fmt"
	"github.com/redmask-hb/GoSimplePrint/goPrint"
	"math"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

var supportVideoTypes = []string{
	".ts",
	".flv",
	".rm",
	".avi",
	".mp4",
	".mov",
	".mpg",
	".mkv",
	".m4v",
	".rmvb",
	".3gp",
	".3g2",
	".webm",
	".wmv",
}

func getVideoDuration(ffmpegExecPath string, videoPath string) (float64, error) {
	// ffprobe命令，-v error 用于减少输出信息，-show_entries format=duration -of compact=p=0,nk=1 用于只输出时长
	cmd := exec.Command(ffmpegExecPath, "-v", "error", "-show_entries", "format=duration", "-of", "default=noprint_wrappers=1:nokey=1", videoPath)
	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	output, err := cmd.Output()
	if err != nil {
		return 0, fmt.Errorf("ffprobe failed with error: %v, stderr: %q", err, stderr.String())
	}

	// 解析输出的时长字符串为浮点数
	durationStr := strings.TrimSpace(string(output))
	duration, err := strconv.ParseFloat(durationStr, 64)
	if err != nil {
		return 0, fmt.Errorf("failed to parse duration: %q, error: %v", durationStr, err)
	}

	return duration, nil
}

func getVideoResolution(ffmpegExecPath string, filePath string) (width int, height int, err error) {
	// 构建ffprobe命令
	cmd := exec.Command(ffmpegExecPath, "-v", "error", "-show_entries", "stream=width,height", "-of", "csv=p=0:s=x", filePath)
	// 执行命令并捕获输出
	output, err := cmd.Output()
	if err != nil {
		return 0, 0, fmt.Errorf("failed to run ffprobe: %w", err)
	}
	// 解析输出字符串，格式应为 "宽度,高度"
	resolutionStr := strings.TrimSpace(string(output))
	parts := strings.Split(resolutionStr, ",")
	if len(parts) == 2 {
		width = String2int(parts[0])
		height = String2int(parts[1])
		return width, height, nil
	}
	parts = strings.Split(resolutionStr, "x")
	if len(parts) == 2 {
		width = String2int(parts[0])
		height = String2int(parts[1])
		return width, height, nil
	}
	parts = strings.Split(resolutionStr, "\r\n\r\n\r\n\r\n")
	if len(parts) == 2 {
		tempHw := parts[0]
		parts = strings.Split(tempHw, "x")
		if len(parts) == 2 {
			width = String2int(parts[0])
			height = String2int(parts[1])
			return width, height, nil
		}
	}
	parts = strings.Split(resolutionStr, "x")
	if len(parts) == 3 {
		width = String2int(parts[0])
		height = String2int(parts[1])
		return width, height, nil
	}
	resolutionStr = strings.ReplaceAll(resolutionStr, "\r", "")
	resolutionStr = strings.ReplaceAll(resolutionStr, "\n", "")
	resolutionStr = strings.ReplaceAll(resolutionStr, "N/AxN/A", "")
	parts = strings.Split(resolutionStr, "x")
	if len(parts) == 2 {
		width = String2int(parts[0])
		height = String2int(parts[1])
		return width, height, nil
	}
	return 0, 0, fmt.Errorf("invalid resolution format: %s", resolutionStr)
}

// ReadVideoDuration 获取视频的时长，单位秒
func ReadVideoDuration(videoFilePath string) int {
	duration, err := getVideoDuration("./ffprobe.exe", videoFilePath)
	if err != nil {
		fmt.Println("=== Error getting video duration:", err)
		return 0
	}
	//fmt.Printf("=== Video duration: %.2f seconds\n", duration)
	return int(math.Floor(duration)) // 向下取整
}

// ReadVideoWidthHeight 获取视频的分辨率
func ReadVideoWidthHeight(videoFilePath string) (int, int, error) {
	width, height, err := getVideoResolution("./ffprobe.exe", videoFilePath)
	if err != nil {
		fmt.Printf("=== Error getting resolution: %v\n", err)
		return 0, 0, err
	}
	//fmt.Printf("=== Video resolution: %dx%d\n", width, height)
	return width, height, nil
}

// IsSupportVideo 判断是否属于支持的视频
func IsSupportVideo(videoType string) bool {
	for _, supportVideoType := range supportVideoTypes {
		if strings.EqualFold(videoType, supportVideoType) {
			return true
		}
	}
	return false
}

// DoMoveFileToDir 批量移动文件到目录
func DoMoveFileToDir(filePathList []string, videoDirPath string) {
	total := len(filePathList)
	var count = 0
	bar := goPrint.NewBar(100)
	bar.SetNotice("=== 移动文件到目录：")
	bar.SetGraph(">")
	pathSeparator := string(os.PathSeparator)
	for _, videoFilePath := range filePathList {
		moveFileToDir(videoFilePath, videoDirPath+pathSeparator)
		count = count + 1
		bar.PrintBar(CalcPercentage(count, total))
	}
	bar.PrintEnd("=== Finish")
}

// 移动文件到目录
func moveFileToDir(sourceFilePath string, targetDirectory string) bool {
	splits := strings.Split(sourceFilePath, string(os.PathSeparator))
	fileName := splits[len(splits)-1]
	targetFilePath := targetDirectory + fileName
	err := os.Rename(sourceFilePath, targetFilePath)
	//fmt.Printf("=== 移动文件, 源: %s, 目标: %s \n", sourceFilePath, targetFilePath)
	return err == nil
}
