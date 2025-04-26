package util

import (
	"image"
	"os"
	"strings"
)

var supportImageTypes = []string{
	".bmp",
	".gif",
	".jpg",
	".jpeg",
	".jpe",
	".png",
	".webp",
	".psb",
	".jfif",
}

// IsSupportImage 判断是否属于支持的图片文件
func IsSupportImage(imageType string) bool {
	for _, supportImageType := range supportImageTypes {
		if strings.EqualFold(supportImageType, strings.ToLower(imageType)) {
			return true
		}
	}
	return false
}

// ReadImageInfo 读取一般图片文件信息
func ReadImageInfo(filePath string) (err error, width int, height int) {
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
