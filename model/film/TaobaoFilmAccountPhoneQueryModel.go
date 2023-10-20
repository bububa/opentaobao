package film

// TaobaoFilmAccountPhoneQueryModel 结构体
type TaobaoFilmAccountPhoneQueryModel struct {
	// 对外开放ID
	MixUserId string `json:"mix_user_id,omitempty" xml:"mix_user_id,omitempty"`
	// 脱敏昵称
	UserNick string `json:"user_nick,omitempty" xml:"user_nick,omitempty"`
	// 是否新用户
	IsNewUser bool `json:"is_new_user,omitempty" xml:"is_new_user,omitempty"`
}
