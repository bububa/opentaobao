package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaservicecenterfulfiltaskcreateAPIRequest 合单生成核销单 API请求
// alibaba.servicecenter.fulfiltask.create
//
// 服务对工单进行合单，合单的结果是生成核销单
type AlibabaservicecenterfulfiltaskcreateAPIRequest struct {
	model.Params
	// 工单id列表
	_workcardIds []int64
	// 外部单号
	_outerId string
}

// NewAlibabaservicecenterfulfiltaskcreateRequest 初始化AlibabaservicecenterfulfiltaskcreateAPIRequest对象
func NewAlibabaservicecenterfulfiltaskcreateRequest() *AlibabaservicecenterfulfiltaskcreateAPIRequest {
	return &AlibabaservicecenterfulfiltaskcreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaservicecenterfulfiltaskcreateAPIRequest) GetApiMethodName() string {
	return "alibaba.servicecenter.fulfiltask.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaservicecenterfulfiltaskcreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaservicecenterfulfiltaskcreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetWorkcardIds is WorkcardIds Setter
// 工单id列表
func (r *AlibabaservicecenterfulfiltaskcreateAPIRequest) SetWorkcardIds(_workcardIds []int64) error {
	r._workcardIds = _workcardIds
	r.Set("workcard_ids", _workcardIds)
	return nil
}

// GetWorkcardIds WorkcardIds Getter
func (r AlibabaservicecenterfulfiltaskcreateAPIRequest) GetWorkcardIds() []int64 {
	return r._workcardIds
}

// SetOuterId is OuterId Setter
// 外部单号
func (r *AlibabaservicecenterfulfiltaskcreateAPIRequest) SetOuterId(_outerId string) error {
	r._outerId = _outerId
	r.Set("outer_id", _outerId)
	return nil
}

// GetOuterId OuterId Getter
func (r AlibabaservicecenterfulfiltaskcreateAPIRequest) GetOuterId() string {
	return r._outerId
}
