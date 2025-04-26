package vars

import (
	_ "embed" // 内嵌文件支持
)

//go:embed ffprobe.exe
var VideoDecodeToolWin64 []byte
var VideoIgnorePathList []string                          // 忽略的文件路径
var VideoReadErrorPathList []string                       // 读取信息异常的路径
var VideoPath2WidthHeightMap = make(map[string]string)    // 视频路径和宽高比
var VideoPath2WidthHeightTagMap = make(map[string]string) // 视频路径和宽高比[640x480]
var VideoPath2DurationMap = make(map[string]string)       // 视频路径和时长
var VideoHorizontalNormalList []string
var VideoHorizontalGifList []string
var VideoHorizontal1KList []string
var VideoHorizontal2KList []string
var VideoHorizontal3KList []string
var VideoHorizontal4KList []string
var VideoHorizontal5KList []string
var VideoHorizontal6KList []string
var VideoHorizontal7KList []string
var VideoHorizontal8KList []string
var VideoHorizontal9KList []string
var VideoHorizontalHKList []string
var VideoHorizontalStandard720PList []string
var VideoHorizontalStandard1D5PList []string
var VideoHorizontalStandard1080PList []string
var VideoHorizontalStandard2DKList []string
var VideoHorizontalStandard2D5KList []string
var VideoHorizontalStandardUltraWide2List []string
var VideoHorizontalStandardUltraWide3List []string
var VideoHorizontalStandardUltraWide4List []string
var VideoHorizontalStandard4KList []string
var VideoHorizontalStandard4KHList []string
var VideoHorizontalStandard8KList []string
var VideoVerticalNormalList []string
var VideoVerticalGifList []string
var VideoVertical1KList []string
var VideoVertical2KList []string
var VideoVertical3KList []string
var VideoVertical4KList []string
var VideoVertical5KList []string
var VideoVertical6KList []string
var VideoVertical7KList []string
var VideoVertical8KList []string
var VideoVertical9KList []string
var VideoVerticalHKList []string
var VideoSquareNormalList []string
var VideoSquareGifList []string
var VideoSquare1KList []string
var VideoSquare2KList []string
var VideoSquare3KList []string
var VideoSquare4KList []string
var VideoSquare5KList []string
var VideoSquare6KList []string
var VideoSquare7KList []string
var VideoSquare8KList []string
var VideoSquare9KList []string
var VideoSquareHKList []string
