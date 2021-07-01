package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallServicecenterWorkcardEvaluateAPIRequest
服务商反馈鉴定结果 API请求
tmall.servicecenter.workcard.evaluate

服务商反馈鉴定结果 */
type TmallServicecenterWorkcardEvaluateAPIRequest struct {
	model.Params
	// 是否鉴定通过
	_passEvaluation bool
	// 鉴定不通过时的原因编码
	_failCode int64
	// 鉴定结果图片列表
	_picUrlList []string
	// 工单id
	_workcardId int64
}

// New
