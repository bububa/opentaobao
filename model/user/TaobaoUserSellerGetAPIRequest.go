package user

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoUserSellerGetAPIRequest
查询卖家用户信息 API请求
taobao.user.seller.get

查询卖家用户信息（只能查询有店铺的用户） 只能卖家类应用调用。 */
type TaobaoUserSellerGetAPIRequest struct {
	model.Params
	// 需要返回的字段列表，可选值为返回示例值中的可以看到的字段
	_fields []string
}

// New
