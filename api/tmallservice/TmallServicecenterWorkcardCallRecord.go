package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

/* TmallServicecenterWorkcardCallRecord
回访记录回传API
tmall.servicecenter.workcard.call.record

客满回访信息登记 */
func TmallServicecenterWorkcardCallRecord(clt *core.SDKClient, req *tmallservice.TmallServicecenterWorkcardCallRecordAPIRequest, session string) (*tmallservice.TmallServicecenterWorkcardCallRecordAPIResponse, error) {
	var resp tmallservice.TmallServicecenterWorkcardCallRecordAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
