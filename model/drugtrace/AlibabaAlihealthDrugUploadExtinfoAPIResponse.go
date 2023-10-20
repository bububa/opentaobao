package drugtrace

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugUploadExtinfoAPIResponse 中药饮片及器械对接 API返回值
// alibaba.alihealth.drug.upload.extinfo
//
// 中药饮片及器械对接
type AlibabaAlihealthDrugUploadExtinfoAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugUploadExtinfoAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugUploadExtinfoAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthDrugUploadExtinfoAPIResponseModel).Reset()
}

// AlibabaAlihealthDrugUploadExtinfoAPIResponseModel is 中药饮片及器械对接 成功返回结果
type AlibabaAlihealthDrugUploadExtinfoAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_upload_extinfo_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *AlibabaAlihealthDrugUploadExtinfoResultModel `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugUploadExtinfoAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthDrugUploadExtinfoAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugUploadExtinfoAPIResponse)
	},
}

// GetAlibabaAlihealthDrugUploadExtinfoAPIResponse 从 sync.Pool 获取 AlibabaAlihealthDrugUploadExtinfoAPIResponse
func GetAlibabaAlihealthDrugUploadExtinfoAPIResponse() *AlibabaAlihealthDrugUploadExtinfoAPIResponse {
	return poolAlibabaAlihealthDrugUploadExtinfoAPIResponse.Get().(*AlibabaAlihealthDrugUploadExtinfoAPIResponse)
}

// ReleaseAlibabaAlihealthDrugUploadExtinfoAPIResponse 将 AlibabaAlihealthDrugUploadExtinfoAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthDrugUploadExtinfoAPIResponse(v *AlibabaAlihealthDrugUploadExtinfoAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthDrugUploadExtinfoAPIResponse.Put(v)
}
