package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaServicecenterIdentifytaskCreateAPIRequest
创建核销单 API请求
alibaba.servicecenter.identifytask.create

创建核销单 */
type AlibabaServicecenterIdentifytaskCreateAPIRequest struct {
	model.Params
	// 工单集合
	_workcardIds []int64
	// 核销单外部标识，服务商保证唯一。如果创建核销单时传入重复的outer_id，系统直接会返回服务商下该outer_id对应的核销单，不会重新创建新核销单。
	_outerId string
}

// New
