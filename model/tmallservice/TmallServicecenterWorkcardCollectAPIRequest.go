package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallservicecenterworkcardcollectAPIRequest 工单揽件 API请求
// tmall.servicecenter.workcard.collect
//
// 服务工单揽件接口
type TmallservicecenterworkcardcollectAPIRequest struct {
	model.Params
	// 扩展信息
	_attributes string
	// 工单id
	_workcardId int64
	// 买家id
	_buyerId int64
}

// NewTmallservicecenterworkcardcollectRequest 初始化TmallservicecenterworkcardcollectAPIRequest对象
func NewTmallservicecenterworkcardcollectRequest() *TmallservicecenterworkcardcollectAPIRequest {
	return &TmallservicecenterworkcardcollectAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallservicecenterworkcardcollectAPIRequest) GetApiMethodName() string {
	return "tmall.servicecenter.workcard.collect"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallservicecenterworkcardcollectAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallservicecenterworkcardcollectAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAttributes is Attributes Setter
// 扩展信息
func (r *TmallservicecenterworkcardcollectAPIRequest) SetAttributes(_attributes string) error {
	r._attributes = _attributes
	r.Set("attributes", _attributes)
	return nil
}

// GetAttributes Attributes Getter
func (r TmallservicecenterworkcardcollectAPIRequest) GetAttributes() string {
	return r._attributes
}

// SetWorkcardId is WorkcardId Setter
// 工单id
func (r *TmallservicecenterworkcardcollectAPIRequest) SetWorkcardId(_workcardId int64) error {
	r._workcardId = _workcardId
	r.Set("workcard_id", _workcardId)
	return nil
}

// GetWorkcardId WorkcardId Getter
func (r TmallservicecenterworkcardcollectAPIRequest) GetWorkcardId() int64 {
	return r._workcardId
}

// SetBuyerId is BuyerId Setter
// 买家id
func (r *TmallservicecenterworkcardcollectAPIRequest) SetBuyerId(_buyerId int64) error {
	r._buyerId = _buyerId
	r.Set("buyer_id", _buyerId)
	return nil
}

// GetBuyerId BuyerId Getter
func (r TmallservicecenterworkcardcollectAPIRequest) GetBuyerId() int64 {
	return r._buyerId
}
