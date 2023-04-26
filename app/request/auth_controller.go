package request

type LoginRequest struct {
	Username string `json:"username" binding:"required,min=3,max=32" required_msg:"用户名不能为空" min_msg:"用户名长度不能小于3" max_msg:"用户名长度不能大于32"`
	Password string `json:"password" binding:"required,min=5,max=32" required_msg:"密码不能为空" min_msg:"密码长度不能小于6" max_msg:"密码长度不能大于32"`
}

type MessageRes struct {
	Content string `json:"content" `
}
