package utilities

func FormatText(rawText string, maxLen int) string {
	// Format RLE
	var fmtText string
	for c := 0; c < len(rawText); c += maxLen {
		var cMax = c + maxLen
		if cMax < len(rawText) {
			fmtText += rawText[c:c+70] + "\n"
		} else {
			fmtText += rawText[c:]
		}
	}
	return fmtText
}
