package tmallsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallsc"
)

// Tmallservicecenteranomalyrecoursehomedecorationappeal 天猫服务平台商家投诉单服务商申诉接口
// tmall.servicecenter.anomalyrecourse.homedecoration.appeal
//
// 天猫服务平台商家投诉单服务商申诉接口
func Tmallservicecenteranomalyrecoursehomedecorationappeal(clt *core.SDKClient, req *tmallsc.TmallservicecenteranomalyrecoursehomedecorationappealAPIRequest, session string) (*tmallsc.TmallservicecenteranomalyrecoursehomedecorationappealAPIResponse, error) {
	var resp tmallsc.TmallservicecenteranomalyrecoursehomedecorationappealAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
