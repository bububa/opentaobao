package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallServicecenterWorkcardCollectAPIRequest 工单揽件 API请求
// tmall.servicecenter.workcard.collect
//
// 服务工单揽件接口
type TmallServicecenterWorkcardCollectAPIRequest struct {
	model.Params
	// 扩展信息
	_attributes string
	// 工单id
	_workcardId int64
	// 买家id
	_buyerId int64
}

// NewTmallServicecenterWorkcardCollectRequest 初始化TmallServicecenterWorkcardCollectAPIRequest对象
func NewTmallServicecenterWorkcardCollectRequest() *TmallServicecenterWorkcardCollectAPIRequest {
	return &TmallServicecenterWorkcardCollectAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallServicecenterWorkcardCollectAPIRequest) GetApiMethodName() string {
	return "tmall.servicecenter.workcard.collect"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallServicecenterWorkcardCollectAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallServicecenterWorkcardCollectAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAttributes is Attributes Setter
// 扩展信息
func (r *TmallServicecenterWorkcardCollectAPIRequest) SetAttributes(_attributes string) error {
	r._attributes = _attributes
	r.Set("attributes", _attributes)
	return nil
}

// GetAttributes Attributes Getter
func (r TmallServicecenterWorkcardCollectAPIRequest) GetAttributes() string {
	return r._attributes
}

// SetWorkcardId is WorkcardId Setter
// 工单id
func (r *TmallServicecenterWorkcardCollectAPIRequest) SetWorkcardId(_workcardId int64) error {
	r._workcardId = _workcardId
	r.Set("workcard_id", _workcardId)
	return nil
}

// GetWorkcardId WorkcardId Getter
func (r TmallServicecenterWorkcardCollectAPIRequest) GetWorkcardId() int64 {
	return r._workcardId
}

// SetBuyerId is BuyerId Setter
// 买家id
func (r *TmallServicecenterWorkcardCollectAPIRequest) SetBuyerId(_buyerId int64) error {
	r._buyerId = _buyerId
	r.Set("buyer_id", _buyerId)
	return nil
}

// GetBuyerId BuyerId Getter
func (r TmallServicecenterWorkcardCollectAPIRequest) GetBuyerId() int64 {
	return r._buyerId
}
