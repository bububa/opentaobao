package tblogistics

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaodeliverytemplategetAPIResponse 获取用户指定运费模板信息 API返回值
// taobao.delivery.template.get
//
// 获取用户指定运费模板信息
type TaobaodeliverytemplategetAPIResponse struct {
	model.CommonResponse
	TaobaodeliverytemplategetAPIResponseModel
}

// TaobaodeliverytemplategetAPIResponseModel is 获取用户指定运费模板信息 成功返回结果
type TaobaodeliverytemplategetAPIResponseModel struct {
	XMLName xml.Name `xml:"delivery_template_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 运费模板列表
	DeliveryTemplates []DeliveryTemplate `json:"delivery_templates,omitempty" xml:"delivery_templates>delivery_template,omitempty"`
	// 获得到符合条件的结果总数
	TotalResults int64 `json:"total_results,omitempty" xml:"total_results,omitempty"`
}
