package tbk

// TaobaoTbkScVegasSendStatusData 结构体
type TaobaoTbkScVegasSendStatusData struct {
	// 返回结果封装对象
	ResultList []TaobaoTbkScVegasSendStatusMapData `json:"result_list,omitempty" xml:"result_list>taobao_tbk_sc_vegas_send_status_map_data,omitempty"`
}
