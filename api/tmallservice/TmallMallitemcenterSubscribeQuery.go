package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// TmallMallitemcenterSubscribeQuery 天猫服务订购信息查询接口
// tmall.mallitemcenter.subscribe.query
//
// 查询商家服务订购信息
func TmallMallitemcenterSubscribeQuery(clt *core.SDKClient, req *tmallservice.TmallMallitemcenterSubscribeQueryAPIRequest, session string) (*tmallservice.TmallMallitemcenterSubscribeQueryAPIResponse, error) {
	var resp tmallservice.TmallMallitemcenterSubscribeQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
