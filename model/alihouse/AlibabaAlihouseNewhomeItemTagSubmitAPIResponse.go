package alihouse

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihousenewhomeitemtagsubmitAPIResponse ETC上翻商品打标接口 API返回值
// alibaba.alihouse.newhome.item.tag.submit
//
// ETC上翻商品打标接口
type AlibabaalihousenewhomeitemtagsubmitAPIResponse struct {
	model.CommonResponse
	AlibabaalihousenewhomeitemtagsubmitAPIResponseModel
}

// AlibabaalihousenewhomeitemtagsubmitAPIResponseModel is ETC上翻商品打标接口 成功返回结果
type AlibabaalihousenewhomeitemtagsubmitAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_newhome_item_tag_submit_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *AlibabaalihousenewhomeitemtagsubmitResult `json:"result,omitempty" xml:"result,omitempty"`
}
