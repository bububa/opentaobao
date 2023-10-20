package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// Tmallservicecenterworkcardcallrecord 回访记录回传API
// tmall.servicecenter.workcard.call.record
//
// 客满回访信息登记
func Tmallservicecenterworkcardcallrecord(clt *core.SDKClient, req *tmallservice.TmallservicecenterworkcardcallrecordAPIRequest, session string) (*tmallservice.TmallservicecenterworkcardcallrecordAPIResponse, error) {
	var resp tmallservice.TmallservicecenterworkcardcallrecordAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
