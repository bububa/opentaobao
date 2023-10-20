package xhotelitem

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelRateDeleteAPIResponse rate删除接口 API返回值
// taobao.xhotel.rate.delete
//
// 酒店产品库rate删除
type TaobaoXhotelRateDeleteAPIResponse struct {
	model.CommonResponse
	TaobaoXhotelRateDeleteAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoXhotelRateDeleteAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoXhotelRateDeleteAPIResponseModel).Reset()
}

// TaobaoXhotelRateDeleteAPIResponseModel is rate删除接口 成功返回结果
type TaobaoXhotelRateDeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"xhotel_rate_delete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TaobaoXhotelRateDeleteResultSet `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoXhotelRateDeleteAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoXhotelRateDeleteAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoXhotelRateDeleteAPIResponse)
	},
}

// GetTaobaoXhotelRateDeleteAPIResponse 从 sync.Pool 获取 TaobaoXhotelRateDeleteAPIResponse
func GetTaobaoXhotelRateDeleteAPIResponse() *TaobaoXhotelRateDeleteAPIResponse {
	return poolTaobaoXhotelRateDeleteAPIResponse.Get().(*TaobaoXhotelRateDeleteAPIResponse)
}

// ReleaseTaobaoXhotelRateDeleteAPIResponse 将 TaobaoXhotelRateDeleteAPIResponse 保存到 sync.Pool
func ReleaseTaobaoXhotelRateDeleteAPIResponse(v *TaobaoXhotelRateDeleteAPIResponse) {
	v.Reset()
	poolTaobaoXhotelRateDeleteAPIResponse.Put(v)
}
