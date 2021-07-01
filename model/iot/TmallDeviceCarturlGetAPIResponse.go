package iot

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TmallDeviceCarturlGetAPIResponse
添加商品到购物车 API返回值
tmall.device.carturl.get

获取二维码，支持添加商品到购物车 */
type TmallDeviceCarturlGetAPIResponse struct {
	model.CommonResponse
	TmallDeviceCarturlGetAPIResponseModel
}

// TmallDeviceCarturlGetAPIResponseModel is 添加商品到购物车 成功返回结果
type TmallDeviceCarturlGetAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_device_carturl_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 二维码短链接地址
	ShortUrl string `json:"short_url,omitempty" xml:"short_url,omitempty"`
	// 二维码图片URL
	ShortImgUrl string `json:"short_img_url,omitempty" xml:"short_img_url,omitempty"`
}
