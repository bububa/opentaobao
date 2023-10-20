package drug

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drug"
)

// TaobaoAlihealthDrugUserShopGet 根据用户id获取店铺id
// taobao.alihealth.drug.user.shop.get
//
// 提供给千牛智能客服，获取用户当前咨询的店铺ID
func TaobaoAlihealthDrugUserShopGet(clt *core.SDKClient, req *drug.TaobaoAlihealthDrugUserShopGetAPIRequest, resp *drug.TaobaoAlihealthDrugUserShopGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
