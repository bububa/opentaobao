package examination

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthExaminationItemsPublishAPIResponse 单项/加项包信息同步 API返回值
// alibaba.alihealth.examination.items.publish
//
// 体检机构对接_单项/加项包信息发布／更新
type AlibabaAlihealthExaminationItemsPublishAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthExaminationItemsPublishAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthExaminationItemsPublishAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthExaminationItemsPublishAPIResponseModel).Reset()
}

// AlibabaAlihealthExaminationItemsPublishAPIResponseModel is 单项/加项包信息同步 成功返回结果
type AlibabaAlihealthExaminationItemsPublishAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_examination_items_publish_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *ServiceResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthExaminationItemsPublishAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthExaminationItemsPublishAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthExaminationItemsPublishAPIResponse)
	},
}

// GetAlibabaAlihealthExaminationItemsPublishAPIResponse 从 sync.Pool 获取 AlibabaAlihealthExaminationItemsPublishAPIResponse
func GetAlibabaAlihealthExaminationItemsPublishAPIResponse() *AlibabaAlihealthExaminationItemsPublishAPIResponse {
	return poolAlibabaAlihealthExaminationItemsPublishAPIResponse.Get().(*AlibabaAlihealthExaminationItemsPublishAPIResponse)
}

// ReleaseAlibabaAlihealthExaminationItemsPublishAPIResponse 将 AlibabaAlihealthExaminationItemsPublishAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthExaminationItemsPublishAPIResponse(v *AlibabaAlihealthExaminationItemsPublishAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthExaminationItemsPublishAPIResponse.Put(v)
}
