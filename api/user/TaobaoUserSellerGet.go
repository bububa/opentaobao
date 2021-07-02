package user

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/user"
)

// TaobaoUserSellerGet 查询卖家用户信息
// taobao.user.seller.get
//
// 查询卖家用户信息（只能查询有店铺的用户） 只能卖家类应用调用。
func TaobaoUserSellerGet(clt *core.SDKClient, req *user.TaobaoUserSellerGetAPIRequest, session string) (*user.TaobaoUserSellerGetAPIResponse, error) {
	var resp user.TaobaoUserSellerGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
