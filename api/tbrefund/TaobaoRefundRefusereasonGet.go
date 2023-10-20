package tbrefund

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbrefund"
)

// Taobaorefundrefusereasonget 获取拒绝原因列表
// taobao.refund.refusereason.get
//
// 获取商家拒绝原因列表
func Taobaorefundrefusereasonget(clt *core.SDKClient, req *tbrefund.TaobaorefundrefusereasongetAPIRequest, session string) (*tbrefund.TaobaorefundrefusereasongetAPIResponse, error) {
	var resp tbrefund.TaobaorefundrefusereasongetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
