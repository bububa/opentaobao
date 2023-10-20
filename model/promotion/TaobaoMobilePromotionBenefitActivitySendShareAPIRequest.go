package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaomobilepromotionbenefitactivitysendshareAPIRequest 手淘专用单用户发放接口 API请求
// taobao.mobile.promotion.benefit.activity.send.share
//
// 卖家活动中需要通过该API来发放对应的权益。手淘专用、验证分享链路。
type TaobaomobilepromotionbenefitactivitysendshareAPIRequest struct {
	model.Params
	// 权益关联的活动ID
	_bizId string
	// 和bizId一起使用，标记分享链路的key。
	_shareKey string
	// 分享链路上的用户及用户是否发奖，u1_true,u2_true
	_shareUsers string
	// 调试线索
	_traceId string
	// 事务id
	_uniqueId string
	// 权益类型    其中ALIPAY_COUPON 对应的type值是1
	_benefitType int64
	// 活动详情id
	_detailId int64
	// 广播ID
	_feedId int64
	// 关联活动id
	_relationId int64
	// 权益发放数量
	_sendCount int64
}

// NewTaobaomobilepromotionbenefitactivitysendshareRequest 初始化TaobaomobilepromotionbenefitactivitysendshareAPIRequest对象
func NewTaobaomobilepromotionbenefitactivitysendshareRequest() *TaobaomobilepromotionbenefitactivitysendshareAPIRequest {
	return &TaobaomobilepromotionbenefitactivitysendshareAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaomobilepromotionbenefitactivitysendshareAPIRequest) GetApiMethodName() string {
	return "taobao.mobile.promotion.benefit.activity.send.share"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaomobilepromotionbenefitactivitysendshareAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaomobilepromotionbenefitactivitysendshareAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBizId is BizId Setter
// 权益关联的活动ID
func (r *TaobaomobilepromotionbenefitactivitysendshareAPIRequest) SetBizId(_bizId string) error {
	r._bizId = _bizId
	r.Set("biz_id", _bizId)
	return nil
}

// GetBizId BizId Getter
func (r TaobaomobilepromotionbenefitactivitysendshareAPIRequest) GetBizId() string {
	return r._bizId
}

// SetShareKey is ShareKey Setter
// 和bizId一起使用，标记分享链路的key。
func (r *TaobaomobilepromotionbenefitactivitysendshareAPIRequest) SetShareKey(_shareKey string) error {
	r._shareKey = _shareKey
	r.Set("share_key", _shareKey)
	return nil
}

// GetShareKey ShareKey Getter
func (r TaobaomobilepromotionbenefitactivitysendshareAPIRequest) GetShareKey() string {
	return r._shareKey
}

// SetShareUsers is ShareUsers Setter
// 分享链路上的用户及用户是否发奖，u1_true,u2_true
func (r *TaobaomobilepromotionbenefitactivitysendshareAPIRequest) SetShareUsers(_shareUsers string) error {
	r._shareUsers = _shareUsers
	r.Set("share_users", _shareUsers)
	return nil
}

// GetShareUsers ShareUsers Getter
func (r TaobaomobilepromotionbenefitactivitysendshareAPIRequest) GetShareUsers() string {
	return r._shareUsers
}

// SetTraceId is TraceId Setter
// 调试线索
func (r *TaobaomobilepromotionbenefitactivitysendshareAPIRequest) SetTraceId(_traceId string) error {
	r._traceId = _traceId
	r.Set("trace_id", _traceId)
	return nil
}

// GetTraceId TraceId Getter
func (r TaobaomobilepromotionbenefitactivitysendshareAPIRequest) GetTraceId() string {
	return r._traceId
}

// SetUniqueId is UniqueId Setter
// 事务id
func (r *TaobaomobilepromotionbenefitactivitysendshareAPIRequest) SetUniqueId(_uniqueId string) error {
	r._uniqueId = _uniqueId
	r.Set("unique_id", _uniqueId)
	return nil
}

// GetUniqueId UniqueId Getter
func (r TaobaomobilepromotionbenefitactivitysendshareAPIRequest) GetUniqueId() string {
	return r._uniqueId
}

// SetBenefitType is BenefitType Setter
// 权益类型    其中ALIPAY_COUPON 对应的type值是1
func (r *TaobaomobilepromotionbenefitactivitysendshareAPIRequest) SetBenefitType(_benefitType int64) error {
	r._benefitType = _benefitType
	r.Set("benefit_type", _benefitType)
	return nil
}

// GetBenefitType BenefitType Getter
func (r TaobaomobilepromotionbenefitactivitysendshareAPIRequest) GetBenefitType() int64 {
	return r._benefitType
}

// SetDetailId is DetailId Setter
// 活动详情id
func (r *TaobaomobilepromotionbenefitactivitysendshareAPIRequest) SetDetailId(_detailId int64) error {
	r._detailId = _detailId
	r.Set("detail_id", _detailId)
	return nil
}

// GetDetailId DetailId Getter
func (r TaobaomobilepromotionbenefitactivitysendshareAPIRequest) GetDetailId() int64 {
	return r._detailId
}

// SetFeedId is FeedId Setter
// 广播ID
func (r *TaobaomobilepromotionbenefitactivitysendshareAPIRequest) SetFeedId(_feedId int64) error {
	r._feedId = _feedId
	r.Set("feed_id", _feedId)
	return nil
}

// GetFeedId FeedId Getter
func (r TaobaomobilepromotionbenefitactivitysendshareAPIRequest) GetFeedId() int64 {
	return r._feedId
}

// SetRelationId is RelationId Setter
// 关联活动id
func (r *TaobaomobilepromotionbenefitactivitysendshareAPIRequest) SetRelationId(_relationId int64) error {
	r._relationId = _relationId
	r.Set("relation_id", _relationId)
	return nil
}

// GetRelationId RelationId Getter
func (r TaobaomobilepromotionbenefitactivitysendshareAPIRequest) GetRelationId() int64 {
	return r._relationId
}

// SetSendCount is SendCount Setter
// 权益发放数量
func (r *TaobaomobilepromotionbenefitactivitysendshareAPIRequest) SetSendCount(_sendCount int64) error {
	r._sendCount = _sendCount
	r.Set("send_count", _sendCount)
	return nil
}

// GetSendCount SendCount Getter
func (r TaobaomobilepromotionbenefitactivitysendshareAPIRequest) GetSendCount() int64 {
	return r._sendCount
}
