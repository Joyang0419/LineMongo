package lineController

type PushMessagesBody struct {
	UserID   string   `example:"Ubbd98d38fbfc1fe0e7db93a2d8bc9c34"`
	Messages []string `example:"Hello, world"`
}

type ReplyMessagesBody struct {
	ReplyToken string   `example:"bd2b0c9ee9124ef9814913165457afce"`
	Messages   []string `example:"Hello, world"`
}
