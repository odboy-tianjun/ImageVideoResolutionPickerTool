package core

import (
	"OdMediaPicker/util"
	"OdMediaPicker/vars"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"image"
	_ "image/gif"  // 导入gif支持
	_ "image/jpeg" // 导入jpeg支持
	_ "image/png"  // 导入png支持
	"os"
	"strings"
)

var ignoreImagePathList []string                       // 忽略的文件路径
var readErrorImagePathList []string                    // 读取信息异常的路径
var imagePath2WidthHeightMap = make(map[string]string) // 图片路径和宽高比
var supportImageTypes = []string{
	".bmp",
	".gif",
	".jpg",
	".jpeg",
	".jpe",
	".png",
	".webp",
	".psb",
}

// gif图片
var gifImageList []string

// 横向图片
var horizontalNormalImageList []string
var horizontal1KImageList []string
var horizontal2KImageList []string
var horizontal3KImageList []string
var horizontal4KImageList []string
var horizontal5KImageList []string
var horizontal6KImageList []string
var horizontal7KImageList []string
var horizontal8KImageList []string
var horizontal9KImageList []string
var horizontalHKImageList []string
var horizontalStandard720PImageList []string
var horizontalStandard1080PImageList []string
var horizontalStandard4KImageList []string
var horizontalStandard8KImageList []string

// 纵向图片
var verticalNormalImageList []string
var vertical1KImageList []string
var vertical2KImageList []string
var vertical3KImageList []string
var vertical4KImageList []string
var vertical5KImageList []string
var vertical6KImageList []string
var vertical7KImageList []string
var vertical8KImageList []string
var vertical9KImageList []string
var verticalHKImageList []string

// 等比图片
var squareNormalImageList []string
var square1KImageList []string
var square2KImageList []string
var square3KImageList []string
var square4KImageList []string
var square5KImageList []string
var square6KImageList []string
var square7KImageList []string
var square8KImageList []string
var square9KImageList []string
var squareHKImageList []string

// psd图片
var psdImageList []string

// 判断是否属于支持的图片文件
func isSupportImage(imageType string) bool {
	for _, supportImageType := range supportImageTypes {
		if strings.EqualFold(supportImageType, imageType) {
			return true
		}
	}
	return false
}

// 读取一般图片文件信息
func readImageInfo(filePath string) (err error, width int, height int) {
	file, err := os.Open(filePath) // 图片文件路径
	if err != nil {
		return err, 0, 0
	}
	defer file.Close()
	img, _, err := image.DecodeConfig(file)
	if err != nil {
		return err, 0, 0
	}
	return nil, img.Width, img.Height
}

func DoHandleImage(rootDir string) {
	total := len(vars.GlobalImagePathList) // 总数
	successCount := 0                      // 成功数
	errorCount := 0                        // 失败数
	ignoreCount := 0                       // 忽略数
	for _, imageFilePath := range vars.GlobalImagePathList {
		suffix := vars.GlobalFilePath2FileExtMap[imageFilePath]
		if isSupportImage(suffix) {
			err, width, height := readImageInfo(imageFilePath)
			if err == nil {
				successCount = successCount + 1
				imagePath2WidthHeightMap[imageFilePath] = fmt.Sprintf("%d-%d", width, height)
				fmt.Printf("=== Image总数: %d, 已读取Info: %d, 成功数: %d, 失败数: %d \n", total, successCount+errorCount+ignoreCount, successCount, errorCount)
			} else {
				errorCount = errorCount + 1
				readErrorImagePathList = append(readErrorImagePathList, imageFilePath)
				fmt.Printf("=== 异常图片: %s \n", imageFilePath)
			}
			continue
		}
		if strings.EqualFold(suffix, ".webp") { // 特殊文件处理, webp为网络常用图片格式
			webpErr, webpWidth, webpHeight := readWebpTypeImage(imageFilePath)
			if webpErr == nil {
				imagePath2WidthHeightMap[imageFilePath] = fmt.Sprintf("%d-%d", webpWidth, webpHeight)
				successCount = successCount + 1
			} else {
				errorCount = errorCount + 1
				fmt.Printf("=== 异常图片: %s \n", imageFilePath)
			}
			continue
		}
		if strings.EqualFold(suffix, ".bmp") { // 特殊文件处理
			bpmErr, bmpWidth, bmpHeight := readBmpInfo(imageFilePath)
			if bpmErr == nil {
				imagePath2WidthHeightMap[imageFilePath] = fmt.Sprintf("%d-%d", bmpWidth, bmpHeight)
				successCount = successCount + 1
			} else {
				errorCount = errorCount + 1
				fmt.Printf("=== 异常图片: %s \n", imageFilePath)
			}
			continue
		}
		if strings.EqualFold(suffix, ".psd") { // 特殊文件处理
			psdImageList = append(psdImageList, imageFilePath)
			successCount = successCount + 1
			continue
		}
		// 其他的直接先忽略吧, 爱改改, 不改拉倒
		ignoreCount = ignoreCount + 1
		ignoreImagePathList = append(ignoreImagePathList, imageFilePath)
	}
	uid := strings.ReplaceAll(uuid.NewV4().String(), "-", "")
	if len(psdImageList) > 0 {
		psdImagePath := rootDir + string(os.PathSeparator) + uid + "-图片_PSD"
		if util.CreateDir(psdImagePath) {
			doMoveFileToDir(psdImageList, psdImagePath)
		}
	}
	if len(readErrorImagePathList) > 0 {
		readInfoErrorPath := rootDir + string(os.PathSeparator) + uid + "-图片_读取异常"
		if util.CreateDir(readInfoErrorPath) {
			doMoveFileToDir(readErrorImagePathList, readInfoErrorPath)
		}
	}
	if len(ignoreImagePathList) > 0 {
		ignorePath := rootDir + string(os.PathSeparator) + uid + "-图片_已忽略"
		if util.CreateDir(ignorePath) {
			doMoveFileToDir(ignoreImagePathList, ignorePath)
		}
	}
	doPickImageFile(uid, rootDir, imagePath2WidthHeightMap)
	fmt.Printf("=== 图片处理完毕(UID): %s \n\n", uid)
}

// 条件图片并分组存放
func doPickImageFile(uid string, rootDir string, imagePath2WidthHeightMap map[string]string) {
	if len(imagePath2WidthHeightMap) == 0 {
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
	gifImageList = append(gifImageList, path)
}

// 处理垂直图片
func handleVerticalImage(currentImagePath string, height int) {
	// 非标规格
	if height < 1000 {
		verticalNormalImageList = append(verticalNormalImageList, currentImagePath)
	} else if height < 2000 {
		vertical1KImageList = append(vertical1KImageList, currentImagePath)
	} else if height < 3000 {
		vertical2KImageList = append(vertical2KImageList, currentImagePath)
	} else if height < 4000 {
		vertical3KImageList = append(vertical3KImageList, currentImagePath)
	} else if height < 5000 {
		vertical4KImageList = append(vertical4KImageList, currentImagePath)
	} else if height < 6000 {
		vertical5KImageList = append(vertical5KImageList, currentImagePath)
	} else if height < 7000 {
		vertical6KImageList = append(vertical6KImageList, currentImagePath)
	} else if height < 8000 {
		vertical7KImageList = append(vertical7KImageList, currentImagePath)
	} else if height < 9000 {
		vertical8KImageList = append(vertical8KImageList, currentImagePath)
	} else if height < 10000 {
		vertical9KImageList = append(vertical9KImageList, currentImagePath)
	} else if height > 10000 {
		verticalHKImageList = append(verticalHKImageList, currentImagePath)
	}
}

// 处理横向图片
func handleHorizontalImage(currentImagePath string, width int, height int) {
	// 1280 x 720 -> 720p
	if width == 1280 && height == 720 {
		horizontalStandard720PImageList = append(horizontalStandard720PImageList, currentImagePath)
		return
	}
	// 1920 x 1080 -> 1080p
	if width == 1920 && height == 1080 {
		horizontalStandard1080PImageList = append(horizontalStandard1080PImageList, currentImagePath)
		return
	}
	// 3840 x 2160 -> 4k
	if width == 3840 && height == 2160 {
		horizontalStandard4KImageList = append(horizontalStandard4KImageList, currentImagePath)
		return
	}
	// 7680 x 4320 -> 8k
	if width == 7680 && height == 4320 {
		horizontalStandard8KImageList = append(horizontalStandard8KImageList, currentImagePath)
		return
	}
	// 非标规格
	if width < 1000 {
		horizontalNormalImageList = append(horizontalNormalImageList, currentImagePath)
	} else if width < 2000 {
		horizontal1KImageList = append(horizontal1KImageList, currentImagePath)
	} else if width < 3000 {
		horizontal2KImageList = append(horizontal2KImageList, currentImagePath)
	} else if width < 4000 {
		horizontal3KImageList = append(horizontal3KImageList, currentImagePath)
	} else if width < 5000 {
		horizontal4KImageList = append(horizontal4KImageList, currentImagePath)
	} else if width < 6000 {
		horizontal5KImageList = append(horizontal5KImageList, currentImagePath)
	} else if width < 7000 {
		horizontal6KImageList = append(horizontal6KImageList, currentImagePath)
	} else if width < 8000 {
		horizontal7KImageList = append(horizontal7KImageList, currentImagePath)
	} else if width < 9000 {
		horizontal8KImageList = append(horizontal8KImageList, currentImagePath)
	} else if width < 10000 {
		horizontal9KImageList = append(horizontal9KImageList, currentImagePath)
	} else if width > 10000 {
		horizontalHKImageList = append(horizontalHKImageList, currentImagePath)
	}
}

// 处理等比图片
func handleSquareImage(currentImagePath string, width int) {
	// 非标规格
	if width < 1000 {
		squareNormalImageList = append(squareNormalImageList, currentImagePath)
	} else if width < 2000 {
		square1KImageList = append(square1KImageList, currentImagePath)
	} else if width < 3000 {
		square2KImageList = append(square2KImageList, currentImagePath)
	} else if width < 4000 {
		square3KImageList = append(square3KImageList, currentImagePath)
	} else if width < 5000 {
		square4KImageList = append(square4KImageList, currentImagePath)
	} else if width < 6000 {
		square5KImageList = append(square5KImageList, currentImagePath)
	} else if width < 7000 {
		square6KImageList = append(square6KImageList, currentImagePath)
	} else if width < 8000 {
		square7KImageList = append(square7KImageList, currentImagePath)
	} else if width < 9000 {
		square8KImageList = append(square8KImageList, currentImagePath)
	} else if width < 10000 {
		square9KImageList = append(square9KImageList, currentImagePath)
	} else if width > 10000 {
		squareHKImageList = append(squareHKImageList, currentImagePath)
	}
}

// 移动水平图片
func moveHorizontalImage(rootDir string, uid string) {
	pathSeparator := string(os.PathSeparator)
	// 标准
	if len(horizontalStandard720PImageList) > 0 {
		horizontalStandard720PImagePath := rootDir + pathSeparator + uid + "-图片_横屏_720P"
		util.CreateDir(horizontalStandard720PImagePath)
		doMoveFileToDir(horizontalStandard720PImageList, horizontalStandard720PImagePath)
	}
	if len(horizontalStandard1080PImageList) > 0 {
		horizontalStandard1080PImagePath := rootDir + pathSeparator + uid + "-图片_横屏_1080P"
		util.CreateDir(horizontalStandard1080PImagePath)
		doMoveFileToDir(horizontalStandard1080PImageList, horizontalStandard1080PImagePath)
	}
	if len(horizontalStandard4KImageList) > 0 {
		horizontalStandard4KImagePath := rootDir + pathSeparator + uid + "-图片_横屏_4KP"
		util.CreateDir(horizontalStandard4KImagePath)
		doMoveFileToDir(horizontalStandard4KImageList, horizontalStandard4KImagePath)
	}
	if len(horizontalStandard8KImageList) > 0 {
		horizontalStandard8KImagePath := rootDir + pathSeparator + uid + "-图片_横屏_8KP"
		util.CreateDir(horizontalStandard8KImagePath)
		doMoveFileToDir(horizontalStandard8KImageList, horizontalStandard8KImagePath)
	}
	// 非标准
	if len(horizontalNormalImageList) > 0 {
		horizontalNormalImagePath := rootDir + pathSeparator + uid + "-图片_横屏_普通"
		util.CreateDir(horizontalNormalImagePath)
		doMoveFileToDir(horizontalNormalImageList, horizontalNormalImagePath)
	}
	if len(horizontal1KImageList) > 0 {
		horizontal1kImagePath := rootDir + pathSeparator + uid + "-图片_横屏_1K"
		util.CreateDir(horizontal1kImagePath)
		doMoveFileToDir(horizontal1KImageList, horizontal1kImagePath)
	}
	if len(horizontal2KImageList) > 0 {
		horizontal2kImagePath := rootDir + pathSeparator + uid + "-图片_横屏_2K"
		util.CreateDir(horizontal2kImagePath)
		doMoveFileToDir(horizontal2KImageList, horizontal2kImagePath)
	}
	if len(horizontal3KImageList) > 0 {
		horizontal3kImagePath := rootDir + pathSeparator + uid + "-图片_横屏_3K"
		util.CreateDir(horizontal3kImagePath)
		doMoveFileToDir(horizontal3KImageList, horizontal3kImagePath)
	}
	if len(horizontal4KImageList) > 0 {
		horizontal4kImagePath := rootDir + pathSeparator + uid + "-图片_横屏_4K"
		util.CreateDir(horizontal4kImagePath)
		doMoveFileToDir(horizontal4KImageList, horizontal4kImagePath)
	}
	if len(horizontal5KImageList) > 0 {
		horizontal5kImagePath := rootDir + pathSeparator + uid + "-图片_横屏_5K"
		util.CreateDir(horizontal5kImagePath)
		doMoveFileToDir(horizontal5KImageList, horizontal5kImagePath)
	}
	if len(horizontal6KImageList) > 0 {
		horizontal6kImagePath := rootDir + pathSeparator + uid + "-图片_横屏_6K"
		util.CreateDir(horizontal6kImagePath)
		doMoveFileToDir(horizontal6KImageList, horizontal6kImagePath)
	}
	if len(horizontal7KImageList) > 0 {
		horizontal7KImagePath := rootDir + pathSeparator + uid + "-图片_横屏_7K"
		util.CreateDir(horizontal7KImagePath)
		doMoveFileToDir(horizontal7KImageList, horizontal7KImagePath)
	}
	if len(horizontal8KImageList) > 0 {
		horizontal8kImagePath := rootDir + pathSeparator + uid + "-图片_横屏_8K"
		util.CreateDir(horizontal8kImagePath)
		doMoveFileToDir(horizontal8KImageList, horizontal8kImagePath)
	}
	if len(horizontal9KImageList) > 0 {
		horizontal9kImagePath := rootDir + pathSeparator + uid + "-图片_横屏_9K"
		util.CreateDir(horizontal9kImagePath)
		doMoveFileToDir(horizontal9KImageList, horizontal9kImagePath)
	}
	if len(horizontalHKImageList) > 0 {
		horizontalHkImagePath := rootDir + pathSeparator + uid + "-图片_横屏_原图"
		util.CreateDir(horizontalHkImagePath)
		doMoveFileToDir(horizontalHKImageList, horizontalHkImagePath)
	}
}

// 移动垂直图片
func moveVerticalImage(rootDir string, uid string) {
	pathSeparator := string(os.PathSeparator)
	// 非标准
	if len(verticalNormalImageList) > 0 {
		verticalNormalImagePath := rootDir + pathSeparator + uid + "-图片_竖屏_普通"
		util.CreateDir(verticalNormalImagePath)
		doMoveFileToDir(verticalNormalImageList, verticalNormalImagePath)
	}
	if len(vertical1KImageList) > 0 {
		vertical1kImagePath := rootDir + pathSeparator + uid + "-图片_竖屏_1K"
		util.CreateDir(vertical1kImagePath)
		doMoveFileToDir(vertical1KImageList, vertical1kImagePath)
	}
	if len(vertical2KImageList) > 0 {
		vertical2kImagePath := rootDir + pathSeparator + uid + "-图片_竖屏_2K"
		util.CreateDir(vertical2kImagePath)
		doMoveFileToDir(vertical2KImageList, vertical2kImagePath)
	}
	if len(vertical3KImageList) > 0 {
		vertical3kImagePath := rootDir + pathSeparator + uid + "-图片_竖屏_3K"
		util.CreateDir(vertical3kImagePath)
		doMoveFileToDir(vertical3KImageList, vertical3kImagePath)
	}
	if len(vertical4KImageList) > 0 {
		vertical4kImagePath := rootDir + pathSeparator + uid + "-图片_竖屏_4K"
		util.CreateDir(vertical4kImagePath)
		doMoveFileToDir(vertical4KImageList, vertical4kImagePath)
	}
	if len(vertical5KImageList) > 0 {
		vertical5kImagePath := rootDir + pathSeparator + uid + "-图片_竖屏_5K"
		util.CreateDir(vertical5kImagePath)
		doMoveFileToDir(vertical5KImageList, vertical5kImagePath)
	}
	if len(vertical6KImageList) > 0 {
		vertical6kImagePath := rootDir + pathSeparator + uid + "-图片_竖屏_6K"
		util.CreateDir(vertical6kImagePath)
		doMoveFileToDir(vertical6KImageList, vertical6kImagePath)
	}
	if len(vertical7KImageList) > 0 {
		vertical7KImagePath := rootDir + pathSeparator + uid + "-图片_竖屏_7K"
		util.CreateDir(vertical7KImagePath)
		doMoveFileToDir(vertical7KImageList, vertical7KImagePath)
	}
	if len(vertical8KImageList) > 0 {
		vertical8kImagePath := rootDir + pathSeparator + uid + "-图片_竖屏_8K"
		util.CreateDir(vertical8kImagePath)
		doMoveFileToDir(vertical8KImageList, vertical8kImagePath)
	}
	if len(vertical9KImageList) > 0 {
		vertical9kImagePath := rootDir + pathSeparator + uid + "-图片_竖屏_9K"
		util.CreateDir(vertical9kImagePath)
		doMoveFileToDir(vertical9KImageList, vertical9kImagePath)
	}
	if len(verticalHKImageList) > 0 {
		verticalHkImagePath := rootDir + pathSeparator + uid + "-图片_竖屏_原图"
		util.CreateDir(verticalHkImagePath)
		doMoveFileToDir(verticalHKImageList, verticalHkImagePath)
	}
}

// 移动等比图片
func moveSquareImage(rootDir string, uid string) {
	pathSeparator := string(os.PathSeparator)
	// 非标准
	if len(squareNormalImageList) > 0 {
		squareNormalImagePath := rootDir + pathSeparator + uid + "-图片_等比_普通"
		util.CreateDir(squareNormalImagePath)
		doMoveFileToDir(squareNormalImageList, squareNormalImagePath)
	}
	if len(square1KImageList) > 0 {
		square1kImagePath := rootDir + pathSeparator + uid + "-图片_等比_1K"
		util.CreateDir(square1kImagePath)
		doMoveFileToDir(square1KImageList, square1kImagePath)
	}
	if len(square2KImageList) > 0 {
		square2kImagePath := rootDir + pathSeparator + uid + "-图片_等比_2K"
		util.CreateDir(square2kImagePath)
		doMoveFileToDir(square2KImageList, square2kImagePath)
	}
	if len(square3KImageList) > 0 {
		square3kImagePath := rootDir + pathSeparator + uid + "-图片_等比_3K"
		util.CreateDir(square3kImagePath)
		doMoveFileToDir(square3KImageList, square3kImagePath)
	}
	if len(square4KImageList) > 0 {
		square4kImagePath := rootDir + pathSeparator + uid + "-图片_等比_4K"
		util.CreateDir(square4kImagePath)
		doMoveFileToDir(square4KImageList, square4kImagePath)
	}
	if len(square5KImageList) > 0 {
		square5kImagePath := rootDir + pathSeparator + uid + "-图片_等比_5K"
		util.CreateDir(square5kImagePath)
		doMoveFileToDir(square5KImageList, square5kImagePath)
	}
	if len(square6KImageList) > 0 {
		square6kImagePath := rootDir + pathSeparator + uid + "-图片_等比_6K"
		util.CreateDir(square6kImagePath)
		doMoveFileToDir(square6KImageList, square6kImagePath)
	}
	if len(square7KImageList) > 0 {
		square7KImagePath := rootDir + pathSeparator + uid + "-图片_等比_7K"
		util.CreateDir(square7KImagePath)
		doMoveFileToDir(square7KImageList, square7KImagePath)
	}
	if len(square8KImageList) > 0 {
		square8kImagePath := rootDir + pathSeparator + uid + "-图片_等比_8K"
		util.CreateDir(square8kImagePath)
		doMoveFileToDir(square8KImageList, square8kImagePath)
	}
	if len(square9KImageList) > 0 {
		square9kImagePath := rootDir + pathSeparator + uid + "-图片_等比_9K"
		util.CreateDir(square9kImagePath)
		doMoveFileToDir(square9KImageList, square9kImagePath)
	}
	if len(squareHKImageList) > 0 {
		squareHkImagePath := rootDir + pathSeparator + uid + "-图片_等比_原图"
		util.CreateDir(squareHkImagePath)
		doMoveFileToDir(squareHKImageList, squareHkImagePath)
	}
}
