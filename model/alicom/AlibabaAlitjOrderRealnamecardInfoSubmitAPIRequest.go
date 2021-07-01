package alicom

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlitjOrderRealnamecardInfoSubmitAPIRequest
阿里实人认证卡片信息回传 API请求
alibaba.alitj.order.realnamecard.info.submit

阿里实人认证卡片信息回传。ISP相关商家在线对接阿里通信的实人认证功能，在线提交订单对应运营商的合约订购相关信息，以便完成在线使用实人认证功能。 */
type AlibabaAlitjOrderRealnamecardInfoSubmitAPIRequest struct {
	model.Params
	// 淘宝订单号
	_orderNo int64
	// sim卡iccid（一般为18位到20位）
	_iccid string
}

// New
