package jstinteractive

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoJstInteractivePointDecreaseAPIRequest
互动积分扣减接口 API请求
taobao.jst.interactive.point.decrease

扣减用户的互动积分 */
type TaobaoJstInteractivePointDecreaseAPIRequest struct {
	model.Params
	// 扣减的积分值
	_amount int64
	// 幂等操作码，要确保唯一性，同一个操作码只能使用一次，避免重复操作
	_operateCode string
}

// New
