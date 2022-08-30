package tmallcar

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallAliautoOrderQrcodeAPIResponse 根据商品id列表获取可扫描下单二维码 API返回值
// tmall.aliauto.order.qrcode
//
// 根据商品id列表获取可扫描下单二维码
type TmallAliautoOrderQrcodeAPIResponse struct {
	model.CommonResponse
	TmallAliautoOrderQrcodeAPIResponseModel
}

// TmallAliautoOrderQrcodeAPIResponseModel is 根据商品id列表获取可扫描下单二维码 成功返回结果
type TmallAliautoOrderQrcodeAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_aliauto_order_qrcode_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回的数据实体
	Data *ConfirmOrderQrCode4IsvDto `json:"data,omitempty" xml:"data,omitempty"`
}
