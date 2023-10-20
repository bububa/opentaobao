package tmallnr

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallnr"
)

// Tmallnrfulfillorderquery 零售商获取品牌商的单笔订单
// tmall.nr.fulfill.order.query
//
// 零售商获取品牌商的单笔订单，后端服务有零售商和品牌商的绑定关系，存在开关控制；返回值存在品牌方用户的电话号码，当前电话号码是屏蔽中间四位
func Tmallnrfulfillorderquery(clt *core.SDKClient, req *tmallnr.TmallnrfulfillorderqueryAPIRequest, session string) (*tmallnr.TmallnrfulfillorderqueryAPIResponse, error) {
	var resp tmallnr.TmallnrfulfillorderqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
