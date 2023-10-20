package tbk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbk"
)

// Taobaotbkdgnewuserordersum 淘宝客-推广者-拉新活动对应数据查询
// taobao.tbk.dg.newuser.order.sum
//
// 拉新活动汇总API
func Taobaotbkdgnewuserordersum(clt *core.SDKClient, req *tbk.TaobaotbkdgnewuserordersumAPIRequest, session string) (*tbk.TaobaotbkdgnewuserordersumAPIResponse, error) {
	var resp tbk.TaobaotbkdgnewuserordersumAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
