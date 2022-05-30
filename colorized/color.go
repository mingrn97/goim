package colorized

import (
	"runtime"
)

// _    表示背景色
// dark 表示暗色

var (
	Reset = "\033[0m" // clear any style

	Bold          = "\033[1m" // 粗体
	Italic        = "\033[3m" // 斜体
	Underline     = "\033[4m" // 下划
	StrikeThrough = "\033[9m" // 中划线

	BlackDark  = "\033[30m" // 黑色(暗色)
	RedDark    = "\033[31m" // 红色(暗色)
	GreenDark  = "\033[32m" // 绿色(暗色)
	YellowDark = "\033[33m" // 黄色(暗色)
	BlueDark   = "\033[34m" // 蓝色(暗色)
	PurpleDark = "\033[35m" // 紫色(暗色)
	CyanDark   = "\033[36m" // 靛青(暗色)
	GrayDark   = "\033[37m" // 灰色(暗色)

	WhiteDark_  = "\033[7m"  // 白色背景(暗色)
	BlackDark_  = "\033[40m" // 黑色背景(暗色)
	RedDark_    = "\033[41m" // 红色背景(暗色)
	GreenDark_  = "\033[42m" // 绿色背景(暗色)
	YellowDark_ = "\033[43m" // 黄色背景(暗色)
	BlueDark_   = "\033[44m" // 蓝色背景(暗色)
	PurpleDark_ = "\033[45m" // 紫色背景(暗色)
	CyanDark_   = "\033[46m" // 靛青背景(暗色)
	GrayDark_   = "\033[47m" // 灰色背景(暗色)

	Gray   = "\033[90m" // 灰色
	Red    = "\033[91m" // 红色
	Green  = "\033[92m" // 绿色
	Yellow = "\033[93m" // 黄色
	Blue   = "\033[94m" // 蓝色
	Purple = "\033[95m" // 紫色
	Cyan   = "\033[96m" // 靛青
	White  = "\033[97m" // 白色

	Gray_   = "\033[100m" // 灰色背景
	Red_    = "\033[101m" // 红色背景
	Green_  = "\033[102m" // 绿色背景
	Yellow_ = "\033[103m" // 黄色背景
	Blue_   = "\033[104m" // 蓝色背景
	Purple_ = "\033[105m" // 紫色背景
	Cyan_   = "\033[106m" // 靛青背景
	White_  = "\033[107m" // 白色背景
)

func init() {

	// windows 不支持输出颜色值
	if runtime.GOOS == "windows" {
		Bold = ""
		Italic = ""
		Underline = ""

		BlackDark = ""
		RedDark = ""
		GreenDark = ""
		YellowDark = ""
		BlueDark = ""
		PurpleDark = ""
		CyanDark = ""
		GrayDark = ""

		WhiteDark_ = ""
		BlackDark_ = ""
		RedDark_ = ""
		GreenDark_ = ""
		YellowDark_ = ""
		BlueDark_ = ""
		PurpleDark_ = ""
		CyanDark_ = ""
		GrayDark_ = ""

		Gray = ""
		Red = ""
		Green = ""
		Yellow = ""
		Blue = ""
		Purple = ""
		Cyan = ""
		White = ""

		Gray_ = ""
		Red_ = ""
		Green_ = ""
		Yellow_ = ""
		Blue_ = ""
		Purple_ = ""
		Cyan_ = ""
		White_ = ""
	}
}
