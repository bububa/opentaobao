package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoMobilePromotionBenefitActivitySendShareAPIRequest 手淘专用单用户发放接口 API请求
// taobao.mobile.promotion.benefit.activity.send.share
//
// 卖家活动中需要通过该API来发放对应的权益。手淘专用、验证分享链路。
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

// NewTaobaoMobilePromotionBenefitActivitySendShareRequest 初始化TaobaoMobilePromotionBenefitActivitySendShareAPIRequest对象
func NewTaobaoMobilePromotionBenefitActivitySendShareRequest() *TaobaoMobilePromotionBenefitActivitySendShareAPIRequest {
	return &TaobaoMobilePromotionBenefitActivitySendShareAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoMobilePromotionBenefitActivitySendShareAPIRequest) GetApiMethodName() string {
	return "taobao.mobile.promotion.benefit.activity.send.share"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoMobilePromotionBenefitActivitySendShareAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetBenefitType is BenefitType Setter
// 权益类型    其中ALIPAY_COUPON 对应的type值是1
func (r *TaobaoMobilePromotionBenefitActivitySendShareAPIRequest) SetBenefitType(_benefitType int64) error {
	r._benefitType = _benefitType
	r.Set("benefit_type", _benefitType)
	return nil
}

// GetBenefitType BenefitType Getter
func (r TaobaoMobilePromotionBenefitActivitySendShareAPIRequest) GetBenefitType() int64 {
	return r._benefitType
}

// SetBizId is BizId Setter
// 权益关联的活动ID
func (r *TaobaoMobilePromotionBenefitActivitySendShareAPIRequest) SetBizId(_bizId string) error {
	r._bizId = _bizId
	r.Set("biz_id", _bizId)
	return nil
}

// GetBizId BizId Getter
func (r TaobaoMobilePromotionBenefitActivitySendShareAPIRequest) GetBizId() string {
	return r._bizId
}

// SetDetailId is DetailId Setter
// 活动详情id
func (r *TaobaoMobilePromotionBenefitActivitySendShareAPIRequest) SetDetailId(_detailId int64) error {
	r._detailId = _detailId
	r.Set("detail_id", _detailId)
	return nil
}

// GetDetailId DetailId Getter
func (r TaobaoMobilePromotionBenefitActivitySendShareAPIRequest) GetDetailId() int64 {
	return r._detailId
}

// SetFeedId is FeedId Setter
// 广播ID
func (r *TaobaoMobilePromotionBenefitActivitySendShareAPIRequest) SetFeedId(_feedId int64) error {
	r._feedId = _feedId
	r.Set("feed_id", _feedId)
	return nil
}

// GetFeedId FeedId Getter
func (r TaobaoMobilePromotionBenefitActivitySendShareAPIRequest) GetFeedId() int64 {
	return r._feedId
}

// SetRelationId is RelationId Setter
// 关联活动id
func (r *TaobaoMobilePromotionBenefitActivitySendShareAPIRequest) SetRelationId(_relationId int64) error {
	r._relationId = _relationId
	r.Set("relation_id", _relationId)
	return nil
}

// GetRelationId RelationId Getter
func (r TaobaoMobilePromotionBenefitActivitySendShareAPIRequest) GetRelationId() int64 {
	return r._relationId
}

// SetSendCount is SendCount Setter
// 权益发放数量
func (r *TaobaoMobilePromotionBenefitActivitySendShareAPIRequest) SetSendCount(_sendCount int64) error {
	r._sendCount = _sendCount
	r.Set("send_count", _sendCount)
	return nil
}

// GetSendCount SendCount Getter
func (r TaobaoMobilePromotionBenefitActivitySendShareAPIRequest) GetSendCount() int64 {
	return r._sendCount
}

// SetShareKey is ShareKey Setter
// 和bizId一起使用，标记分享链路的key。
func (r *TaobaoMobilePromotionBenefitActivitySendShareAPIRequest) SetShareKey(_shareKey string) error {
	r._shareKey = _shareKey
	r.Set("share_key", _shareKey)
	return nil
}

// GetShareKey ShareKey Getter
func (r TaobaoMobilePromotionBenefitActivitySendShareAPIRequest) GetShareKey() string {
	return r._shareKey
}

// SetShareUsers is ShareUsers Setter
// 分享链路上的用户及用户是否发奖，u1_true,u2_true
func (r *TaobaoMobilePromotionBenefitActivitySendShareAPIRequest) SetShareUsers(_shareUsers string) error {
	r._shareUsers = _shareUsers
	r.Set("share_users", _shareUsers)
	return nil
}

// GetShareUsers ShareUsers Getter
func (r TaobaoMobilePromotionBenefitActivitySendShareAPIRequest) GetShareUsers() string {
	return r._shareUsers
}

// SetTraceId is TraceId Setter
// 调试线索
func (r *TaobaoMobilePromotionBenefitActivitySendShareAPIRequest) SetTraceId(_traceId string) error {
	r._traceId = _traceId
	r.Set("trace_id", _traceId)
	return nil
}

// GetTraceId TraceId Getter
func (r TaobaoMobilePromotionBenefitActivitySendShareAPIRequest) GetTraceId() string {
	return r._traceId
}

// SetUniqueId is UniqueId Setter
// 事务id
func (r *TaobaoMobilePromotionBenefitActivitySendShareAPIRequest) SetUniqueId(_uniqueId string) error {
	r._uniqueId = _uniqueId
	r.Set("unique_id", _uniqueId)
	return nil
}

// GetUniqueId UniqueId Getter
func (r TaobaoMobilePromotionBenefitActivitySendShareAPIRequest) GetUniqueId() string {
	return r._uniqueId
}
