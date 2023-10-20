package cainiaoecc

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// CainiaoEccExceptionsDelayGetAPIResponse 菜鸟控制塔包裹滞留异常信息获取 API返回值
// cainiao.ecc.exceptions.delay.get
//
// 菜鸟控制塔包裹滞留异常信息获取
type CainiaoEccExceptionsDelayGetAPIResponse struct {
	model.CommonResponse
	CainiaoEccExceptionsDelayGetAPIResponseModel
}

// Reset 清空结构体
func (m *CainiaoEccExceptionsDelayGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.CainiaoEccExceptionsDelayGetAPIResponseModel).Reset()
}

// CainiaoEccExceptionsDelayGetAPIResponseModel is 菜鸟控制塔包裹滞留异常信息获取 成功返回结果
type CainiaoEccExceptionsDelayGetAPIResponseModel struct {
	XMLName xml.Name `xml:"cainiao_ecc_exceptions_delay_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *SingleResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *CainiaoEccExceptionsDelayGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolCainiaoEccExceptionsDelayGetAPIResponse = sync.Pool{
	New: func() any {
		return new(CainiaoEccExceptionsDelayGetAPIResponse)
	},
}

// GetCainiaoEccExceptionsDelayGetAPIResponse 从 sync.Pool 获取 CainiaoEccExceptionsDelayGetAPIResponse
func GetCainiaoEccExceptionsDelayGetAPIResponse() *CainiaoEccExceptionsDelayGetAPIResponse {
	return poolCainiaoEccExceptionsDelayGetAPIResponse.Get().(*CainiaoEccExceptionsDelayGetAPIResponse)
}

// ReleaseCainiaoEccExceptionsDelayGetAPIResponse 将 CainiaoEccExceptionsDelayGetAPIResponse 保存到 sync.Pool
func ReleaseCainiaoEccExceptionsDelayGetAPIResponse(v *CainiaoEccExceptionsDelayGetAPIResponse) {
	v.Reset()
	poolCainiaoEccExceptionsDelayGetAPIResponse.Put(v)
}
