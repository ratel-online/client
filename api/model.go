package api

type Common struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type LoginResp struct {
	Common
	Data LoginRespData `json:"data"`
}

type LoginRespData struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Token    string `json:"token"`
}
