package xhotelitem

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelRateplanAddAPIResponse 酒店产品库rateplan添加 API返回值
// taobao.xhotel.rateplan.add
//
// 酒店产品库rateplan
type TaobaoXhotelRateplanAddAPIResponse struct {
	model.CommonResponse
	TaobaoXhotelRateplanAddAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoXhotelRateplanAddAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoXhotelRateplanAddAPIResponseModel).Reset()
}

// TaobaoXhotelRateplanAddAPIResponseModel is 酒店产品库rateplan添加 成功返回结果
type TaobaoXhotelRateplanAddAPIResponseModel struct {
	XMLName xml.Name `xml:"xhotel_rateplan_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 生成的rp id
	Rpid int64 `json:"rpid,omitempty" xml:"rpid,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoXhotelRateplanAddAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Rpid = 0
}

var poolTaobaoXhotelRateplanAddAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoXhotelRateplanAddAPIResponse)
	},
}

// GetTaobaoXhotelRateplanAddAPIResponse 从 sync.Pool 获取 TaobaoXhotelRateplanAddAPIResponse
func GetTaobaoXhotelRateplanAddAPIResponse() *TaobaoXhotelRateplanAddAPIResponse {
	return poolTaobaoXhotelRateplanAddAPIResponse.Get().(*TaobaoXhotelRateplanAddAPIResponse)
}

// ReleaseTaobaoXhotelRateplanAddAPIResponse 将 TaobaoXhotelRateplanAddAPIResponse 保存到 sync.Pool
func ReleaseTaobaoXhotelRateplanAddAPIResponse(v *TaobaoXhotelRateplanAddAPIResponse) {
	v.Reset()
	poolTaobaoXhotelRateplanAddAPIResponse.Put(v)
}
