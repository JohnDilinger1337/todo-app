package utility

func DoneMessage(done bool) string {
	message := map[bool]string{
		true:  "This is already done!",
		false: "Not done yet!",
	}
	return message[done]

}
