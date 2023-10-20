package alihealth2

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthdentalstoreinsertorupdateAPIResponse ISV新增/修改口腔门店 API返回值
// alibaba.alihealth.dental.store.insertorupdate
//
// ISV新增/修改口腔门店
type AlibabaalihealthdentalstoreinsertorupdateAPIResponse struct {
	model.CommonResponse
	AlibabaalihealthdentalstoreinsertorupdateAPIResponseModel
}

// AlibabaalihealthdentalstoreinsertorupdateAPIResponseModel is ISV新增/修改口腔门店 成功返回结果
type AlibabaalihealthdentalstoreinsertorupdateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_dental_store_insertorupdate_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *AlibabaalihealthdentalstoreinsertorupdateMtopResult `json:"result,omitempty" xml:"result,omitempty"`
}
