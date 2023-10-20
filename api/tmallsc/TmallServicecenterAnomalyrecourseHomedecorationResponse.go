package tmallsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallsc"
)

// Tmallservicecenteranomalyrecoursehomedecorationresponse 天猫服务平台商家投诉单服务商响应接口
// tmall.servicecenter.anomalyrecourse.homedecoration.response
//
// 天猫服务平台商家投诉单服务商响应接口
func Tmallservicecenteranomalyrecoursehomedecorationresponse(clt *core.SDKClient, req *tmallsc.TmallservicecenteranomalyrecoursehomedecorationresponseAPIRequest, session string) (*tmallsc.TmallservicecenteranomalyrecoursehomedecorationresponseAPIResponse, error) {
	var resp tmallsc.TmallservicecenteranomalyrecoursehomedecorationresponseAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
