package newretail

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaitapaddresssetAPIResponse setApAddressNew API返回值
// alibaba.it.ap.address.set
//
// 该接口可为ISV系统提供 ap位置信息维护的功能
type AlibabaitapaddresssetAPIResponse struct {
	model.CommonResponse
	AlibabaitapaddresssetAPIResponseModel
}

// AlibabaitapaddresssetAPIResponseModel is setApAddressNew 成功返回结果
type AlibabaitapaddresssetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_it_ap_address_set_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaitapaddresssetResult `json:"result,omitempty" xml:"result,omitempty"`
}
