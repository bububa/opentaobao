package omniorder

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/omniorder"
)

// Taobaoomnidealerodersget 获取单笔全渠道经销商订单的详细信息
// taobao.omni.dealer.oders.get
//
// 全渠道经销商获取单笔订单的详细信息
func Taobaoomnidealerodersget(clt *core.SDKClient, req *omniorder.TaobaoomnidealerodersgetAPIRequest, session string) (*omniorder.TaobaoomnidealerodersgetAPIResponse, error) {
	var resp omniorder.TaobaoomnidealerodersgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
