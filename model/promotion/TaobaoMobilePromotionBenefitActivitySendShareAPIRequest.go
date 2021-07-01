package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoMobilePromotionBenefitActivitySendShareAPIRequest
手淘专用单用户发放接口 API请求
taobao.mobile.promotion.benefit.activity.send.share

卖家活动中需要通过该API来发放对应的权益。手淘专用、验证分享链路。 */
type TaobaoMobilePromotionBenefitActivitySendShareAPIRequest struct {
	model.Params
	// 权益类型    其中ALIPAY_COUPON 对应的type值是1
	_benefitType int64
	// 权益关联的活动ID
	_bizId string
	// 活动详情id
	_detailId int64
	// 广播ID
	_feedId int64
	// 关联活动id
	_relationId int64
	// 权益发放数量
	_sendCount int64
	// 和bizId一起使用，标记分享链路的key。
	_shareKey string
	// 分享链路上的用户及用户是否发奖，u1_true,u2_true
	_shareUsers string
	// 调试线索
	_traceId string
	// 事务id
	_uniqueId string
}

// New
