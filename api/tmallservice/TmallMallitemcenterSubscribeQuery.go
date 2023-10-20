package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// Tmallmallitemcentersubscribequery 天猫服务订购信息查询接口
// tmall.mallitemcenter.subscribe.query
//
// 查询商家服务订购信息
func Tmallmallitemcentersubscribequery(clt *core.SDKClient, req *tmallservice.TmallmallitemcentersubscribequeryAPIRequest, session string) (*tmallservice.TmallmallitemcentersubscribequeryAPIResponse, error) {
	var resp tmallservice.TmallmallitemcentersubscribequeryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
