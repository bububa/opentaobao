package tbuser

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbuser"
)

// TaobaoUserSellerGet 查询卖家用户信息
// taobao.user.seller.get
//
// 查询卖家用户信息（只能查询有店铺的用户） 只能卖家类应用调用。
func TaobaoUserSellerGet(clt *core.SDKClient, req *tbuser.TaobaoUserSellerGetAPIRequest, session string) (*tbuser.TaobaoUserSellerGetAPIResponse, error) {
	var resp tbuser.TaobaoUserSellerGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
