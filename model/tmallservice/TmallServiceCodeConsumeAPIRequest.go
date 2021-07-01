package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallServiceCodeConsumeAPIRequest
天猫服务平台服务核销 API请求
tmall.service.code.consume

天猫服务平台－服务核销 */
type TmallServiceCodeConsumeAPIRequest struct {
	model.Params
	// 核销码
	_consumeCode string
	// 核销帐号
	_operatorNick string
	// 门店id
	_shopId string
}

// New
