package jstinteractive

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoJstInteractivePointIncreaseAPIRequest
互动积分发放接口 API请求
taobao.jst.interactive.point.increase

向用户发放互动积分 */
type TaobaoJstInteractivePointIncreaseAPIRequest struct {
	model.Params
	// 发放的积分值
	_amount int64
	// 幂等操作码，要确保唯一性，同一个操作码只能使用一次，避免重复操作
	_operateCode string
}

// New
