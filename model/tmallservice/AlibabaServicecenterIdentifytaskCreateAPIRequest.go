package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaservicecenteridentifytaskcreateAPIRequest 创建核销单 API请求
// alibaba.servicecenter.identifytask.create
//
// 创建核销单
type AlibabaservicecenteridentifytaskcreateAPIRequest struct {
	model.Params
	// 工单集合
	_workcardIds []int64
	// 核销单外部标识，服务商保证唯一。如果创建核销单时传入重复的outer_id，系统直接会返回服务商下该outer_id对应的核销单，不会重新创建新核销单。
	_outerId string
}

// NewAlibabaservicecenteridentifytaskcreateRequest 初始化AlibabaservicecenteridentifytaskcreateAPIRequest对象
func NewAlibabaservicecenteridentifytaskcreateRequest() *AlibabaservicecenteridentifytaskcreateAPIRequest {
	return &AlibabaservicecenteridentifytaskcreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaservicecenteridentifytaskcreateAPIRequest) GetApiMethodName() string {
	return "alibaba.servicecenter.identifytask.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaservicecenteridentifytaskcreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaservicecenteridentifytaskcreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetWorkcardIds is WorkcardIds Setter
// 工单集合
func (r *AlibabaservicecenteridentifytaskcreateAPIRequest) SetWorkcardIds(_workcardIds []int64) error {
	r._workcardIds = _workcardIds
	r.Set("workcard_ids", _workcardIds)
	return nil
}

// GetWorkcardIds WorkcardIds Getter
func (r AlibabaservicecenteridentifytaskcreateAPIRequest) GetWorkcardIds() []int64 {
	return r._workcardIds
}

// SetOuterId is OuterId Setter
// 核销单外部标识，服务商保证唯一。如果创建核销单时传入重复的outer_id，系统直接会返回服务商下该outer_id对应的核销单，不会重新创建新核销单。
func (r *AlibabaservicecenteridentifytaskcreateAPIRequest) SetOuterId(_outerId string) error {
	r._outerId = _outerId
	r.Set("outer_id", _outerId)
	return nil
}

// GetOuterId OuterId Getter
func (r AlibabaservicecenteridentifytaskcreateAPIRequest) GetOuterId() string {
	return r._outerId
}
