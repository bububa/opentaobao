package tuanhotel

// StoreDetailVoList 结构体
type StoreDetailVoList struct {
	// 联系电话列表
	AppointPhoneList []AppointPhoneVoList `json:"appoint_phone_list,omitempty" xml:"appoint_phone_list>appoint_phone_vo_list,omitempty"`
	// 房型列表
	RoomTypeList []RoomTypeVoList `json:"room_type_list,omitempty" xml:"room_type_list>room_type_vo_list,omitempty"`
	// 物理酒店ID
	Hid string `json:"hid,omitempty" xml:"hid,omitempty"`
	// 分账描述
	BillDescs string `json:"bill_descs,omitempty" xml:"bill_descs,omitempty"`
	// 地址
	Address string `json:"address,omitempty" xml:"address,omitempty"`
	// 核销账户
	WriteOffAccounts string `json:"write_off_accounts,omitempty" xml:"write_off_accounts,omitempty"`
	// 名字
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 分账信息
	BillInfos string `json:"bill_infos,omitempty" xml:"bill_infos,omitempty"`
	// 标准酒店ID
	Shid int64 `json:"shid,omitempty" xml:"shid,omitempty"`
	// 门店id
	StoreId int64 `json:"store_id,omitempty" xml:"store_id,omitempty"`
}
