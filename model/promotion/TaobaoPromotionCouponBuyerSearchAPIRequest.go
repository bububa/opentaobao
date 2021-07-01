package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoPromotionCouponBuyerSearchAPIRequest
查询买家在相关app领取的优惠券信息 API请求
taobao.promotion.coupon.buyer.search

查询买家在相关app领取的优惠券信息 */
type TaobaoPromotionCouponBuyerSearchAPIRequest struct {
	model.Params
	// 卖家昵称
	_sellerNick string
	// 券状态  "正常",1 "已删除",-1 "已使用",-2 "冻结",0
	_status int64
	// 每页数据 建议20左右
	_pageSize int64
	// 当前第几页  从第一页开始
	_currentPage int64
	// 结束时间
	_endTime string
}

// New
