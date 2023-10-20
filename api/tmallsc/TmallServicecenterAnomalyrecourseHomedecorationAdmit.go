package tmallsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallsc"
)

// Tmallservicecenteranomalyrecoursehomedecorationadmit 天猫服务平台商家投诉单服务商认责接口
// tmall.servicecenter.anomalyrecourse.homedecoration.admit
//
// 天猫服务平台商家投诉单服务商认责接口
func Tmallservicecenteranomalyrecoursehomedecorationadmit(clt *core.SDKClient, req *tmallsc.TmallservicecenteranomalyrecoursehomedecorationadmitAPIRequest, session string) (*tmallsc.TmallservicecenteranomalyrecoursehomedecorationadmitAPIResponse, error) {
	var resp tmallsc.TmallservicecenteranomalyrecoursehomedecorationadmitAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
