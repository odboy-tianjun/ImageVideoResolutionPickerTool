package core

import (
	"ImageVideoResolutionPickerTool/util"
	"ImageVideoResolutionPickerTool/vars"
	"fmt"
	uuid "github.com/satori/go.uuid"
	_ "image/gif"  // 导入gif支持
	_ "image/jpeg" // 导入jpeg支持
	_ "image/png"  // 导入png支持
	"os"
	"strings"
)

func DoHandleImage(rootDir string) {
	total := len(vars.GlobalImagePathList) // 总数
	successCount := 0                      // 成功数
	errorCount := 0                        // 失败数
	ignoreCount := 0                       // 忽略数
	for _, imageFilePath := range vars.GlobalImagePathList {
		suffix := vars.GlobalFilePath2FileExtMap[imageFilePath]
		if util.IsSupportImage(suffix) {
			err, width, height := util.ReadImageInfo(imageFilePath)
			if err == nil {
				successCount = successCount + 1
				vars.ImagePath2WidthHeightMap[imageFilePath] = fmt.Sprintf("%d-%d", width, height)
				fmt.Printf("=== Image总数: %d, 已读取Info: %d, 成功数: %d, 失败数: %d \n", total, successCount+errorCount+ignoreCount, successCount, errorCount)
			} else {
				errorCount = errorCount + 1
				vars.ImageReadErrorPathList = append(vars.ImageReadErrorPathList, imageFilePath)
				fmt.Printf("=== 异常图片: %s \n", imageFilePath)
			}
			continue
		}
		if strings.EqualFold(suffix, ".webp") { // 特殊文件处理, webp为网络常用图片格式
			webpErr, webpWidth, webpHeight := ReadWebpImage(imageFilePath)
			if webpErr == nil {
				vars.ImagePath2WidthHeightMap[imageFilePath] = fmt.Sprintf("%d-%d", webpWidth, webpHeight)
				successCount = successCount + 1
			} else {
				errorCount = errorCount + 1
				fmt.Printf("=== 异常图片: %s \n", imageFilePath)
			}
			continue
		}
		if strings.EqualFold(suffix, ".bmp") { // 特殊文件处理
			bpmErr, bmpWidth, bmpHeight := ReadBmpImage(imageFilePath)
			if bpmErr == nil {
				vars.ImagePath2WidthHeightMap[imageFilePath] = fmt.Sprintf("%d-%d", bmpWidth, bmpHeight)
				successCount = successCount + 1
			} else {
				errorCount = errorCount + 1
				fmt.Printf("=== 异常图片: %s \n", imageFilePath)
			}
			continue
		}
		if strings.EqualFold(suffix, ".psd") { // 特殊文件处理
			vars.ImagePsdList = append(vars.ImagePsdList, imageFilePath)
			successCount = successCount + 1
			continue
		}
		// 其他的直接先忽略吧, 爱改改, 不改拉倒
		ignoreCount = ignoreCount + 1
		vars.ImageIgnorePathList = append(vars.ImageIgnorePathList, imageFilePath)
	}
	uid := strings.ReplaceAll(uuid.NewV4().String(), "-", "")
	if len(vars.ImagePsdList) > 0 {
		psdImagePath := rootDir + string(os.PathSeparator) + uid + "-图片_PSD"
		if util.CreateDir(psdImagePath) {
			util.DoMoveFileToDir(vars.ImagePsdList, psdImagePath)
		}
	}
	if len(vars.ImageReadErrorPathList) > 0 {
		readInfoErrorPath := rootDir + string(os.PathSeparator) + uid + "-图片_读取异常"
		if util.CreateDir(readInfoErrorPath) {
			util.DoMoveFileToDir(vars.ImageReadErrorPathList, readInfoErrorPath)
		}
	}
	if len(vars.ImageIgnorePathList) > 0 {
		ignorePath := rootDir + string(os.PathSeparator) + uid + "-图片_已忽略"
		if util.CreateDir(ignorePath) {
			util.DoMoveFileToDir(vars.ImageIgnorePathList, ignorePath)
		}
	}
	doPickImageFile(uid, rootDir, vars.ImagePath2WidthHeightMap)
	fmt.Printf("=== 图片处理完毕(UID): %s \n\n", uid)
}

// 条件图片并分组存放
func doPickImageFile(uid string, rootDir string, imagePath2WidthHeightMap map[string]string) {
	if len(vars.ImagePath2WidthHeightMap) == 0 {
		fmt.Printf("=== 当前目录下没有扫描到图片文件, %s \n", rootDir)
		return
	}
	for currentImagePath, infoStr := range imagePath2WidthHeightMap {
		width2Height := strings.Split(infoStr, "-")
		width := util.String2int(width2Height[0])
		height := util.String2int(width2Height[1])
		suffix := vars.GlobalFilePath2FileExtMap[currentImagePath]
		if strings.EqualFold(suffix, ".gif") {
			handleGifImage(currentImagePath)
			continue
		}
		if width > height {
			handleHorizontalImage(currentImagePath, width, height)
			continue
		}
		if width < height {
			handleVerticalImage(currentImagePath, height)
			continue
		}
		handleSquareImage(currentImagePath, width)
	}
	moveHorizontalImage(rootDir, uid)
	moveVerticalImage(rootDir, uid)
	moveSquareImage(rootDir, uid)
}

// 统一处理gif
func handleGifImage(path string) {
	vars.ImageGifList = append(vars.ImageGifList, path)
}

// 处理垂直图片
func handleVerticalImage(currentImagePath string, height int) {
	// 非标规格
	if height < 1000 {
		vars.ImageVerticalNormalList = append(vars.ImageVerticalNormalList, currentImagePath)
	} else if height < 2000 {
		vars.ImageVertical1KList = append(vars.ImageVertical1KList, currentImagePath)
	} else if height < 3000 {
		vars.ImageVertical2KList = append(vars.ImageVertical2KList, currentImagePath)
	} else if height < 4000 {
		vars.ImageVertical3KList = append(vars.ImageVertical3KList, currentImagePath)
	} else if height < 5000 {
		vars.ImageVertical4KList = append(vars.ImageVertical4KList, currentImagePath)
	} else if height < 6000 {
		vars.ImageVertical5KList = append(vars.ImageVertical5KList, currentImagePath)
	} else if height < 7000 {
		vars.ImageVertical6KList = append(vars.ImageVertical6KList, currentImagePath)
	} else if height < 8000 {
		vars.ImageVertical7KList = append(vars.ImageVertical7KList, currentImagePath)
	} else if height < 9000 {
		vars.ImageVertical8KList = append(vars.ImageVertical8KList, currentImagePath)
	} else if height < 10000 {
		vars.ImageVertical9KList = append(vars.ImageVertical9KList, currentImagePath)
	} else if height >= 10000 { // 提示请忽略
		vars.ImageVerticalHKList = append(vars.ImageVerticalHKList, currentImagePath)
	}
}

// 处理横向图片
func handleHorizontalImage(currentImagePath string, width int, height int) {
	// 1280×720（720p）—— 标准高清（HD），常用于早期高清视频或小型屏幕
	if width == 1280 && height == 720 {
		vars.ImageHorizontalStandard720PList = append(vars.ImageHorizontalStandard720PList, currentImagePath)
		return
	}
	// 1600×900—— 介于HD和FHD之间
	if width == 1600 && height == 900 {
		vars.ImageHorizontalStandard1D5KList = append(vars.ImageHorizontalStandard1D5KList, currentImagePath)
		return
	}
	// 1920×1080（1080p）—— 全高清，主流显示器、电视和视频的分辨率
	if width == 1920 && height == 1080 {
		vars.ImageHorizontalStandard1080PList = append(vars.ImageHorizontalStandard1080PList, currentImagePath)
		return
	}
	// 2048×1080（2dk）—— 影院2k
	if width == 2048 && height == 1080 {
		vars.ImageHorizontalStandard2DKList = append(vars.ImageHorizontalStandard2DKList, currentImagePath)
		return
	}
	// 2560×1440（1440p）—— 俗称“2.5K”，电竞显示器或高端手机屏幕
	if width == 2560 && height == 1440 {
		vars.ImageHorizontalStandard2D5KList = append(vars.ImageHorizontalStandard2D5KList, currentImagePath)
		return
	}
	// 2560×1080—— 带鱼屏显示器（超宽屏）
	if width == 2560 && height == 1080 {
		vars.ImageHorizontalStandardUltraWide2List = append(vars.ImageHorizontalStandardUltraWide2List, currentImagePath)
		return
	}
	// 3440×1440—— 带鱼屏显示器（超宽屏）
	if width == 3440 && height == 1440 {
		vars.ImageHorizontalStandardUltraWide3List = append(vars.ImageHorizontalStandardUltraWide3List, currentImagePath)
		return
	}
	// 5120×2160—— 带鱼屏显示器（超宽屏）
	if width == 5120 && height == 2160 {
		vars.ImageHorizontalStandardUltraWide4List = append(vars.ImageHorizontalStandardUltraWide4List, currentImagePath)
		return
	}
	// 3840×2160（主流4K）—— 超高清，现代高端显示器、电视和影视制作
	if width == 3840 && height == 2160 {
		vars.ImageHorizontalStandard4KList = append(vars.ImageHorizontalStandard4KList, currentImagePath)
		return
	}
	// 4096×2160（DCI 4K）—— 电影行业标准
	if width == 4096 && height == 2160 {
		vars.ImageHorizontalStandard4KHList = append(vars.ImageHorizontalStandard4KHList, currentImagePath)
		return
	}
	// 7680×4320—— 超高清，用于专业影视或高端设备
	if width == 7680 && height == 4320 {
		vars.ImageHorizontalStandard8KList = append(vars.ImageHorizontalStandard8KList, currentImagePath)
		return
	}
	// 非标规格
	if width < 1000 {
		vars.ImageHorizontalNormalList = append(vars.ImageHorizontalNormalList, currentImagePath)
	} else if width < 2000 {
		vars.ImageHorizontal1KList = append(vars.ImageHorizontal1KList, currentImagePath)
	} else if width < 3000 {
		vars.ImageHorizontal2KList = append(vars.ImageHorizontal2KList, currentImagePath)
	} else if width < 4000 {
		vars.ImageHorizontal3KList = append(vars.ImageHorizontal3KList, currentImagePath)
	} else if width < 5000 {
		vars.ImageHorizontal4KList = append(vars.ImageHorizontal4KList, currentImagePath)
	} else if width < 6000 {
		vars.ImageHorizontal5KList = append(vars.ImageHorizontal5KList, currentImagePath)
	} else if width < 7000 {
		vars.ImageHorizontal6KList = append(vars.ImageHorizontal6KList, currentImagePath)
	} else if width < 8000 {
		vars.ImageHorizontal7KList = append(vars.ImageHorizontal7KList, currentImagePath)
	} else if width < 9000 {
		vars.ImageHorizontal8KList = append(vars.ImageHorizontal8KList, currentImagePath)
	} else if width < 10000 {
		vars.ImageHorizontal9KList = append(vars.ImageHorizontal9KList, currentImagePath)
	} else if width >= 10000 { // 提示请忽略
		vars.ImageHorizontalHKList = append(vars.ImageHorizontalHKList, currentImagePath)
	}
}

// 处理等比图片
func handleSquareImage(currentImagePath string, width int) {
	// 非标规格
	if width < 1000 {
		vars.ImageSquareNormalList = append(vars.ImageSquareNormalList, currentImagePath)
	} else if width < 2000 {
		vars.ImageSquare1KList = append(vars.ImageSquare1KList, currentImagePath)
	} else if width < 3000 {
		vars.ImageSquare2KList = append(vars.ImageSquare2KList, currentImagePath)
	} else if width < 4000 {
		vars.ImageSquare3KList = append(vars.ImageSquare3KList, currentImagePath)
	} else if width < 5000 {
		vars.ImageSquare4KList = append(vars.ImageSquare4KList, currentImagePath)
	} else if width < 6000 {
		vars.ImageSquare5KList = append(vars.ImageSquare5KList, currentImagePath)
	} else if width < 7000 {
		vars.ImageSquare6KList = append(vars.ImageSquare6KList, currentImagePath)
	} else if width < 8000 {
		vars.ImageSquare7KList = append(vars.ImageSquare7KList, currentImagePath)
	} else if width < 9000 {
		vars.ImageSquare8KList = append(vars.ImageSquare8KList, currentImagePath)
	} else if width < 10000 {
		vars.ImageSquare9KList = append(vars.ImageSquare9KList, currentImagePath)
	} else if width > 10000 {
		vars.ImageSquareHKList = append(vars.ImageSquareHKList, currentImagePath)
	}
}

// 移动水平图片
func moveHorizontalImage(rootDir string, uid string) {
	pathSeparator := string(os.PathSeparator)
	// 标准
	if len(vars.ImageHorizontalStandard720PList) > 0 {
		imageDirPath := rootDir + pathSeparator + uid + "-图片_横屏_720P"
		util.CreateDir(imageDirPath)
		util.DoMoveFileToDir(vars.ImageHorizontalStandard720PList, imageDirPath)
	}
	if len(vars.ImageHorizontalStandard1D5KList) > 0 {
		imageDirPath := rootDir + pathSeparator + uid + "-图片_横屏_1D5KP"
		util.CreateDir(imageDirPath)
		util.DoMoveFileToDir(vars.ImageHorizontalStandard1D5KList, imageDirPath)
	}
	if len(vars.ImageHorizontalStandard1080PList) > 0 {
		imageDirPath := rootDir + pathSeparator + uid + "-图片_横屏_1080P"
		util.CreateDir(imageDirPath)
		util.DoMoveFileToDir(vars.ImageHorizontalStandard1080PList, imageDirPath)
	}
	if len(vars.ImageHorizontalStandard2DKList) > 0 {
		imageDirPath := rootDir + pathSeparator + uid + "-图片_横屏_2DKP"
		util.CreateDir(imageDirPath)
		util.DoMoveFileToDir(vars.ImageHorizontalStandard2DKList, imageDirPath)
	}
	if len(vars.ImageHorizontalStandard2D5KList) > 0 {
		imageDirPath := rootDir + pathSeparator + uid + "-图片_横屏_2D5KP"
		util.CreateDir(imageDirPath)
		util.DoMoveFileToDir(vars.ImageHorizontalStandard2D5KList, imageDirPath)
	}
	if len(vars.ImageHorizontalStandardUltraWide2List) > 0 {
		imageDirPath := rootDir + pathSeparator + uid + "-图片_横屏_ultraWide2KP"
		util.CreateDir(imageDirPath)
		util.DoMoveFileToDir(vars.ImageHorizontalStandardUltraWide2List, imageDirPath)
	}
	if len(vars.ImageHorizontalStandardUltraWide3List) > 0 {
		imageDirPath := rootDir + pathSeparator + uid + "-图片_横屏_ultraWide3KP"
		util.CreateDir(imageDirPath)
		util.DoMoveFileToDir(vars.ImageHorizontalStandardUltraWide3List, imageDirPath)
	}
	if len(vars.ImageHorizontalStandardUltraWide4List) > 0 {
		imageDirPath := rootDir + pathSeparator + uid + "-图片_横屏_ultraWide4KP"
		util.CreateDir(imageDirPath)
		util.DoMoveFileToDir(vars.ImageHorizontalStandardUltraWide4List, imageDirPath)
	}
	if len(vars.ImageHorizontalStandard4KList) > 0 {
		imageDirPath := rootDir + pathSeparator + uid + "-图片_横屏_4KP"
		util.CreateDir(imageDirPath)
		util.DoMoveFileToDir(vars.ImageHorizontalStandard4KList, imageDirPath)
	}
	if len(vars.ImageHorizontalStandard4KHList) > 0 {
		imageDirPath := rootDir + pathSeparator + uid + "-图片_横屏_4KHP"
		util.CreateDir(imageDirPath)
		util.DoMoveFileToDir(vars.ImageHorizontalStandard4KHList, imageDirPath)
	}
	if len(vars.ImageHorizontalStandard8KList) > 0 {
		imageDirPath := rootDir + pathSeparator + uid + "-图片_横屏_8KP"
		util.CreateDir(imageDirPath)
		util.DoMoveFileToDir(vars.ImageHorizontalStandard8KList, imageDirPath)
	}
	// 非标准
	if len(vars.ImageHorizontalNormalList) > 0 {
		imageDirPath := rootDir + pathSeparator + uid + "-图片_横屏_普通"
		util.CreateDir(imageDirPath)
		util.DoMoveFileToDir(vars.ImageHorizontalNormalList, imageDirPath)
	}
	if len(vars.ImageHorizontal1KList) > 0 {
		imageDirPath := rootDir + pathSeparator + uid + "-图片_横屏_1K"
		util.CreateDir(imageDirPath)
		util.DoMoveFileToDir(vars.ImageHorizontal1KList, imageDirPath)
	}
	if len(vars.ImageHorizontal2KList) > 0 {
		imageDirPath := rootDir + pathSeparator + uid + "-图片_横屏_2K"
		util.CreateDir(imageDirPath)
		util.DoMoveFileToDir(vars.ImageHorizontal2KList, imageDirPath)
	}
	if len(vars.ImageHorizontal3KList) > 0 {
		imageDirPath := rootDir + pathSeparator + uid + "-图片_横屏_3K"
		util.CreateDir(imageDirPath)
		util.DoMoveFileToDir(vars.ImageHorizontal3KList, imageDirPath)
	}
	if len(vars.ImageHorizontal4KList) > 0 {
		imageDirPath := rootDir + pathSeparator + uid + "-图片_横屏_4K"
		util.CreateDir(imageDirPath)
		util.DoMoveFileToDir(vars.ImageHorizontal4KList, imageDirPath)
	}
	if len(vars.ImageHorizontal5KList) > 0 {
		imageDirPath := rootDir + pathSeparator + uid + "-图片_横屏_5K"
		util.CreateDir(imageDirPath)
		util.DoMoveFileToDir(vars.ImageHorizontal5KList, imageDirPath)
	}
	if len(vars.ImageHorizontal6KList) > 0 {
		imageDirPath := rootDir + pathSeparator + uid + "-图片_横屏_6K"
		util.CreateDir(imageDirPath)
		util.DoMoveFileToDir(vars.ImageHorizontal6KList, imageDirPath)
	}
	if len(vars.ImageHorizontal7KList) > 0 {
		imageDirPath := rootDir + pathSeparator + uid + "-图片_横屏_7K"
		util.CreateDir(imageDirPath)
		util.DoMoveFileToDir(vars.ImageHorizontal7KList, imageDirPath)
	}
	if len(vars.ImageHorizontal8KList) > 0 {
		imageDirPath := rootDir + pathSeparator + uid + "-图片_横屏_8K"
		util.CreateDir(imageDirPath)
		util.DoMoveFileToDir(vars.ImageHorizontal8KList, imageDirPath)
	}
	if len(vars.ImageHorizontal9KList) > 0 {
		imageDirPath := rootDir + pathSeparator + uid + "-图片_横屏_9K"
		util.CreateDir(imageDirPath)
		util.DoMoveFileToDir(vars.ImageHorizontal9KList, imageDirPath)
	}
	if len(vars.ImageHorizontalHKList) > 0 {
		imageDirPath := rootDir + pathSeparator + uid + "-图片_横屏_原图"
		util.CreateDir(imageDirPath)
		util.DoMoveFileToDir(vars.ImageHorizontalHKList, imageDirPath)
	}
}

// 移动垂直图片
func moveVerticalImage(rootDir string, uid string) {
	pathSeparator := string(os.PathSeparator)
	// 非标准
	if len(vars.ImageVerticalNormalList) > 0 {
		imageDirPath := rootDir + pathSeparator + uid + "-图片_竖屏_普通"
		util.CreateDir(imageDirPath)
		util.DoMoveFileToDir(vars.ImageVerticalNormalList, imageDirPath)
	}
	if len(vars.ImageVertical1KList) > 0 {
		imageDirPath := rootDir + pathSeparator + uid + "-图片_竖屏_1K"
		util.CreateDir(imageDirPath)
		util.DoMoveFileToDir(vars.ImageVertical1KList, imageDirPath)
	}
	if len(vars.ImageVertical2KList) > 0 {
		imageDirPath := rootDir + pathSeparator + uid + "-图片_竖屏_2K"
		util.CreateDir(imageDirPath)
		util.DoMoveFileToDir(vars.ImageVertical2KList, imageDirPath)
	}
	if len(vars.ImageVertical3KList) > 0 {
		imageDirPath := rootDir + pathSeparator + uid + "-图片_竖屏_3K"
		util.CreateDir(imageDirPath)
		util.DoMoveFileToDir(vars.ImageVertical3KList, imageDirPath)
	}
	if len(vars.ImageVertical4KList) > 0 {
		imageDirPath := rootDir + pathSeparator + uid + "-图片_竖屏_4K"
		util.CreateDir(imageDirPath)
		util.DoMoveFileToDir(vars.ImageVertical4KList, imageDirPath)
	}
	if len(vars.ImageVertical5KList) > 0 {
		imageDirPath := rootDir + pathSeparator + uid + "-图片_竖屏_5K"
		util.CreateDir(imageDirPath)
		util.DoMoveFileToDir(vars.ImageVertical5KList, imageDirPath)
	}
	if len(vars.ImageVertical6KList) > 0 {
		imageDirPath := rootDir + pathSeparator + uid + "-图片_竖屏_6K"
		util.CreateDir(imageDirPath)
		util.DoMoveFileToDir(vars.ImageVertical6KList, imageDirPath)
	}
	if len(vars.ImageVertical7KList) > 0 {
		imageDirPath := rootDir + pathSeparator + uid + "-图片_竖屏_7K"
		util.CreateDir(imageDirPath)
		util.DoMoveFileToDir(vars.ImageVertical7KList, imageDirPath)
	}
	if len(vars.ImageVertical8KList) > 0 {
		imageDirPath := rootDir + pathSeparator + uid + "-图片_竖屏_8K"
		util.CreateDir(imageDirPath)
		util.DoMoveFileToDir(vars.ImageVertical8KList, imageDirPath)
	}
	if len(vars.ImageVertical9KList) > 0 {
		imageDirPath := rootDir + pathSeparator + uid + "-图片_竖屏_9K"
		util.CreateDir(imageDirPath)
		util.DoMoveFileToDir(vars.ImageVertical9KList, imageDirPath)
	}
	if len(vars.ImageVerticalHKList) > 0 {
		imageDirPath := rootDir + pathSeparator + uid + "-图片_竖屏_原图"
		util.CreateDir(imageDirPath)
		util.DoMoveFileToDir(vars.ImageVerticalHKList, imageDirPath)
	}
}

// 移动等比图片
func moveSquareImage(rootDir string, uid string) {
	pathSeparator := string(os.PathSeparator)
	// 非标准
	if len(vars.ImageSquareNormalList) > 0 {
		imageDirPath := rootDir + pathSeparator + uid + "-图片_等比_普通"
		util.CreateDir(imageDirPath)
		util.DoMoveFileToDir(vars.ImageSquareNormalList, imageDirPath)
	}
	if len(vars.ImageSquare1KList) > 0 {
		imageDirPath := rootDir + pathSeparator + uid + "-图片_等比_1K"
		util.CreateDir(imageDirPath)
		util.DoMoveFileToDir(vars.ImageSquare1KList, imageDirPath)
	}
	if len(vars.ImageSquare2KList) > 0 {
		imageDirPath := rootDir + pathSeparator + uid + "-图片_等比_2K"
		util.CreateDir(imageDirPath)
		util.DoMoveFileToDir(vars.ImageSquare2KList, imageDirPath)
	}
	if len(vars.ImageSquare3KList) > 0 {
		imageDirPath := rootDir + pathSeparator + uid + "-图片_等比_3K"
		util.CreateDir(imageDirPath)
		util.DoMoveFileToDir(vars.ImageSquare3KList, imageDirPath)
	}
	if len(vars.ImageSquare4KList) > 0 {
		imageDirPath := rootDir + pathSeparator + uid + "-图片_等比_4K"
		util.CreateDir(imageDirPath)
		util.DoMoveFileToDir(vars.ImageSquare4KList, imageDirPath)
	}
	if len(vars.ImageSquare5KList) > 0 {
		imageDirPath := rootDir + pathSeparator + uid + "-图片_等比_5K"
		util.CreateDir(imageDirPath)
		util.DoMoveFileToDir(vars.ImageSquare5KList, imageDirPath)
	}
	if len(vars.ImageSquare6KList) > 0 {
		imageDirPath := rootDir + pathSeparator + uid + "-图片_等比_6K"
		util.CreateDir(imageDirPath)
		util.DoMoveFileToDir(vars.ImageSquare6KList, imageDirPath)
	}
	if len(vars.ImageSquare7KList) > 0 {
		imageDirPath := rootDir + pathSeparator + uid + "-图片_等比_7K"
		util.CreateDir(imageDirPath)
		util.DoMoveFileToDir(vars.ImageSquare7KList, imageDirPath)
	}
	if len(vars.ImageSquare8KList) > 0 {
		imageDirPath := rootDir + pathSeparator + uid + "-图片_等比_8K"
		util.CreateDir(imageDirPath)
		util.DoMoveFileToDir(vars.ImageSquare8KList, imageDirPath)
	}
	if len(vars.ImageSquare9KList) > 0 {
		imageDirPath := rootDir + pathSeparator + uid + "-图片_等比_9K"
		util.CreateDir(imageDirPath)
		util.DoMoveFileToDir(vars.ImageSquare9KList, imageDirPath)
	}
	if len(vars.ImageSquareHKList) > 0 {
		imageDirPath := rootDir + pathSeparator + uid + "-图片_等比_原图"
		util.CreateDir(imageDirPath)
		util.DoMoveFileToDir(vars.ImageSquareHKList, imageDirPath)
	}
}
