package drugtrace

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthdruguploadextinfoAPIResponse 中药饮片及器械对接 API返回值
// alibaba.alihealth.drug.upload.extinfo
//
// 中药饮片及器械对接
type AlibabaalihealthdruguploadextinfoAPIResponse struct {
	model.CommonResponse
	AlibabaalihealthdruguploadextinfoAPIResponseModel
}

// AlibabaalihealthdruguploadextinfoAPIResponseModel is 中药饮片及器械对接 成功返回结果
type AlibabaalihealthdruguploadextinfoAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_upload_extinfo_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *AlibabaalihealthdruguploadextinfoResultModel `json:"result,omitempty" xml:"result,omitempty"`
}
