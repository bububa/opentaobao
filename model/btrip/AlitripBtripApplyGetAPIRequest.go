package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripBtripApplyGetAPIRequest
获取单个审批单 API请求
alitrip.btrip.apply.get

获取单个审批单的详情数据 */
type AlitripBtripApplyGetAPIRequest struct {
	model.Params
	// 外部审批单id
	_thirdpartApplyId string
	// 阿里商旅审批单id
	_applyId int64
	// 企业id
	_corpId string
	// 审批单展示id
	_applyShowId string
}

// New
