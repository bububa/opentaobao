package qimen

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoqimeninventoryrulecreateAPIResponse 渠道间库存规则设置接口 API返回值
// taobao.qimen.inventoryrule.create
//
// 渠道间库存规则设置
type TaobaoqimeninventoryrulecreateAPIResponse struct {
	model.CommonResponse
	TaobaoqimeninventoryrulecreateAPIResponseModel
}

// TaobaoqimeninventoryrulecreateAPIResponseModel is 渠道间库存规则设置接口 成功返回结果
type TaobaoqimeninventoryrulecreateAPIResponseModel struct {
	XMLName xml.Name `xml:"qimen_inventoryrule_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	//
	Response *ResponseDo `json:"response,omitempty" xml:"response,omitempty"`
}
