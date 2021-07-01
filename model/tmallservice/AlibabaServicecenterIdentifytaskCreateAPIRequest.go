package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaServicecenterIdentifytaskCreateAPIRequest
创建核销单 API请求
alibaba.servicecenter.identifytask.create

创建核销单 */
type AlibabaServicecenterIdentifytaskCreateAPIRequest struct {
	model.Params
	// 工单集合
	_workcardIds []int64
	// 核销单外部标识，服务商保证唯一。如果创建核销单时传入重复的outer_id，系统直接会返回服务商下该outer_id对应的核销单，不会重新创建新核销单。
	_outerId string
}

// NewAlibabaServicecenterIdentifytaskCreateRequest 初始化AlibabaServicecenterIdentifytaskCreateAPIRequest对象
func NewAlibabaServicecenterIdentifytaskCreateRequest() *AlibabaServicecenterIdentifytaskCreateAPIRequest {
	return &AlibabaServicecenterIdentifytaskCreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaServicecenterIdentifytaskCreateAPIRequest) GetApiMethodName() string {
	return "alibaba.servicecenter.identifytask.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaServicecenterIdentifytaskCreateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is WorkcardIds Setter
// 工单集合
func (r *AlibabaServicecenterIdentifytaskCreateAPIRequest) SetWorkcardIds(_workcardIds []int64) error {
	r._workcardIds = _workcardIds
	r.Set("workcard_ids", _workcardIds)
	return nil
}

// Get WorkcardIds Getter
func (r AlibabaServicecenterIdentifytaskCreateAPIRequest) GetWorkcardIds() []int64 {
	return r._workcardIds
}

// Set is OuterId Setter
// 核销单外部标识，服务商保证唯一。如果创建核销单时传入重复的outer_id，系统直接会返回服务商下该outer_id对应的核销单，不会重新创建新核销单。
func (r *AlibabaServicecenterIdentifytaskCreateAPIRequest) SetOuterId(_outerId string) error {
	r._outerId = _outerId
	r.Set("outer_id", _outerId)
	return nil
}

// Get OuterId Getter
func (r AlibabaServicecenterIdentifytaskCreateAPIRequest) GetOuterId() string {
	return r._outerId
}
