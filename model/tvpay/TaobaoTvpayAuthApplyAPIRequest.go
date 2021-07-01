package tvpay

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTvpayAuthApplyAPIRequest
tv支付申请设备授权 API请求
taobao.tvpay.auth.apply

为用户在指定设备上申请支付授权 */
type TaobaoTvpayAuthApplyAPIRequest struct {
	model.Params
	// 设备id
	_deviceId string
	// 请求来源
	_from string
	// 场景
	_bizScene string
	// 商品名称
	_itemName string
	// 授权类型
	_operateType string
	// 外部订单号
	_outApproveId string
	// 金额
	_totalFee string
}

// New
