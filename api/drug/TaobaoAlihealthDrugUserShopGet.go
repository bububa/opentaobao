package drug

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drug"
)

// Taobaoalihealthdrugusershopget 根据用户id获取店铺id
// taobao.alihealth.drug.user.shop.get
//
// 提供给千牛智能客服，获取用户当前咨询的店铺ID
func Taobaoalihealthdrugusershopget(clt *core.SDKClient, req *drug.TaobaoalihealthdrugusershopgetAPIRequest, session string) (*drug.TaobaoalihealthdrugusershopgetAPIResponse, error) {
	var resp drug.TaobaoalihealthdrugusershopgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
