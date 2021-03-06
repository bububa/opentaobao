package westcrm

// DataResult 结构体
type DataResult struct {
	// data
	Datas []MallListVo `json:"datas,omitempty" xml:"datas>mall_list_vo,omitempty"`
	// 商店列表
	ShopList []ShopListVo `json:"shop_list,omitempty" xml:"shop_list>shop_list_vo,omitempty"`
	// 返回信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// data
	Data *MemberAccountIdVo `json:"data,omitempty" xml:"data,omitempty"`
	// 参数code
	Code int64 `json:"code,omitempty" xml:"code,omitempty"`
	// data
	DataList *AlibabaWestcrmActivityInfoGetData `json:"data_list,omitempty" xml:"data_list,omitempty"`
}
