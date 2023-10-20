package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// TmallMallitemcenterSubscribeQuery 天猫服务订购信息查询接口
// tmall.mallitemcenter.subscribe.query
//
// 查询商家服务订购信息
func TmallMallitemcenterSubscribeQuery(clt *core.SDKClient, req *tmallservice.TmallMallitemcenterSubscribeQueryAPIRequest, resp *tmallservice.TmallMallitemcenterSubscribeQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
