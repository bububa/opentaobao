package drugtrace

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytDrStorageuploadAPIResponse 疫苗存储温度上传 API返回值
// alibaba.alihealth.drug.kyt.dr.storageupload
//
// 疫苗存储温度上传
type AlibabaAlihealthDrugKytDrStorageuploadAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugKytDrStorageuploadAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytDrStorageuploadAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthDrugKytDrStorageuploadAPIResponseModel).Reset()
}

// AlibabaAlihealthDrugKytDrStorageuploadAPIResponseModel is 疫苗存储温度上传 成功返回结果
type AlibabaAlihealthDrugKytDrStorageuploadAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_dr_storageupload_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaAlihealthDrugKytDrStorageuploadResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthDrugKytDrStorageuploadAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthDrugKytDrStorageuploadAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytDrStorageuploadAPIResponse)
	},
}

// GetAlibabaAlihealthDrugKytDrStorageuploadAPIResponse 从 sync.Pool 获取 AlibabaAlihealthDrugKytDrStorageuploadAPIResponse
func GetAlibabaAlihealthDrugKytDrStorageuploadAPIResponse() *AlibabaAlihealthDrugKytDrStorageuploadAPIResponse {
	return poolAlibabaAlihealthDrugKytDrStorageuploadAPIResponse.Get().(*AlibabaAlihealthDrugKytDrStorageuploadAPIResponse)
}

// ReleaseAlibabaAlihealthDrugKytDrStorageuploadAPIResponse 将 AlibabaAlihealthDrugKytDrStorageuploadAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthDrugKytDrStorageuploadAPIResponse(v *AlibabaAlihealthDrugKytDrStorageuploadAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthDrugKytDrStorageuploadAPIResponse.Put(v)
}
