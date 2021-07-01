package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoPromotionCouponSellerSearchAPIRequest
查询绑定卖家优惠券相关信息 API请求
taobao.promotion.coupon.seller.search

查询绑定卖家相关优惠券信息  如isv  百川 等外部业务方 */
type TaobaoPromotionCouponSellerSearchAPIRequest struct {
	model.Params
	// 卖家昵称
	_sellerNick string
	// 当前第几页  从第一页开始
	_currentPage int64
	// 每页数据 最大20左右
	_pageSize int64
	// 券id集合
	_spreadIds []string
}

// New
