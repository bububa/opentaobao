package tbk

import (
	"sync"
)

// TaobaoTbkDgVegasSendStatusData 结构体
type TaobaoTbkDgVegasSendStatusData struct {
	// 返回结果封装对象
	ResultList []TaobaoTbkDgVegasSendStatusMapData `json:"result_list,omitempty" xml:"result_list>taobao_tbk_dg_vegas_send_status_map_data,omitempty"`
}

var poolTaobaoTbkDgVegasSendStatusData = sync.Pool{
	New: func() any {
		return new(TaobaoTbkDgVegasSendStatusData)
	},
}

// GetTaobaoTbkDgVegasSendStatusData() 从对象池中获取TaobaoTbkDgVegasSendStatusData
func GetTaobaoTbkDgVegasSendStatusData() *TaobaoTbkDgVegasSendStatusData {
	return poolTaobaoTbkDgVegasSendStatusData.Get().(*TaobaoTbkDgVegasSendStatusData)
}

// ReleaseTaobaoTbkDgVegasSendStatusData 释放TaobaoTbkDgVegasSendStatusData
func ReleaseTaobaoTbkDgVegasSendStatusData(v *TaobaoTbkDgVegasSendStatusData) {
	v.ResultList = v.ResultList[:0]
	poolTaobaoTbkDgVegasSendStatusData.Put(v)
}
