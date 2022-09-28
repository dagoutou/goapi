package message

const (
	LoginMesType    = "LoginMes"
	RegisterMesType = "RegisterMes"
	LoginResMesType = "LoginResMes"
)

type Message struct {
	MessageType string `json:"message_type"`
	MessageData string `json:"message_data"`
}

type Login struct {
	UserId       int    `json:"user_id"`
	UserName     string `json:"user_name"`
	UserPassword string `json:"user_password"`
}

type ResMessage struct {
	ResCode     int    `json:"res_code"`
	MessageData string `json:"message_data"`
}
