package alihouse

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeShopconfigDetailSubmitAPIResponse 店铺配置信息接口 API返回值
// alibaba.alihouse.newhome.shopconfig.detail.submit
//
// 提供店铺配置的能力
type AlibabaAlihouseNewhomeShopconfigDetailSubmitAPIResponse struct {
	model.CommonResponse
	AlibabaAlihouseNewhomeShopconfigDetailSubmitAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihouseNewhomeShopconfigDetailSubmitAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihouseNewhomeShopconfigDetailSubmitAPIResponseModel).Reset()
}

// AlibabaAlihouseNewhomeShopconfigDetailSubmitAPIResponseModel is 店铺配置信息接口 成功返回结果
type AlibabaAlihouseNewhomeShopconfigDetailSubmitAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_newhome_shopconfig_detail_submit_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaAlihouseNewhomeShopconfigDetailSubmitResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihouseNewhomeShopconfigDetailSubmitAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihouseNewhomeShopconfigDetailSubmitAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseNewhomeShopconfigDetailSubmitAPIResponse)
	},
}

// GetAlibabaAlihouseNewhomeShopconfigDetailSubmitAPIResponse 从 sync.Pool 获取 AlibabaAlihouseNewhomeShopconfigDetailSubmitAPIResponse
func GetAlibabaAlihouseNewhomeShopconfigDetailSubmitAPIResponse() *AlibabaAlihouseNewhomeShopconfigDetailSubmitAPIResponse {
	return poolAlibabaAlihouseNewhomeShopconfigDetailSubmitAPIResponse.Get().(*AlibabaAlihouseNewhomeShopconfigDetailSubmitAPIResponse)
}

// ReleaseAlibabaAlihouseNewhomeShopconfigDetailSubmitAPIResponse 将 AlibabaAlihouseNewhomeShopconfigDetailSubmitAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihouseNewhomeShopconfigDetailSubmitAPIResponse(v *AlibabaAlihouseNewhomeShopconfigDetailSubmitAPIResponse) {
	v.Reset()
	poolAlibabaAlihouseNewhomeShopconfigDetailSubmitAPIResponse.Put(v)
}
