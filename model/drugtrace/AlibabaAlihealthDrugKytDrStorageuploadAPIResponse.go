package drugtrace

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthdrugkytdrstorageuploadAPIResponse 疫苗存储温度上传 API返回值
// alibaba.alihealth.drug.kyt.dr.storageupload
//
// 疫苗存储温度上传
type AlibabaalihealthdrugkytdrstorageuploadAPIResponse struct {
	model.CommonResponse
	AlibabaalihealthdrugkytdrstorageuploadAPIResponseModel
}

// AlibabaalihealthdrugkytdrstorageuploadAPIResponseModel is 疫苗存储温度上传 成功返回结果
type AlibabaalihealthdrugkytdrstorageuploadAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_dr_storageupload_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaalihealthdrugkytdrstorageuploadResult `json:"result,omitempty" xml:"result,omitempty"`
}
