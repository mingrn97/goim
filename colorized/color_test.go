package colorized_test

import (
	"testing"

	"itumate.com/im/colorized"
)

func TestColor(t *testing.T) {

	t.Log(colorized.Bold, "This is Bold", colorized.Reset)
	t.Log(colorized.Italic, "This is Italic", colorized.Reset)
	t.Log(colorized.Underline, "This is Underline", colorized.Reset)
	t.Log(colorized.StrikeThrough, "This is StrikeThrough", colorized.Reset)

	t.Log(colorized.BlackDark, "This is BlackDark", colorized.Reset)
	t.Log(colorized.RedDark, "This is RedDark", colorized.Reset)
	t.Log(colorized.GreenDark, "This is GreenDark", colorized.Reset)
	t.Log(colorized.YellowDark, "This is YellowDark", colorized.Reset)
	t.Log(colorized.BlueDark, "This is BlueDark", colorized.Reset)
	t.Log(colorized.PurpleDark, "This is PurpleDark", colorized.Reset)
	t.Log(colorized.CyanDark, "This is CyanDark", colorized.Reset)
	t.Log(colorized.GrayDark, "This is GrayDark", colorized.Reset)

	t.Log(colorized.WhiteDark_, "This is WhiteDark_", colorized.Reset)
	t.Log(colorized.BlackDark_, "This is BlackDark_", colorized.Reset)
	t.Log(colorized.RedDark_, "This is RedDark_", colorized.Reset)
	t.Log(colorized.GreenDark_, "This is GreenDark_", colorized.Reset)
	t.Log(colorized.YellowDark_, "This is YellowDark_", colorized.Reset)
	t.Log(colorized.BlueDark_, "This is BlueDark_", colorized.Reset)
	t.Log(colorized.PurpleDark_, "This is PurpleDark_", colorized.Reset)
	t.Log(colorized.CyanDark_, "This is CyanDark_", colorized.Reset)
	t.Log(colorized.GrayDark_, "This is GrayDark_", colorized.Reset)

	t.Log(colorized.Gray, "This is Gray", colorized.Reset)
	t.Log(colorized.Red, "This is Red", colorized.Reset)
	t.Log(colorized.Green, "This is Green", colorized.Reset)
	t.Log(colorized.Yellow, "This is Yellow", colorized.Reset)
	t.Log(colorized.Blue, "This is Blue", colorized.Reset)
	t.Log(colorized.Purple, "This is Purple", colorized.Reset)
	t.Log(colorized.Cyan, "This is Cyan", colorized.Reset)
	t.Log(colorized.White, "This is White", colorized.Reset)

	t.Log(colorized.Gray_, "This is Gray_", colorized.Reset)
	t.Log(colorized.Red_, "This is Red_", colorized.Reset)
	t.Log(colorized.Green_, "This is Green_", colorized.Reset)
	t.Log(colorized.Yellow_, "This is Yellow_", colorized.Reset)
	t.Log(colorized.Blue_, "This is Blue_", colorized.Reset)
	t.Log(colorized.Purple_, "This is Purple_", colorized.Reset)
	t.Log(colorized.Cyan_, "This is Cyan_", colorized.Reset)
	t.Log(colorized.White_, "This is White_", colorized.Reset)
}

func TestAllColor(t *testing.T) {
	for i := 0; i < 110; i++ {
		t.Logf("\033[%dmThis is Color %d%s\n", i, i, colorized.Reset)
	}
}
