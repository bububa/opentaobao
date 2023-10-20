package drugtrace

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthdrugkytscqydelbillinfoAPIResponse 根据单据号删除单据 API返回值
// alibaba.alihealth.drug.kyt.scqy.delbillinfo
//
// 根据单据号删除单据
type AlibabaalihealthdrugkytscqydelbillinfoAPIResponse struct {
	model.CommonResponse
	AlibabaalihealthdrugkytscqydelbillinfoAPIResponseModel
}

// AlibabaalihealthdrugkytscqydelbillinfoAPIResponseModel is 根据单据号删除单据 成功返回结果
type AlibabaalihealthdrugkytscqydelbillinfoAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_scqy_delbillinfo_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 最外层结果
	Result *AlibabaalihealthdrugkytscqydelbillinfoResultModel `json:"result,omitempty" xml:"result,omitempty"`
}
