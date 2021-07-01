package tmallsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallServicecenterAnomalyrecourseRemarkUpdateAPIRequest
天猫服务平台一键求助单服务商备注更新接口 API请求
tmall.servicecenter.anomalyrecourse.remark.update

一键求助服务商可以回传备注 */
type TmallServicecenterAnomalyrecourseRemarkUpdateAPIRequest struct {
	model.Params
	// 需要更新的一键求助单id
	_id int64
	// 需要更新的服务商备注
	_remark string
}

// New
