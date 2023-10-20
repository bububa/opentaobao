package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// TmallServicecenterWorkcardCallRecord 回访记录回传API
// tmall.servicecenter.workcard.call.record
//
// 客满回访信息登记
func TmallServicecenterWorkcardCallRecord(clt *core.SDKClient, req *tmallservice.TmallServicecenterWorkcardCallRecordAPIRequest, resp *tmallservice.TmallServicecenterWorkcardCallRecordAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
