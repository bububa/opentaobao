package alihouse

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseExistinghomeRegionInfoSubmitAPIResponse 商圈专家信息同步 API返回值
// alibaba.alihouse.existinghome.region.info.submit
//
// 商圈专家信息同步
type AlibabaAlihouseExistinghomeRegionInfoSubmitAPIResponse struct {
	model.CommonResponse
	AlibabaAlihouseExistinghomeRegionInfoSubmitAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihouseExistinghomeRegionInfoSubmitAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihouseExistinghomeRegionInfoSubmitAPIResponseModel).Reset()
}

// AlibabaAlihouseExistinghomeRegionInfoSubmitAPIResponseModel is 商圈专家信息同步 成功返回结果
type AlibabaAlihouseExistinghomeRegionInfoSubmitAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_existinghome_region_info_submit_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *AlibabaAlihouseExistinghomeRegionInfoSubmitResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihouseExistinghomeRegionInfoSubmitAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihouseExistinghomeRegionInfoSubmitAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseExistinghomeRegionInfoSubmitAPIResponse)
	},
}

// GetAlibabaAlihouseExistinghomeRegionInfoSubmitAPIResponse 从 sync.Pool 获取 AlibabaAlihouseExistinghomeRegionInfoSubmitAPIResponse
func GetAlibabaAlihouseExistinghomeRegionInfoSubmitAPIResponse() *AlibabaAlihouseExistinghomeRegionInfoSubmitAPIResponse {
	return poolAlibabaAlihouseExistinghomeRegionInfoSubmitAPIResponse.Get().(*AlibabaAlihouseExistinghomeRegionInfoSubmitAPIResponse)
}

// ReleaseAlibabaAlihouseExistinghomeRegionInfoSubmitAPIResponse 将 AlibabaAlihouseExistinghomeRegionInfoSubmitAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihouseExistinghomeRegionInfoSubmitAPIResponse(v *AlibabaAlihouseExistinghomeRegionInfoSubmitAPIResponse) {
	v.Reset()
	poolAlibabaAlihouseExistinghomeRegionInfoSubmitAPIResponse.Put(v)
}
