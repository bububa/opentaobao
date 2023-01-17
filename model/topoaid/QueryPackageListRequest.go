package topoaid

// QueryPackageListRequest 结构体
type QueryPackageListRequest struct {
	// 用户包裹的身份，淘宝账号/收件人
	QueryRole string `json:"query_role,omitempty" xml:"query_role,omitempty"`
	// 末端类型
	StationType string `json:"station_type,omitempty" xml:"station_type,omitempty"`
	// 用户的唯一标识
	OpenId string `json:"open_id,omitempty" xml:"open_id,omitempty"`
	// 客服场景:CUSTOM，可查询已取件15天内和未取件包裹，收件场景：RECEIVE，仅可查询未取件包裹
	Scene string `json:"scene,omitempty" xml:"scene,omitempty"`
	// 收件人手机号
	Mobile string `json:"mobile,omitempty" xml:"mobile,omitempty"`
	// 1-淘宝端
	OrderShowApp int64 `json:"order_show_app,omitempty" xml:"order_show_app,omitempty"`
	// 页数，从1开始
	PageNo int64 `json:"page_no,omitempty" xml:"page_no,omitempty"`
	// 页大小
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
}
