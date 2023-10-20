package xhotelitem

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelRatesUpdateAPIResponse 价格推送接口（批量全量） API返回值
// taobao.xhotel.rates.update
//
// 酒店产品库rate批量更新房态信息
type TaobaoXhotelRatesUpdateAPIResponse struct {
	model.CommonResponse
	TaobaoXhotelRatesUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoXhotelRatesUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoXhotelRatesUpdateAPIResponseModel).Reset()
}

// TaobaoXhotelRatesUpdateAPIResponseModel is 价格推送接口（批量全量） 成功返回结果
type TaobaoXhotelRatesUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"xhotel_rates_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// gid_and_rateplan_ids
	GidAndRpids []string `json:"gid_and_rpids,omitempty" xml:"gid_and_rpids>string,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoXhotelRatesUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.GidAndRpids = m.GidAndRpids[:0]
}

var poolTaobaoXhotelRatesUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoXhotelRatesUpdateAPIResponse)
	},
}

// GetTaobaoXhotelRatesUpdateAPIResponse 从 sync.Pool 获取 TaobaoXhotelRatesUpdateAPIResponse
func GetTaobaoXhotelRatesUpdateAPIResponse() *TaobaoXhotelRatesUpdateAPIResponse {
	return poolTaobaoXhotelRatesUpdateAPIResponse.Get().(*TaobaoXhotelRatesUpdateAPIResponse)
}

// ReleaseTaobaoXhotelRatesUpdateAPIResponse 将 TaobaoXhotelRatesUpdateAPIResponse 保存到 sync.Pool
func ReleaseTaobaoXhotelRatesUpdateAPIResponse(v *TaobaoXhotelRatesUpdateAPIResponse) {
	v.Reset()
	poolTaobaoXhotelRatesUpdateAPIResponse.Put(v)
}
