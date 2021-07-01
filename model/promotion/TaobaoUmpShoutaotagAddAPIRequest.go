package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoUmpShoutaotagAddAPIRequest
手淘定向优惠打标接口 API请求
taobao.ump.shoutaotag.add

手淘定向优惠的优惠标签打标接口
给特定的手淘买家打上优惠标记，标记承载在自己的业务标签库中，标签有效期为7天。 */
type TaobaoUmpShoutaotagAddAPIRequest struct {
	model.Params
	// 商品ID
	_itemId int64
	// 买家ID
	_buyerId int64
	// 渠道KEY
	_channelKey string
}

// New
