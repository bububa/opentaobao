package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// Tmallservicecenterworkcardquery 工单查询接口
// tmall.servicecenter.workcard.query
//
// 工单查询接口
func Tmallservicecenterworkcardquery(clt *core.SDKClient, req *tmallservice.TmallservicecenterworkcardqueryAPIRequest, session string) (*tmallservice.TmallservicecenterworkcardqueryAPIResponse, error) {
	var resp tmallservice.TmallservicecenterworkcardqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
