package newretail

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaItApAddressGetAPIResponse
getApAddressByMacNew API返回值
alibaba.it.ap.address.get

根据ap 的mac地址查询ap的结构化位置信息 */
type AlibabaItApAddressGetAPIResponse struct {
	model.CommonResponse
	AlibabaItApAddressGetAPIResponseModel
}

// AlibabaItApAddressGetAPIResponseModel is getApAddressByMacNew 成功返回结果
type AlibabaItApAddressGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_it_ap_address_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaItApAddressGetResult `json:"result,omitempty" xml:"result,omitempty"`
}
