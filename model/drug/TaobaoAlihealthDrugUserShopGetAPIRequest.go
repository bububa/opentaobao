package drug

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAlihealthDrugUserShopGetAPIRequest
根据用户id获取店铺id API请求
taobao.alihealth.drug.user.shop.get

提供给千牛智能客服，获取用户当前咨询的店铺ID */
type TaobaoAlihealthDrugUserShopGetAPIRequest struct {
	model.Params
	// 用户昵称
	_userNick string
}

// New
