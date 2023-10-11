package nrt

// TopAssetDataAuthReqDto 结构体
type TopAssetDataAuthReqDto struct {
	// 员工手机号
	PhoneList []string `json:"phone_list,omitempty" xml:"phone_list>string,omitempty"`
	// 同城站ID
	StoreId int64 `json:"store_id,omitempty" xml:"store_id,omitempty"`
}
