package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// Tmallservicecenteranomalyrecoursehomedecorationquerybyid 天猫服务平台服务商查询商家投诉单
// tmall.servicecenter.anomalyrecourse.homedecoration.querybyid
//
// 天猫服务平台服务商查询商家投诉单
func Tmallservicecenteranomalyrecoursehomedecorationquerybyid(clt *core.SDKClient, req *tmallservice.TmallservicecenteranomalyrecoursehomedecorationquerybyidAPIRequest, session string) (*tmallservice.TmallservicecenteranomalyrecoursehomedecorationquerybyidAPIResponse, error) {
	var resp tmallservice.TmallservicecenteranomalyrecoursehomedecorationquerybyidAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
