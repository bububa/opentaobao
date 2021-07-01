package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoMobilePromotionCouponSellerSearchAPIRequest
查询绑定卖家优惠券相关信息(手淘专用) API请求
taobao.mobile.promotion.coupon.seller.search

查询绑定卖家相关优惠券信息 如isv 百川 等外部业务方 */
type TaobaoMobilePromotionCouponSellerSearchAPIRequest struct {
	model.Params
	// 请求id 排查线索 需保证单次调用唯一
	_traceId string
	// 券id集合
	_spreadIds string
	// 每页数据 最大20左右
	_pageSize int64
	// 当前第几页 从第一页开始
	_currentPage int64
}

// New
