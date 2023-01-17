package tblogistics

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoDeliveryTemplateUpdateAPIResponse 修改运费模板 API返回值
// taobao.delivery.template.update
//
// 修改运费模板
type TaobaoDeliveryTemplateUpdateAPIResponse struct {
	model.CommonResponse
	TaobaoDeliveryTemplateUpdateAPIResponseModel
}

// TaobaoDeliveryTemplateUpdateAPIResponseModel is 修改运费模板 成功返回结果
type TaobaoDeliveryTemplateUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"delivery_template_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 表示修改是否成功
	Complete bool `json:"complete,omitempty" xml:"complete,omitempty"`
}
