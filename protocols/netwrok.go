package protocols

// 授权协议第一版
type AuthMethod struct {
	Key     []byte
	Usrname []byte
}

// 客户端发向服务端格式
type SendMsg struct {
	Content []byte `json:"content"`
	Action  int    `json:"action"`
	MessBox string `json:"messbox"`
}

// 客户端发向服务端格式(加强版)
type SendMsgPlus struct {
	Content []byte     `json:"content"`
	Action  int        `json:"action"`
	MessBox string     `json:"messbox"`
	Auth    AuthMethod `json:"auth"`
}
type Response struct {
	StatusCode int
	Content    []byte
	Footer     string
}

// 通用指令协议
type CommonCommand struct {
	Header   string
	Cmd      map[string]string
	Actionid int
	//version 2.1
	Auth AuthMethod
}
