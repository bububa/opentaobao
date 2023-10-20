package rhino

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoRhinoCrmReviewDeliveryAPIRequest crm实体预询期 API请求
// taobao.rhino.crm.review.delivery
//
// crm实体预询期
type TaobaoRhinoCrmReviewDeliveryAPIRequest struct {
	model.Params
	// crm实体数据
	_crmEntity *CrmEntity
}

// NewTaobaoRhinoCrmReviewDeliveryRequest 初始化TaobaoRhinoCrmReviewDeliveryAPIRequest对象
func NewTaobaoRhinoCrmReviewDeliveryRequest() *TaobaoRhinoCrmReviewDeliveryAPIRequest {
	return &TaobaoRhinoCrmReviewDeliveryAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoRhinoCrmReviewDeliveryAPIRequest) Reset() {
	r._crmEntity = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoRhinoCrmReviewDeliveryAPIRequest) GetApiMethodName() string {
	return "taobao.rhino.crm.review.delivery"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoRhinoCrmReviewDeliveryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoRhinoCrmReviewDeliveryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCrmEntity is CrmEntity Setter
// crm实体数据
func (r *TaobaoRhinoCrmReviewDeliveryAPIRequest) SetCrmEntity(_crmEntity *CrmEntity) error {
	r._crmEntity = _crmEntity
	r.Set("crm_entity", _crmEntity)
	return nil
}

// GetCrmEntity CrmEntity Getter
func (r TaobaoRhinoCrmReviewDeliveryAPIRequest) GetCrmEntity() *CrmEntity {
	return r._crmEntity
}

var poolTaobaoRhinoCrmReviewDeliveryAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoRhinoCrmReviewDeliveryRequest()
	},
}

// GetTaobaoRhinoCrmReviewDeliveryRequest 从 sync.Pool 获取 TaobaoRhinoCrmReviewDeliveryAPIRequest
func GetTaobaoRhinoCrmReviewDeliveryAPIRequest() *TaobaoRhinoCrmReviewDeliveryAPIRequest {
	return poolTaobaoRhinoCrmReviewDeliveryAPIRequest.Get().(*TaobaoRhinoCrmReviewDeliveryAPIRequest)
}

// ReleaseTaobaoRhinoCrmReviewDeliveryAPIRequest 将 TaobaoRhinoCrmReviewDeliveryAPIRequest 放入 sync.Pool
func ReleaseTaobaoRhinoCrmReviewDeliveryAPIRequest(v *TaobaoRhinoCrmReviewDeliveryAPIRequest) {
	v.Reset()
	poolTaobaoRhinoCrmReviewDeliveryAPIRequest.Put(v)
}
