package tbk

import (
	"sync"
)

// TaobaoTbkScVegasSendStatusData 结构体
type TaobaoTbkScVegasSendStatusData struct {
	// 返回结果封装对象
	ResultList []TaobaoTbkScVegasSendStatusMapData `json:"result_list,omitempty" xml:"result_list>taobao_tbk_sc_vegas_send_status_map_data,omitempty"`
}

var poolTaobaoTbkScVegasSendStatusData = sync.Pool{
	New: func() any {
		return new(TaobaoTbkScVegasSendStatusData)
	},
}

// GetTaobaoTbkScVegasSendStatusData() 从对象池中获取TaobaoTbkScVegasSendStatusData
func GetTaobaoTbkScVegasSendStatusData() *TaobaoTbkScVegasSendStatusData {
	return poolTaobaoTbkScVegasSendStatusData.Get().(*TaobaoTbkScVegasSendStatusData)
}

// ReleaseTaobaoTbkScVegasSendStatusData 释放TaobaoTbkScVegasSendStatusData
func ReleaseTaobaoTbkScVegasSendStatusData(v *TaobaoTbkScVegasSendStatusData) {
	v.ResultList = v.ResultList[:0]
	poolTaobaoTbkScVegasSendStatusData.Put(v)
}
