package core

import (
	"ImageVideoResolutionPickerTool/util"
	"ImageVideoResolutionPickerTool/vars"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"os"
	"strings"
	"time"
)

func DoHandleVideo(rootDir string) {
	// 释放ffprobe
	readerFileName := "./ffprobe.exe"
	if util.CheckFileIsExist(readerFileName) {
		_ = os.Remove(readerFileName)
	}
	err := util.WriteByteArraysToFile(vars.VideoDecodeToolWin64, readerFileName)
	if err != nil {
		fmt.Println("=== 释放解码器失败, 5秒后将自动退出", err)
		time.Sleep(time.Second * 5)
		return
	}
	total := len(vars.GlobalVideoPathList) // 总数
	successCount := 0                      // 成功数
	errorCount := 0                        // 失败数
	ignoreCount := 0                       // 忽略数
	for _, videoFilePath := range vars.GlobalVideoPathList {
		suffix := vars.GlobalFilePath2FileExtMap[videoFilePath]
		if util.IsSupportVideo(suffix) {
			width, height, err := util.ReadVideoWidthHeight(videoFilePath)
			if err == nil {
				successCount = successCount + 1
				vars.VideoPath2WidthHeightMap[videoFilePath] = fmt.Sprintf("%d-%d", width, height)
				vars.VideoPath2WidthHeightTagMap[videoFilePath] = fmt.Sprintf("[%dx%d]", width, height)
				fmt.Printf("=== Video总数: %d, 已读取Info: %d, 成功数: %d, 失败数: %d \n", total, successCount+errorCount+ignoreCount, successCount, errorCount)
				duration := util.ReadVideoDuration(videoFilePath)
				if duration == 0 {
					vars.VideoPath2DurationMap[videoFilePath] = "0H0M0S"
				} else {
					vars.VideoPath2DurationMap[videoFilePath] = util.SecondsToHms(duration)
				}
			} else {
				errorCount = errorCount + 1
				vars.VideoReadErrorPathList = append(vars.VideoReadErrorPathList, videoFilePath)
				fmt.Printf("=== 异常视频: %s \n", videoFilePath)
			}
			continue
		}
		// 其他的直接先忽略吧, 爱改改, 不改拉倒
		ignoreCount = ignoreCount + 1
		vars.VideoIgnorePathList = append(vars.VideoIgnorePathList, videoFilePath)
	}
	if len(vars.VideoReadErrorPathList) > 0 {
		readInfoErrorPath := rootDir + string(os.PathSeparator) + "读取异常"
		if util.CreateDir(readInfoErrorPath) {
			util.DoMoveFileToDir(vars.VideoReadErrorPathList, readInfoErrorPath)
		}
	}
	if len(vars.VideoIgnorePathList) > 0 {
		ignorePath := rootDir + string(os.PathSeparator) + "已忽略"
		if util.CreateDir(ignorePath) {
			util.DoMoveFileToDir(vars.VideoIgnorePathList, ignorePath)
		}
	}
	doPickVideoFile(rootDir, vars.VideoPath2WidthHeightMap)
	// 删除ffprobe
	if util.CheckFileIsExist(readerFileName) {
		_ = os.Remove(readerFileName)
	}
	fmt.Printf("=== 视频处理完毕 \n\n")
}

// 条件视频并分组存放
func doPickVideoFile(rootDir string, videoPath2WidthHeightMap map[string]string) {
	if len(videoPath2WidthHeightMap) == 0 {
		fmt.Printf("=== 当前目录下没有扫描到视频文件, %s \n", rootDir)
		readerFileName := "./ffprobe.exe"
		if util.CheckFileIsExist(readerFileName) {
			_ = os.Remove(readerFileName)
		}
		return
	}
	for currentVideoPath, infoStr := range videoPath2WidthHeightMap {
		width2Height := strings.Split(infoStr, "-")
		width := util.String2int(width2Height[0])
		height := util.String2int(width2Height[1])
		suffix := vars.GlobalFilePath2FileExtMap[currentVideoPath]
		if width > height {
			handleHorizontalVideo(currentVideoPath, width, height, suffix)
			continue
		}
		if width < height {
			handleVerticalVideo(currentVideoPath, height, suffix)
			continue
		}
		handleSquareVideo(currentVideoPath, width, height, suffix)
	}
	uid := strings.ReplaceAll(uuid.NewV4().String(), "-", "")
	moveNormalVideo(uid, rootDir)
	moveHorizontalVideo(uid, rootDir)
	moveVerticalVideo(uid, rootDir)
	moveSquareVideo(uid, rootDir)
}

// 移动垂直视频
func moveVerticalVideo(uid string, rootDir string) {
	pathSeparator := string(os.PathSeparator)
	if len(vars.VideoVertical1KList) > 0 {
		videoDirPath := rootDir + pathSeparator + uid + "-视频_竖屏_1k"
		util.CreateDir(videoDirPath)
		util.DoMoveFileToDir(vars.VideoVertical1KList, videoDirPath)
	}
	if len(vars.VideoVertical2KList) > 0 {
		videoDirPath := rootDir + pathSeparator + uid + "-视频_竖屏_2k"
		util.CreateDir(videoDirPath)
		util.DoMoveFileToDir(vars.VideoVertical2KList, videoDirPath)
	}
	if len(vars.VideoVertical3KList) > 0 {
		videoDirPath := rootDir + pathSeparator + uid + "-视频_竖屏_3k"
		util.CreateDir(videoDirPath)
		util.DoMoveFileToDir(vars.VideoVertical3KList, videoDirPath)
	}
	if len(vars.VideoVertical4KList) > 0 {
		videoDirPath := rootDir + pathSeparator + uid + "-视频_竖屏_4k"
		util.CreateDir(videoDirPath)
		util.DoMoveFileToDir(vars.VideoVertical4KList, videoDirPath)
	}
	if len(vars.VideoVertical5KList) > 0 {
		videoDirPath := rootDir + pathSeparator + uid + "-视频_竖屏_5k"
		util.CreateDir(videoDirPath)
		util.DoMoveFileToDir(vars.VideoVertical5KList, videoDirPath)
	}
	if len(vars.VideoVertical6KList) > 0 {
		videoDirPath := rootDir + pathSeparator + uid + "-视频_竖屏_6k"
		util.CreateDir(videoDirPath)
		util.DoMoveFileToDir(vars.VideoVertical6KList, videoDirPath)
	}
	if len(vars.VideoVertical7KList) > 0 {
		videoDirPath := rootDir + pathSeparator + uid + "-视频_竖屏_7k"
		util.CreateDir(videoDirPath)
		util.DoMoveFileToDir(vars.VideoVertical7KList, videoDirPath)
	}
	if len(vars.VideoVertical8KList) > 0 {
		videoDirPath := rootDir + pathSeparator + uid + "-视频_竖屏_8k"
		util.CreateDir(videoDirPath)
		util.DoMoveFileToDir(vars.VideoVertical8KList, videoDirPath)
	}
	if len(vars.VideoVertical9KList) > 0 {
		videoDirPath := rootDir + pathSeparator + uid + "-视频_竖屏_9k"
		util.CreateDir(videoDirPath)
		util.DoMoveFileToDir(vars.VideoVertical9KList, videoDirPath)
	}
	if len(vars.VideoVerticalHKList) > 0 {
		videoDirPath := rootDir + pathSeparator + uid + "-视频_竖屏_原画质"
		util.CreateDir(videoDirPath)
		util.DoMoveFileToDir(vars.VideoVerticalHKList, videoDirPath)
	}
}

// 移动水平视频
func moveHorizontalVideo(uid string, rootDir string) {
	pathSeparator := string(os.PathSeparator)
	if len(vars.VideoHorizontal1KList) > 0 {
		videoDirPath := rootDir + pathSeparator + uid + "-视频_横屏_1k"
		util.CreateDir(videoDirPath)
		util.DoMoveFileToDir(vars.VideoHorizontal1KList, videoDirPath)
	}
	if len(vars.VideoHorizontal2KList) > 0 {
		videoDirPath := rootDir + pathSeparator + uid + "-视频_横屏_2k"
		util.CreateDir(videoDirPath)
		util.DoMoveFileToDir(vars.VideoHorizontal2KList, videoDirPath)
	}
	if len(vars.VideoHorizontal3KList) > 0 {
		videoDirPath := rootDir + pathSeparator + uid + "-视频_横屏_3k"
		util.CreateDir(videoDirPath)
		util.DoMoveFileToDir(vars.VideoHorizontal3KList, videoDirPath)
	}
	if len(vars.VideoHorizontal4KList) > 0 {
		videoDirPath := rootDir + pathSeparator + uid + "-视频_横屏_4k"
		util.CreateDir(videoDirPath)
		util.DoMoveFileToDir(vars.VideoHorizontal4KList, videoDirPath)
	}
	if len(vars.VideoHorizontal5KList) > 0 {
		videoDirPath := rootDir + pathSeparator + uid + "-视频_横屏_5k"
		util.CreateDir(videoDirPath)
		util.DoMoveFileToDir(vars.VideoHorizontal5KList, videoDirPath)
	}
	if len(vars.VideoHorizontal6KList) > 0 {
		videoDirPath := rootDir + pathSeparator + uid + "-视频_横屏_6k"
		util.CreateDir(videoDirPath)
		util.DoMoveFileToDir(vars.VideoHorizontal6KList, videoDirPath)
	}
	if len(vars.VideoHorizontal7KList) > 0 {
		videoDirPath := rootDir + pathSeparator + uid + "-视频_横屏_7k"
		util.CreateDir(videoDirPath)
		util.DoMoveFileToDir(vars.VideoHorizontal7KList, videoDirPath)
	}
	if len(vars.VideoHorizontal8KList) > 0 {
		videoDirPath := rootDir + pathSeparator + uid + "-视频_横屏_8k"
		util.CreateDir(videoDirPath)
		util.DoMoveFileToDir(vars.VideoHorizontal8KList, videoDirPath)
	}
	if len(vars.VideoHorizontal9KList) > 0 {
		videoDirPath := rootDir + pathSeparator + uid + "-视频_横屏_9k"
		util.CreateDir(videoDirPath)
		util.DoMoveFileToDir(vars.VideoHorizontal9KList, videoDirPath)
	}
	if len(vars.VideoHorizontalHKList) > 0 {
		videoDirPath := rootDir + pathSeparator + uid + "-视频_横屏_原画质"
		util.CreateDir(videoDirPath)
		util.DoMoveFileToDir(vars.VideoHorizontalHKList, videoDirPath)
	}
	if len(vars.VideoHorizontalStandard720PList) > 0 {
		videoDirPath := rootDir + pathSeparator + uid + "-视频_横屏_720P"
		util.CreateDir(videoDirPath)
		util.DoMoveFileToDir(vars.VideoHorizontalStandard720PList, videoDirPath)
	}
	if len(vars.VideoHorizontalStandard1D5PList) > 0 {
		videoDirPath := rootDir + pathSeparator + uid + "-视频_横屏_1D5KP"
		util.CreateDir(videoDirPath)
		util.DoMoveFileToDir(vars.VideoHorizontalStandard1D5PList, videoDirPath)
	}
	if len(vars.VideoHorizontalStandard1080PList) > 0 {
		videoDirPath := rootDir + pathSeparator + uid + "-视频_横屏_1080P"
		util.CreateDir(videoDirPath)
		util.DoMoveFileToDir(vars.VideoHorizontalStandard1080PList, videoDirPath)
	}
	if len(vars.VideoHorizontalStandard2DKList) > 0 {
		videoDirPath := rootDir + pathSeparator + uid + "-视频_横屏_2DKP"
		util.CreateDir(videoDirPath)
		util.DoMoveFileToDir(vars.VideoHorizontalStandard2DKList, videoDirPath)
	}
	if len(vars.VideoHorizontalStandard2D5KList) > 0 {
		videoDirPath := rootDir + pathSeparator + uid + "-视频_横屏_2D5KP"
		util.CreateDir(videoDirPath)
		util.DoMoveFileToDir(vars.VideoHorizontalStandard2D5KList, videoDirPath)
	}
	if len(vars.VideoHorizontalStandardUltraWide2List) > 0 {
		videoDirPath := rootDir + pathSeparator + uid + "-视频_横屏_UltraWide2KP"
		util.CreateDir(videoDirPath)
		util.DoMoveFileToDir(vars.VideoHorizontalStandardUltraWide2List, videoDirPath)
	}
	if len(vars.VideoHorizontalStandardUltraWide3List) > 0 {
		videoDirPath := rootDir + pathSeparator + uid + "-视频_横屏_UltraWide3KP"
		util.CreateDir(videoDirPath)
		util.DoMoveFileToDir(vars.VideoHorizontalStandardUltraWide3List, videoDirPath)
	}
	if len(vars.VideoHorizontalStandardUltraWide4List) > 0 {
		videoDirPath := rootDir + pathSeparator + uid + "-视频_横屏_UltraWide4KP"
		util.CreateDir(videoDirPath)
		util.DoMoveFileToDir(vars.VideoHorizontalStandardUltraWide4List, videoDirPath)
	}
	if len(vars.VideoHorizontalStandard4KList) > 0 {
		videoDirPath := rootDir + pathSeparator + uid + "-视频_横屏_4KP"
		util.CreateDir(videoDirPath)
		util.DoMoveFileToDir(vars.VideoHorizontalStandard4KList, videoDirPath)
	}
	if len(vars.VideoHorizontalStandard4KHList) > 0 {
		videoDirPath := rootDir + pathSeparator + uid + "-视频_横屏_4KHP"
		util.CreateDir(videoDirPath)
		util.DoMoveFileToDir(vars.VideoHorizontalStandard4KHList, videoDirPath)
	}
	if len(vars.VideoHorizontalStandard8KList) > 0 {
		videoDirPath := rootDir + pathSeparator + uid + "-视频_横屏_8KP"
		util.CreateDir(videoDirPath)
		util.DoMoveFileToDir(vars.VideoHorizontalStandard8KList, videoDirPath)
	}
}

// 移动等比视频
func moveSquareVideo(uid string, rootDir string) {
	pathSeparator := string(os.PathSeparator)
	if len(vars.VideoSquare1KList) > 0 {
		videoDirPath := rootDir + pathSeparator + uid + "-视频_等比_1k"
		util.CreateDir(videoDirPath)
		util.DoMoveFileToDir(vars.VideoSquare1KList, videoDirPath)
	}
	if len(vars.VideoSquare2KList) > 0 {
		videoDirPath := rootDir + pathSeparator + uid + "-视频_等比_2k"
		util.CreateDir(videoDirPath)
		util.DoMoveFileToDir(vars.VideoSquare2KList, videoDirPath)
	}
	if len(vars.VideoSquare3KList) > 0 {
		videoDirPath := rootDir + pathSeparator + uid + "-视频_等比_3k"
		util.CreateDir(videoDirPath)
		util.DoMoveFileToDir(vars.VideoSquare3KList, videoDirPath)
	}
	if len(vars.VideoSquare4KList) > 0 {
		videoDirPath := rootDir + pathSeparator + uid + "-视频_等比_4k"
		util.CreateDir(videoDirPath)
		util.DoMoveFileToDir(vars.VideoSquare4KList, videoDirPath)
	}
	if len(vars.VideoSquare5KList) > 0 {
		videoDirPath := rootDir + pathSeparator + uid + "-视频_等比_5k"
		util.CreateDir(videoDirPath)
		util.DoMoveFileToDir(vars.VideoSquare5KList, videoDirPath)
	}
	if len(vars.VideoSquare6KList) > 0 {
		videoDirPath := rootDir + pathSeparator + uid + "-视频_等比_6k"
		util.CreateDir(videoDirPath)
		util.DoMoveFileToDir(vars.VideoSquare6KList, videoDirPath)
	}
	if len(vars.VideoSquare7KList) > 0 {
		videoDirPath := rootDir + pathSeparator + uid + "-视频_等比_7k"
		util.CreateDir(videoDirPath)
		util.DoMoveFileToDir(vars.VideoSquare7KList, videoDirPath)
	}
	if len(vars.VideoSquare8KList) > 0 {
		videoDirPath := rootDir + pathSeparator + uid + "-视频_等比_8k"
		util.CreateDir(videoDirPath)
		util.DoMoveFileToDir(vars.VideoSquare8KList, videoDirPath)
	}
	if len(vars.VideoSquare9KList) > 0 {
		videoDirPath := rootDir + pathSeparator + uid + "-视频_等比_9k"
		util.CreateDir(videoDirPath)
		util.DoMoveFileToDir(vars.VideoSquare9KList, videoDirPath)
	}
	if len(vars.VideoSquareHKList) > 0 {
		videoDirPath := rootDir + pathSeparator + uid + "-视频_等比_原画质"
		util.CreateDir(videoDirPath)
		util.DoMoveFileToDir(vars.VideoSquareHKList, videoDirPath)
	}
}

// 移动普通视频
func moveNormalVideo(uid string, rootDir string) {
	pathSeparator := string(os.PathSeparator)
	if len(vars.VideoHorizontalNormalList) > 0 {
		videoDirPath := rootDir + pathSeparator + uid + "-视频_横屏_普通"
		util.CreateDir(videoDirPath)
		util.DoMoveFileToDir(vars.VideoHorizontalNormalList, videoDirPath)
	}
	if len(vars.VideoVerticalNormalList) > 0 {
		videoDirPath := rootDir + pathSeparator + uid + "-视频_竖屏_普通"
		util.CreateDir(videoDirPath)
		util.DoMoveFileToDir(vars.VideoVerticalNormalList, videoDirPath)
	}
	if len(vars.VideoSquareNormalList) > 0 {
		videoDirPath := rootDir + pathSeparator + uid + "-视频_等比_普通"
		util.CreateDir(videoDirPath)
		util.DoMoveFileToDir(vars.VideoSquareNormalList, videoDirPath)
	}
}

// 处理垂直视频
func handleVerticalVideo(currentVideoPath string, height int, suffix string) {
	if strings.EqualFold(suffix, ".gif") {
		vars.VideoVerticalGifList = append(vars.VideoVerticalGifList, currentVideoPath)
		return
	}
	if height < 1000 {
		vars.VideoVerticalNormalList = append(vars.VideoVerticalNormalList, currentVideoPath)
	} else if height < 2000 {
		vars.VideoVertical1KList = append(vars.VideoVertical1KList, currentVideoPath)
	} else if height < 3000 {
		vars.VideoVertical2KList = append(vars.VideoVertical2KList, currentVideoPath)
	} else if height < 4000 {
		vars.VideoVertical3KList = append(vars.VideoVertical3KList, currentVideoPath)
	} else if height < 5000 {
		vars.VideoVertical4KList = append(vars.VideoVertical4KList, currentVideoPath)
	} else if height < 6000 {
		vars.VideoVertical5KList = append(vars.VideoVertical5KList, currentVideoPath)
	} else if height < 7000 {
		vars.VideoVertical6KList = append(vars.VideoVertical6KList, currentVideoPath)
	} else if height < 8000 {
		vars.VideoVertical7KList = append(vars.VideoVertical7KList, currentVideoPath)
	} else if height < 9000 {
		vars.VideoVertical8KList = append(vars.VideoVertical8KList, currentVideoPath)
	} else if height < 10000 {
		vars.VideoVertical9KList = append(vars.VideoVertical9KList, currentVideoPath)
	} else if height > 10000 {
		vars.VideoVerticalHKList = append(vars.VideoVerticalHKList, currentVideoPath)
	}
}

// 处理横向视频
func handleHorizontalVideo(currentVideoPath string, width int, height int, suffix string) {
	if strings.EqualFold(suffix, ".gif") {
		vars.VideoHorizontalGifList = append(vars.VideoHorizontalGifList, currentVideoPath)
		return
	}
	// 1280×720（720p）—— 标准高清（HD），常用于早期高清视频或小型屏幕
	if width == 1280 && height == 720 {
		vars.VideoHorizontalStandard720PList = append(vars.VideoHorizontalStandard720PList, currentVideoPath)
		return
	}
	// 1600×900—— 介于HD和FHD之间
	if width == 1600 && height == 900 {
		vars.VideoHorizontalStandard1D5PList = append(vars.VideoHorizontalStandard1D5PList, currentVideoPath)
		return
	}
	// 1920×1080（1080p）—— 全高清，主流显示器、电视和视频的分辨率
	if width == 1920 && height == 1080 {
		vars.VideoHorizontalStandard1080PList = append(vars.VideoHorizontalStandard1080PList, currentVideoPath)
		return
	}
	// 2048×1080（2dk）—— 影院2k
	if width == 2048 && height == 1080 {
		vars.VideoHorizontalStandard2DKList = append(vars.VideoHorizontalStandard2DKList, currentVideoPath)
		return
	}
	// 2560×1440（1440p）—— 俗称“2.5K”，电竞显示器或高端手机屏幕
	if width == 2560 && height == 1440 {
		vars.VideoHorizontalStandard2D5KList = append(vars.VideoHorizontalStandard2D5KList, currentVideoPath)
		return
	}
	// 2560×1080—— 带鱼屏显示器（超宽屏）
	if width == 2560 && height == 1080 {
		vars.VideoHorizontalStandardUltraWide2List = append(vars.VideoHorizontalStandardUltraWide2List, currentVideoPath)
		return
	}
	// 3440×1440—— 带鱼屏显示器（超宽屏）
	if width == 3440 && height == 1440 {
		vars.VideoHorizontalStandardUltraWide3List = append(vars.VideoHorizontalStandardUltraWide3List, currentVideoPath)
		return
	}
	// 5120×2160—— 带鱼屏显示器（超宽屏）
	if width == 5120 && height == 2160 {
		vars.VideoHorizontalStandardUltraWide4List = append(vars.VideoHorizontalStandardUltraWide4List, currentVideoPath)
		return
	}
	// 3840×2160（主流4K）—— 超高清，现代高端显示器、电视和影视制作
	if width == 3840 && height == 2160 {
		vars.VideoHorizontalStandard4KList = append(vars.VideoHorizontalStandard4KList, currentVideoPath)
		return
	}
	// 4096×2160（DCI 4K）—— 电影行业标准
	if width == 4096 && height == 2160 {
		vars.VideoHorizontalStandard4KHList = append(vars.VideoHorizontalStandard4KHList, currentVideoPath)
		return
	}
	// 7680×4320—— 超高清，用于专业影视或高端设备
	if width == 7680 && height == 4320 {
		vars.VideoHorizontalStandard8KList = append(vars.VideoHorizontalStandard8KList, currentVideoPath)
		return
	}
	// 非标规格
	if width < 1000 {
		vars.VideoHorizontalNormalList = append(vars.VideoHorizontalNormalList, currentVideoPath)
	} else if width < 2000 {
		vars.VideoHorizontal1KList = append(vars.VideoHorizontal1KList, currentVideoPath)
	} else if width < 3000 {
		vars.VideoHorizontal2KList = append(vars.VideoHorizontal2KList, currentVideoPath)
	} else if width < 4000 {
		vars.VideoHorizontal3KList = append(vars.VideoHorizontal3KList, currentVideoPath)
	} else if width < 5000 {
		vars.VideoHorizontal4KList = append(vars.VideoHorizontal4KList, currentVideoPath)
	} else if width < 6000 {
		vars.VideoHorizontal5KList = append(vars.VideoHorizontal5KList, currentVideoPath)
	} else if width < 7000 {
		vars.VideoHorizontal6KList = append(vars.VideoHorizontal6KList, currentVideoPath)
	} else if width < 8000 {
		vars.VideoHorizontal7KList = append(vars.VideoHorizontal7KList, currentVideoPath)
	} else if width < 9000 {
		vars.VideoHorizontal8KList = append(vars.VideoHorizontal8KList, currentVideoPath)
	} else if width < 10000 {
		vars.VideoHorizontal9KList = append(vars.VideoHorizontal9KList, currentVideoPath)
	} else if width >= 10000 { // 提示请忽略
		vars.VideoHorizontalHKList = append(vars.VideoHorizontalHKList, currentVideoPath)
	}
}

// 处理等比视频
func handleSquareVideo(currentVideoPath string, width int, height int, suffix string) {
	if strings.EqualFold(suffix, ".gif") {
		vars.VideoSquareGifList = append(vars.VideoSquareGifList, currentVideoPath)
		return
	}
	if width < 1000 {
		vars.VideoSquareNormalList = append(vars.VideoSquareNormalList, currentVideoPath)
	} else if width < 2000 {
		vars.VideoSquare1KList = append(vars.VideoSquare1KList, currentVideoPath)
	} else if width < 3000 {
		vars.VideoSquare2KList = append(vars.VideoSquare2KList, currentVideoPath)
	} else if width < 4000 {
		vars.VideoSquare3KList = append(vars.VideoSquare3KList, currentVideoPath)
	} else if width < 5000 {
		vars.VideoSquare4KList = append(vars.VideoSquare4KList, currentVideoPath)
	} else if width < 6000 {
		vars.VideoSquare5KList = append(vars.VideoSquare5KList, currentVideoPath)
	} else if width < 7000 {
		vars.VideoSquare6KList = append(vars.VideoSquare6KList, currentVideoPath)
	} else if width < 8000 {
		vars.VideoSquare7KList = append(vars.VideoSquare7KList, currentVideoPath)
	} else if width < 9000 {
		vars.VideoSquare8KList = append(vars.VideoSquare8KList, currentVideoPath)
	} else if width < 10000 {
		vars.VideoSquare9KList = append(vars.VideoSquare9KList, currentVideoPath)
	} else if width > 10000 {
		vars.VideoSquareHKList = append(vars.VideoSquareHKList, currentVideoPath)
	}
}
