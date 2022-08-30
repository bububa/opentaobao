package customizemarket

// ProfileQuery 结构体
type ProfileQuery struct {
	// 宠物名称
	PetNick string `json:"pet_nick,omitempty" xml:"pet_nick,omitempty"`
	// 每页数量
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 用户id
	UserId int64 `json:"user_id,omitempty" xml:"user_id,omitempty"`
	// 第几页
	PageNum int64 `json:"page_num,omitempty" xml:"page_num,omitempty"`
}
