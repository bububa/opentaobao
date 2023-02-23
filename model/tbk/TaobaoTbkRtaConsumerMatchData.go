package tbk

// TaobaoTbkRtaConsumerMatchData 结构体
type TaobaoTbkRtaConsumerMatchData struct {
	// 返回结果列表
	ResultList []Resultlist `json:"result_list,omitempty" xml:"result_list>resultlist,omitempty"`
}
