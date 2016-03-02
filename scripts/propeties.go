package necrotheif

import(
	"encoding/hex"
	"html/template"
)

// Website properties
var websiteTitle = "NecroTheif"
var author = "Andrew Wetmore"
var backgroundColor = "551A8B"
var foregroundColor = "FFFFFF"
var linkColor = "FFE303"

// Template Properties
var templatesFolder = "templates"
var basicTemplateName = "basic"
var jsFolder = "js"
var jsonFolder = "json"
var errorPage = "error"

// Template structs
type BasicCss struct{
	BackgroundColor			string
	ForegroundColor			string
	LinkColor				string
	SidebarColor			string
	TitleColor				string
	VisitedLinkColor		string
	HoverLinkColor			string
	MenuHoverColor			string
	MessageBackgroundColor		string
	MessageColor			string
}
type BasicHtml struct{
	WebTitle		string
	Author			string
	OtherAuthors		[]string
	StyleSheet		template.CSS
	Javascript		template.JS
	PageTitle		string
	Updated			string
	Message			template.HTML
	Menu			[]BasicItem
	Content			[]BasicContent
}
type BasicContent struct{
	Title	string
	Id	string
	Images	[]BasicImage
	Text	[]template.HTML
}
type BasicItem struct{
	Link	string
	Name	string
	Items	[]BasicItem
}
type BasicImage struct{
	Size	string
	Src	string
	Alt	string
}

// Css properties
var cssBasic *BasicCss
func loadColors(backgroundColor string, foregroundColor string, linkColor string){
	backgroundColor = clampHex(backgroundColor, 48)
	linkColor = clampHex(linkColor, 32)
	cssBasic = &BasicCss{BackgroundColor:backgroundColor, ForegroundColor:foregroundColor, LinkColor:linkColor, SidebarColor:hexMinusInt(backgroundColor, 16), TitleColor:hexMinusInt(backgroundColor, 32), VisitedLinkColor:hexMinusInt(linkColor, 32), HoverLinkColor:hexMinusInt(linkColor, -32), MenuHoverColor:hexMinusInt(backgroundColor, 48), MessageBackgroundColor:hexMinusInt(backgroundColor, -32), MessageColor:hexInvert(hexMinusInt(backgroundColor, -32))}
}

func hexMinusInt(hexVal string, intVal int8) string{
	hexByte, _ := hex.DecodeString(hexVal)
	if(intVal<0){
		for index,_ := range hexByte{
			if(255-hexByte[index]<uint8(-intVal)){
				hexByte[index]=255
			}else{
				hexByte[index]=hexByte[index]+uint8(-intVal)
			}
		}
	} else {
		for index,_ := range hexByte{
			if(hexByte[index]<uint8(intVal)){
				hexByte[index]=0
			}else{
				hexByte[index]=hexByte[index]-uint8(intVal)
			}
		}
	}
	return hex.EncodeToString(hexByte)
}

func hexInvert(hexVal string) string{
	hexByte, _ := hex.DecodeString(hexVal)
	for index,_ := range hexByte{
		hexByte[index]-=128
	}
	return hex.EncodeToString(hexByte)
}

func clampHex(hexVal string, clamp uint8) string{
	hexByte, _ := hex.DecodeString(hexVal)
	for(hexByte[0]<clamp && hexByte[1]<clamp && hexByte[2]<clamp){
		hexByte[0]++
		hexByte[1]++
		hexByte[2]++
	}
	for(hexByte[0]>255-clamp && hexByte[1]>255-clamp && hexByte[2]>255-clamp){
		hexByte[0]--
		hexByte[1]--
		hexByte[2]--
	}
	return hex.EncodeToString(hexByte)
}