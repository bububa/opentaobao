package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallServicecenterIdentifytaskCreateAPIRequest 服务商创建核销单 API请求
// tmall.servicecenter.identifytask.create
//
// 服务商调用该接口进行创建核销单操作
type TmallServicecenterIdentifytaskCreateAPIRequest struct {
	model.Params
	// 工单列表
	_workcardIds []int64
	// 服务商自定义的外部核销单id
	_outerId string
	// 是否改派
	_reassign bool
}

// NewTmallServicecenterIdentifytaskCreateRequest 初始化TmallServicecenterIdentifytaskCreateAPIRequest对象
func NewTmallServicecenterIdentifytaskCreateRequest() *TmallServicecenterIdentifytaskCreateAPIRequest {
	return &TmallServicecenterIdentifytaskCreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallServicecenterIdentifytaskCreateAPIRequest) GetApiMethodName() string {
	return "tmall.servicecenter.identifytask.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallServicecenterIdentifytaskCreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallServicecenterIdentifytaskCreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetWorkcardIds is WorkcardIds Setter
// 工单列表
func (r *TmallServicecenterIdentifytaskCreateAPIRequest) SetWorkcardIds(_workcardIds []int64) error {
	r._workcardIds = _workcardIds
	r.Set("workcard_ids", _workcardIds)
	return nil
}

// GetWorkcardIds WorkcardIds Getter
func (r TmallServicecenterIdentifytaskCreateAPIRequest) GetWorkcardIds() []int64 {
	return r._workcardIds
}

// SetOuterId is OuterId Setter
// 服务商自定义的外部核销单id
func (r *TmallServicecenterIdentifytaskCreateAPIRequest) SetOuterId(_outerId string) error {
	r._outerId = _outerId
	r.Set("outer_id", _outerId)
	return nil
}

// GetOuterId OuterId Getter
func (r TmallServicecenterIdentifytaskCreateAPIRequest) GetOuterId() string {
	return r._outerId
}

// SetReassign is Reassign Setter
// 是否改派
func (r *TmallServicecenterIdentifytaskCreateAPIRequest) SetReassign(_reassign bool) error {
	r._reassign = _reassign
	r.Set("reassign", _reassign)
	return nil
}

// GetReassign Reassign Getter
func (r TmallServicecenterIdentifytaskCreateAPIRequest) GetReassign() bool {
	return r._reassign
}
