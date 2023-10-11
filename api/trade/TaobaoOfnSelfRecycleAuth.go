package trade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/trade"
)

// TaobaoOfnSelfRecycleAuth 自助回收鉴权
// taobao.ofn.self.recycle.auth
//
// 自助回收鉴权
func TaobaoOfnSelfRecycleAuth(clt *core.SDKClient, req *trade.TaobaoOfnSelfRecycleAuthAPIRequest, session string) (*trade.TaobaoOfnSelfRecycleAuthAPIResponse, error) {
	var resp trade.TaobaoOfnSelfRecycleAuthAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
