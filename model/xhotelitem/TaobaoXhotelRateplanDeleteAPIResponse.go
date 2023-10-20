package xhotelitem

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelRateplanDeleteAPIResponse 价格计划rateplan删除 API返回值
// taobao.xhotel.rateplan.delete
//
// 酒店产品库rateplan删除，同时删除级联的rate，请谨慎使用
type TaobaoXhotelRateplanDeleteAPIResponse struct {
	model.CommonResponse
	TaobaoXhotelRateplanDeleteAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoXhotelRateplanDeleteAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoXhotelRateplanDeleteAPIResponseModel).Reset()
}

// TaobaoXhotelRateplanDeleteAPIResponseModel is 价格计划rateplan删除 成功返回结果
type TaobaoXhotelRateplanDeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"xhotel_rateplan_delete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TaobaoXhotelRateplanDeleteResultSet `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoXhotelRateplanDeleteAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoXhotelRateplanDeleteAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoXhotelRateplanDeleteAPIResponse)
	},
}

// GetTaobaoXhotelRateplanDeleteAPIResponse 从 sync.Pool 获取 TaobaoXhotelRateplanDeleteAPIResponse
func GetTaobaoXhotelRateplanDeleteAPIResponse() *TaobaoXhotelRateplanDeleteAPIResponse {
	return poolTaobaoXhotelRateplanDeleteAPIResponse.Get().(*TaobaoXhotelRateplanDeleteAPIResponse)
}

// ReleaseTaobaoXhotelRateplanDeleteAPIResponse 将 TaobaoXhotelRateplanDeleteAPIResponse 保存到 sync.Pool
func ReleaseTaobaoXhotelRateplanDeleteAPIResponse(v *TaobaoXhotelRateplanDeleteAPIResponse) {
	v.Reset()
	poolTaobaoXhotelRateplanDeleteAPIResponse.Put(v)
}
