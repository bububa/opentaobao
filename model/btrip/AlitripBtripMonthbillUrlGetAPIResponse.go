package btrip

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripMonthbillUrlGetAPIResponse 月账单数据查询 API返回值
// alitrip.btrip.monthbill.url.get
//
// 月账单数据查询
type AlitripBtripMonthbillUrlGetAPIResponse struct {
	model.CommonResponse
	AlitripBtripMonthbillUrlGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripBtripMonthbillUrlGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripBtripMonthbillUrlGetAPIResponseModel).Reset()
}

// AlitripBtripMonthbillUrlGetAPIResponseModel is 月账单数据查询 成功返回结果
type AlitripBtripMonthbillUrlGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_btrip_monthbill_url_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果对象
	Results []OpenAccountRs `json:"results,omitempty" xml:"results>open_account_rs,omitempty"`
	// 结果信息
	ResultMsg string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 结果码
	ResultCode int64 `json:"result_code,omitempty" xml:"result_code,omitempty"`
}

// Reset 清空结构体
func (m *AlitripBtripMonthbillUrlGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Results = m.Results[:0]
	m.ResultMsg = ""
	m.ResultCode = 0
}

var poolAlitripBtripMonthbillUrlGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripBtripMonthbillUrlGetAPIResponse)
	},
}

// GetAlitripBtripMonthbillUrlGetAPIResponse 从 sync.Pool 获取 AlitripBtripMonthbillUrlGetAPIResponse
func GetAlitripBtripMonthbillUrlGetAPIResponse() *AlitripBtripMonthbillUrlGetAPIResponse {
	return poolAlitripBtripMonthbillUrlGetAPIResponse.Get().(*AlitripBtripMonthbillUrlGetAPIResponse)
}

// ReleaseAlitripBtripMonthbillUrlGetAPIResponse 将 AlitripBtripMonthbillUrlGetAPIResponse 保存到 sync.Pool
func ReleaseAlitripBtripMonthbillUrlGetAPIResponse(v *AlitripBtripMonthbillUrlGetAPIResponse) {
	v.Reset()
	poolAlitripBtripMonthbillUrlGetAPIResponse.Put(v)
}
