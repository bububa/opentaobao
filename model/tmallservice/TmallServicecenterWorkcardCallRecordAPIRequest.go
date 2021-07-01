package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallServicecenterWorkcardCallRecordAPIRequest
回访记录回传API API请求
tmall.servicecenter.workcard.call.record

客满回访信息登记 */
type TmallServicecenterWorkcardCallRecordAPIRequest struct {
	model.Params
	// 请求入参
	_busiRequest *UpdateAttributeRequest
}

// New
