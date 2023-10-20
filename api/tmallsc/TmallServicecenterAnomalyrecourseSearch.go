package tmallsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallsc"
)

// Tmallservicecenteranomalyrecoursesearch 天猫服务平台服务商一键求助单查询
// tmall.servicecenter.anomalyrecourse.search
//
// 天猫服务平台服务商一键求助单查询
func Tmallservicecenteranomalyrecoursesearch(clt *core.SDKClient, req *tmallsc.TmallservicecenteranomalyrecoursesearchAPIRequest, session string) (*tmallsc.TmallservicecenteranomalyrecoursesearchAPIResponse, error) {
	var resp tmallsc.TmallservicecenteranomalyrecoursesearchAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
