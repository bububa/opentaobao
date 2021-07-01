package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoPromotionmiscCommonItemDetailUpdateAPIRequest
修改通用单品优惠详情 API请求
taobao.promotionmisc.common.item.detail.update

修改通用单品优惠详情。
1、该接口只修改活动下参与的商品的优惠信息，如需要增加、删除活动，请调用taobao.promotionmisc.common.item.activity.add和taobao.promotionmisc.common.item.activity.delete接口；
2、使用该接口时需要把未做修改的字段值也传入；
3、此接口受卖家最低折扣限制，如果优惠力度大于卖家设置的最低折扣则不能修改 */
type TaobaoPromotionmiscCommonItemDetailUpdateAPIRequest struct {
	model.Params
	// 优惠活动ID
	_activityId int64
	// 优惠详情ID
	_detailId int64
	// 商品ID
	_itemId int64
	// 优惠类型，只有两种可选值：0-减钱；1-打折
	_promotionType int64
	// 优惠力度，其值的解释方式由promotion_type定义：当为减钱时解释成减钱数量，如：900表示减9元；当为打折时解释成打折折扣，如：900表示打9折
	_promotionValue int64
}

// New
