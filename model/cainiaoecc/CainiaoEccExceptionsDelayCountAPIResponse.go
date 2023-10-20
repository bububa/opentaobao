package cainiaoecc

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// CainiaoEccExceptionsDelayCountAPIResponse 菜鸟控制塔包裹滞留异常统计信息获取 API返回值
// cainiao.ecc.exceptions.delay.count
//
// 菜鸟控制塔包裹滞留异常统计信息获取
type CainiaoEccExceptionsDelayCountAPIResponse struct {
	model.CommonResponse
	CainiaoEccExceptionsDelayCountAPIResponseModel
}

// Reset 清空结构体
func (m *CainiaoEccExceptionsDelayCountAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.CainiaoEccExceptionsDelayCountAPIResponseModel).Reset()
}

// CainiaoEccExceptionsDelayCountAPIResponseModel is 菜鸟控制塔包裹滞留异常统计信息获取 成功返回结果
type CainiaoEccExceptionsDelayCountAPIResponseModel struct {
	XMLName xml.Name `xml:"cainiao_ecc_exceptions_delay_count_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *SingleResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *CainiaoEccExceptionsDelayCountAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolCainiaoEccExceptionsDelayCountAPIResponse = sync.Pool{
	New: func() any {
		return new(CainiaoEccExceptionsDelayCountAPIResponse)
	},
}

// GetCainiaoEccExceptionsDelayCountAPIResponse 从 sync.Pool 获取 CainiaoEccExceptionsDelayCountAPIResponse
func GetCainiaoEccExceptionsDelayCountAPIResponse() *CainiaoEccExceptionsDelayCountAPIResponse {
	return poolCainiaoEccExceptionsDelayCountAPIResponse.Get().(*CainiaoEccExceptionsDelayCountAPIResponse)
}

// ReleaseCainiaoEccExceptionsDelayCountAPIResponse 将 CainiaoEccExceptionsDelayCountAPIResponse 保存到 sync.Pool
func ReleaseCainiaoEccExceptionsDelayCountAPIResponse(v *CainiaoEccExceptionsDelayCountAPIResponse) {
	v.Reset()
	poolCainiaoEccExceptionsDelayCountAPIResponse.Put(v)
}
