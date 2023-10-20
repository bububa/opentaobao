package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// Tmallservicecenterworkcarddesensitizationquery 工单查询接口
// tmall.servicecenter.workcard.desensitization.query
//
// 工单查询接口
func Tmallservicecenterworkcarddesensitizationquery(clt *core.SDKClient, req *tmallservice.TmallservicecenterworkcarddesensitizationqueryAPIRequest, session string) (*tmallservice.TmallservicecenterworkcarddesensitizationqueryAPIResponse, error) {
	var resp tmallservice.TmallservicecenterworkcarddesensitizationqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
