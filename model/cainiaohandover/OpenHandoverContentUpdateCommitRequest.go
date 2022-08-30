package cainiaohandover

// OpenHandoverContentUpdateCommitRequest 结构体
type OpenHandoverContentUpdateCommitRequest struct {
	// 新增和修改大包小包号;新增和修改的时候必填
	ParcelOrderList []HandoverContentUpdateParcelDto `json:"parcel_order_list,omitempty" xml:"parcel_order_list>handover_content_update_parcel_dto,omitempty"`
	// 大包号;LP开头的单号(必填)
	HandoverContentCode string `json:"handover_content_code,omitempty" xml:"handover_content_code,omitempty"`
	// 业务类型(选填)
	BizSource string `json:"biz_source,omitempty" xml:"biz_source,omitempty"`
	// 客户端标示(选填)
	Client string `json:"client,omitempty" xml:"client,omitempty"`
	// 语言(选填)
	Locale string `json:"locale,omitempty" xml:"locale,omitempty"`
	// 更新类型:add、remove、commit;add新增parcel_order_list中小包,remove删除parcel_order_list中小包;commit,忽略parcel_order_list小包,将大包设置为完成组包状态,将大包信息下发给司机上门揽收
	UpdateType string `json:"update_type,omitempty" xml:"update_type,omitempty"`
	// 用户信息
	UserInfo *UserInfoDto `json:"user_info,omitempty" xml:"user_info,omitempty"`
	// 完成更新(必填)例如传入true则代表大包已经完成修改,不能再做修改的操作
	CompleteUpdate bool `json:"complete_update,omitempty" xml:"complete_update,omitempty"`
}
