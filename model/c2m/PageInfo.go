package c2m

// PageInfo 结构体
type PageInfo struct {
	// 邀请关系信息
	List []OrganizationInviteInfoVo `json:"list,omitempty" xml:"list>organization_invite_info_vo,omitempty"`
	// 总条数
	Total int64 `json:"total,omitempty" xml:"total,omitempty"`
	// 总页数
	Pages int64 `json:"pages,omitempty" xml:"pages,omitempty"`
	// 页大小
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 第几页
	PageNum int64 `json:"page_num,omitempty" xml:"page_num,omitempty"`
}
