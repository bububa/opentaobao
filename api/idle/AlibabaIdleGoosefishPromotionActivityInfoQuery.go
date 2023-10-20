package idle

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idle"
)

// Alibabaidlegoosefishpromotionactivityinfoquery 闲鱼三方活动参与信息查询
// alibaba.idle.goosefish.promotion.activity.info.query
//
// 闲鱼三方活动参与信息查询
func Alibabaidlegoosefishpromotionactivityinfoquery(clt *core.SDKClient, req *idle.AlibabaidlegoosefishpromotionactivityinfoqueryAPIRequest, session string) (*idle.AlibabaidlegoosefishpromotionactivityinfoqueryAPIResponse, error) {
	var resp idle.AlibabaidlegoosefishpromotionactivityinfoqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
