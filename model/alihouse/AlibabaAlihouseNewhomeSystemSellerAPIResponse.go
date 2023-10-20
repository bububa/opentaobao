package alihouse

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeSystemSellerAPIResponse 商品发布授权 API返回值
// alibaba.alihouse.newhome.system.seller
//
// 商品发布授权
type AlibabaAlihouseNewhomeSystemSellerAPIResponse struct {
	model.CommonResponse
	AlibabaAlihouseNewhomeSystemSellerAPIResponseModel
}

// AlibabaAlihouseNewhomeSystemSellerAPIResponseModel is 商品发布授权 成功返回结果
type AlibabaAlihouseNewhomeSystemSellerAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_newhome_system_seller_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaAlihouseNewhomeSystemSellerResult `json:"result,omitempty" xml:"result,omitempty"`
}
