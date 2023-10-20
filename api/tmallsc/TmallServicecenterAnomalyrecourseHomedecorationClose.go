package tmallsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallsc"
)

// Tmallservicecenteranomalyrecoursehomedecorationclose 天猫服务平台商家投诉单服务商完结接口
// tmall.servicecenter.anomalyrecourse.homedecoration.close
//
// 天猫服务平台商家投诉单服务商完结接口
func Tmallservicecenteranomalyrecoursehomedecorationclose(clt *core.SDKClient, req *tmallsc.TmallservicecenteranomalyrecoursehomedecorationcloseAPIRequest, session string) (*tmallsc.TmallservicecenteranomalyrecoursehomedecorationcloseAPIResponse, error) {
	var resp tmallsc.TmallservicecenteranomalyrecoursehomedecorationcloseAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
