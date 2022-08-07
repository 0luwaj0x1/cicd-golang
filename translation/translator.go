package translation


func Translate(word string, language string) string {
	if word != "hello" {
		return ""
	}
	switch language {
	case "english":
		return "hello"
	case "finnish":
		return "hei"
	case "german":
		return "hallo"
	default: 
	return ""			
		
	}
}