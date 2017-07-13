package constants

const Version string = "0.0.1"
const AppName string = "lemonfeeder"
const AppPrefix string = "LF"
const Reset string = "%{F-}%{B-}"

var Icons map[string]string
var BspwmCommand []string

func init() {
	Icons = map[string]string{
		"cpu_chip":        "\uF2DB",
		"gpu_chip":        "\uF11B",
		"cpu_core":        "\uF054",
		"calendar":        "\uF073",
		"check_empty":     "\uF096",
		"check":           "\uF046",
		"unknown":         "\uF059",
		"desktop_generic": "\uF07B",
		"monitor_generic": "\uF108",
		"toggle_off":      "\uF204",
		"toggle_on":       "\uF205",
		"desktop1":        "\uF268",
		"desktop2":        "\uF085",
		"desktop3":        "\uF08E",
		"desktop4":        "\uF075",
		"desktop5":        "\uF0C3",
		"desktop6":        "\uF0C2",
		"arrow_right":     "\uF0DA",
		"temp0":           "\uF2CB",
		"temp1":           "\uF2CA",
		"temp2":           "\uF2C9",
		"temp3":           "\uF2C8",
		"temp4":           "\uF2C7",
	}

	BspwmCommand = []string{"bspc", "query", "-T", "-m"}
}
