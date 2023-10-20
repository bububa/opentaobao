package tbrefund

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbrefund"
)

// Taobaorefundget 获取单笔退款详情
// taobao.refund.get
//
// 获取单笔退款详情
func Taobaorefundget(clt *core.SDKClient, req *tbrefund.TaobaorefundgetAPIRequest, session string) (*tbrefund.TaobaorefundgetAPIResponse, error) {
	var resp tbrefund.TaobaorefundgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
