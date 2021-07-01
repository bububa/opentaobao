package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripBtripCostCenterGetAPIRequest
获取用户费用归属 API请求
alitrip.btrip.cost.center.get

获取差旅申请用户的费用归属列表 */
type AlitripBtripCostCenterGetAPIRequest struct {
	model.Params
	// 企业id
	_corpId string
	// 用户id
	_userId string
}

// New
