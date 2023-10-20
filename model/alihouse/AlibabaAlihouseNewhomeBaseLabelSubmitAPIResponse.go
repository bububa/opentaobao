package alihouse

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeBaseLabelSubmitAPIResponse 提交标签库 API返回值
// alibaba.alihouse.newhome.base.label.submit
//
// 提交标签库
type AlibabaAlihouseNewhomeBaseLabelSubmitAPIResponse struct {
	model.CommonResponse
	AlibabaAlihouseNewhomeBaseLabelSubmitAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihouseNewhomeBaseLabelSubmitAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihouseNewhomeBaseLabelSubmitAPIResponseModel).Reset()
}

// AlibabaAlihouseNewhomeBaseLabelSubmitAPIResponseModel is 提交标签库 成功返回结果
type AlibabaAlihouseNewhomeBaseLabelSubmitAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_newhome_base_label_submit_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaAlihouseNewhomeBaseLabelSubmitResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihouseNewhomeBaseLabelSubmitAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihouseNewhomeBaseLabelSubmitAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseNewhomeBaseLabelSubmitAPIResponse)
	},
}

// GetAlibabaAlihouseNewhomeBaseLabelSubmitAPIResponse 从 sync.Pool 获取 AlibabaAlihouseNewhomeBaseLabelSubmitAPIResponse
func GetAlibabaAlihouseNewhomeBaseLabelSubmitAPIResponse() *AlibabaAlihouseNewhomeBaseLabelSubmitAPIResponse {
	return poolAlibabaAlihouseNewhomeBaseLabelSubmitAPIResponse.Get().(*AlibabaAlihouseNewhomeBaseLabelSubmitAPIResponse)
}

// ReleaseAlibabaAlihouseNewhomeBaseLabelSubmitAPIResponse 将 AlibabaAlihouseNewhomeBaseLabelSubmitAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihouseNewhomeBaseLabelSubmitAPIResponse(v *AlibabaAlihouseNewhomeBaseLabelSubmitAPIResponse) {
	v.Reset()
	poolAlibabaAlihouseNewhomeBaseLabelSubmitAPIResponse.Put(v)
}
