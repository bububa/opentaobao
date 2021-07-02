package xhotelitem

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelEntityConfigAPIResponse 飞猪商品各实体通用配置 API返回值
// taobao.xhotel.entity.config
//
// 飞猪商品各实体通用配置服务
type TaobaoXhotelEntityConfigAPIResponse struct {
	model.CommonResponse
	TaobaoXhotelEntityConfigAPIResponseModel
}

// TaobaoXhotelEntityConfigAPIResponseModel is 飞猪商品各实体通用配置 成功返回结果
type TaobaoXhotelEntityConfigAPIResponseModel struct {
	XMLName xml.Name `xml:"xhotel_entity_config_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 请勿关注该值
	Result int64 `json:"result,omitempty" xml:"result,omitempty"`
}
