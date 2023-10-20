package xhotelitem

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelRateplanUpdateAPIResponse 价格计划rateplan更新或添加 API返回值
// taobao.xhotel.rateplan.update
//
// 酒店产品库rateplan更新或添加
type TaobaoXhotelRateplanUpdateAPIResponse struct {
	model.CommonResponse
	TaobaoXhotelRateplanUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoXhotelRateplanUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoXhotelRateplanUpdateAPIResponseModel).Reset()
}

// TaobaoXhotelRateplanUpdateAPIResponseModel is 价格计划rateplan更新或添加 成功返回结果
type TaobaoXhotelRateplanUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"xhotel_rateplan_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 修改的rp id
	Rpid int64 `json:"rpid,omitempty" xml:"rpid,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoXhotelRateplanUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Rpid = 0
}

var poolTaobaoXhotelRateplanUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoXhotelRateplanUpdateAPIResponse)
	},
}

// GetTaobaoXhotelRateplanUpdateAPIResponse 从 sync.Pool 获取 TaobaoXhotelRateplanUpdateAPIResponse
func GetTaobaoXhotelRateplanUpdateAPIResponse() *TaobaoXhotelRateplanUpdateAPIResponse {
	return poolTaobaoXhotelRateplanUpdateAPIResponse.Get().(*TaobaoXhotelRateplanUpdateAPIResponse)
}

// ReleaseTaobaoXhotelRateplanUpdateAPIResponse 将 TaobaoXhotelRateplanUpdateAPIResponse 保存到 sync.Pool
func ReleaseTaobaoXhotelRateplanUpdateAPIResponse(v *TaobaoXhotelRateplanUpdateAPIResponse) {
	v.Reset()
	poolTaobaoXhotelRateplanUpdateAPIResponse.Put(v)
}
