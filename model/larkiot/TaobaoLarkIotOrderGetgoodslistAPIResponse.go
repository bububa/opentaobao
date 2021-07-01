package larkiot

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoLarkIotOrderGetgoodslistAPIResponse
iot渠道获取卖品信息 API返回值
taobao.lark.iot.order.getgoodslist

iot无人超市服务商通过接口获取影院的可售卖品数据 */
type TaobaoLarkIotOrderGetgoodslistAPIResponse struct {
	model.CommonResponse
	TaobaoLarkIotOrderGetgoodslistAPIResponseModel
}

// TaobaoLarkIotOrderGetgoodslistAPIResponseModel is iot渠道获取卖品信息 成功返回结果
type TaobaoLarkIotOrderGetgoodslistAPIResponseModel struct {
	XMLName xml.Name `xml:"lark_iot_order_getgoodslist_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 卖品信息列表
	Data *BizSingleResult `json:"data,omitempty" xml:"data,omitempty"`
}
