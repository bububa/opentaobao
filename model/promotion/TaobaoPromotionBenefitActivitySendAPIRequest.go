package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoPromotionBenefitActivitySendAPIRequest
活动权益发放接口 API请求
taobao.promotion.benefit.activity.send

活动权益发放接口，用于卖家针对活动进行权益发放 */
type TaobaoPromotionBenefitActivitySendAPIRequest struct {
	model.Params
	// 单个权益发放请求
	_sendRequest *BenefitSingleSendRequest
	// 非混淆的接收者id
	_receiverId int64
	// 混淆的接收者nick
	_nick string
	// 非混淆的接收者nick
	_platNick string
	// 混淆的接收者id
	_mixReceiverId string
}

// New
