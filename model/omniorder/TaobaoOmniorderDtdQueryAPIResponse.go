package omniorder

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOmniorderDtdQueryAPIResponse 门店自送根据核销码查订单 API返回值
// taobao.omniorder.dtd.query
//
// 门店自送根据核销码码查询订单信息
type TaobaoOmniorderDtdQueryAPIResponse struct {
	model.CommonResponse
	TaobaoOmniorderDtdQueryAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoOmniorderDtdQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoOmniorderDtdQueryAPIResponseModel).Reset()
}

// TaobaoOmniorderDtdQueryAPIResponseModel is 门店自送根据核销码查订单 成功返回结果
type TaobaoOmniorderDtdQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"omniorder_dtd_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误码，为0表示成功，非0表示失败
	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 错误信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// data
	Data *Door2doorQueryResult `json:"data,omitempty" xml:"data,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoOmniorderDtdQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ResultCode = ""
	m.Message = ""
	m.Data = nil
}

var poolTaobaoOmniorderDtdQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoOmniorderDtdQueryAPIResponse)
	},
}

// GetTaobaoOmniorderDtdQueryAPIResponse 从 sync.Pool 获取 TaobaoOmniorderDtdQueryAPIResponse
func GetTaobaoOmniorderDtdQueryAPIResponse() *TaobaoOmniorderDtdQueryAPIResponse {
	return poolTaobaoOmniorderDtdQueryAPIResponse.Get().(*TaobaoOmniorderDtdQueryAPIResponse)
}

// ReleaseTaobaoOmniorderDtdQueryAPIResponse 将 TaobaoOmniorderDtdQueryAPIResponse 保存到 sync.Pool
func ReleaseTaobaoOmniorderDtdQueryAPIResponse(v *TaobaoOmniorderDtdQueryAPIResponse) {
	v.Reset()
	poolTaobaoOmniorderDtdQueryAPIResponse.Put(v)
}
