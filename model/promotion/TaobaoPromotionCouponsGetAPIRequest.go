package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoPromotionCouponsGetAPIRequest
查询卖家优惠券 API请求
taobao.promotion.coupons.get

查询卖家已经创建的优惠券，接口返回信息：优惠券ID，面值，创建时间，有效期，使用条件，使用渠道，创建渠道，优惠券总数量 */
type TaobaoPromotionCouponsGetAPIRequest struct {
	model.Params
	// 优惠券的id，唯一标识这个优惠券
	_couponId int64
	// 优惠券的截止日期
	_endTime string
	// 优惠券的面额，必须是3，5，10，20，50,100
	_denominations int64
	// 查询的页号，结果集是分页返回的，每页20条
	_pageNo int64
	// 每页条数
	_pageSize int64
}

// New
