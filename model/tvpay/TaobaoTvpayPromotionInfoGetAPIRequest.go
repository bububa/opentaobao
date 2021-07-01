package tvpay

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTvpayPromotionInfoGetAPIRequest
tv支付查询消费抽奖配置 API请求
taobao.tvpay.promotion.info.get

查询消费抽奖配置 */
type TaobaoTvpayPromotionInfoGetAPIRequest struct {
	model.Params
	// 设备id
	_deviceId string
	// 来源
	_from string
	// 商品id
	_subjectId string
	// 淘系订单号
	_extOrderId string
	// 是否淘系
	_isTao bool
	// 标题
	_subject string
}

// New
