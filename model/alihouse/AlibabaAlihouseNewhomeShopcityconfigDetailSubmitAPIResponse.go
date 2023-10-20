package alihouse

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeShopcityconfigDetailSubmitAPIResponse 城市配置信息接口 API返回值
// alibaba.alihouse.newhome.shopcityconfig.detail.submit
//
// 上传城市配置信息
type AlibabaAlihouseNewhomeShopcityconfigDetailSubmitAPIResponse struct {
	model.CommonResponse
	AlibabaAlihouseNewhomeShopcityconfigDetailSubmitAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihouseNewhomeShopcityconfigDetailSubmitAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihouseNewhomeShopcityconfigDetailSubmitAPIResponseModel).Reset()
}

// AlibabaAlihouseNewhomeShopcityconfigDetailSubmitAPIResponseModel is 城市配置信息接口 成功返回结果
type AlibabaAlihouseNewhomeShopcityconfigDetailSubmitAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_newhome_shopcityconfig_detail_submit_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaAlihouseNewhomeShopcityconfigDetailSubmitResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihouseNewhomeShopcityconfigDetailSubmitAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihouseNewhomeShopcityconfigDetailSubmitAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseNewhomeShopcityconfigDetailSubmitAPIResponse)
	},
}

// GetAlibabaAlihouseNewhomeShopcityconfigDetailSubmitAPIResponse 从 sync.Pool 获取 AlibabaAlihouseNewhomeShopcityconfigDetailSubmitAPIResponse
func GetAlibabaAlihouseNewhomeShopcityconfigDetailSubmitAPIResponse() *AlibabaAlihouseNewhomeShopcityconfigDetailSubmitAPIResponse {
	return poolAlibabaAlihouseNewhomeShopcityconfigDetailSubmitAPIResponse.Get().(*AlibabaAlihouseNewhomeShopcityconfigDetailSubmitAPIResponse)
}

// ReleaseAlibabaAlihouseNewhomeShopcityconfigDetailSubmitAPIResponse 将 AlibabaAlihouseNewhomeShopcityconfigDetailSubmitAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihouseNewhomeShopcityconfigDetailSubmitAPIResponse(v *AlibabaAlihouseNewhomeShopcityconfigDetailSubmitAPIResponse) {
	v.Reset()
	poolAlibabaAlihouseNewhomeShopcityconfigDetailSubmitAPIResponse.Put(v)
}
