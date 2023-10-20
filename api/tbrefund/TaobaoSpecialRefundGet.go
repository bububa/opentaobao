package tbrefund

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbrefund"
)

// Taobaospecialrefundget 特殊部分退纠纷单查询
// taobao.special.refund.get
//
// 获取单笔特殊部分退的纠纷单查询
func Taobaospecialrefundget(clt *core.SDKClient, req *tbrefund.TaobaospecialrefundgetAPIRequest, session string) (*tbrefund.TaobaospecialrefundgetAPIResponse, error) {
	var resp tbrefund.TaobaospecialrefundgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
