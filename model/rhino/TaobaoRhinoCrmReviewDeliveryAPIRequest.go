package rhino

import (
	"net/url"

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
		Params: model.NewParams(),
	}
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
