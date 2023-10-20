package rhino

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaorhinocrmreviewdeliveryAPIRequest crm实体预询期 API请求
// taobao.rhino.crm.review.delivery
//
// crm实体预询期
type TaobaorhinocrmreviewdeliveryAPIRequest struct {
	model.Params
	// crm实体数据
	_crmEntity *CrmEntity
}

// NewTaobaorhinocrmreviewdeliveryRequest 初始化TaobaorhinocrmreviewdeliveryAPIRequest对象
func NewTaobaorhinocrmreviewdeliveryRequest() *TaobaorhinocrmreviewdeliveryAPIRequest {
	return &TaobaorhinocrmreviewdeliveryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaorhinocrmreviewdeliveryAPIRequest) GetApiMethodName() string {
	return "taobao.rhino.crm.review.delivery"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaorhinocrmreviewdeliveryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaorhinocrmreviewdeliveryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCrmEntity is CrmEntity Setter
// crm实体数据
func (r *TaobaorhinocrmreviewdeliveryAPIRequest) SetCrmEntity(_crmEntity *CrmEntity) error {
	r._crmEntity = _crmEntity
	r.Set("crm_entity", _crmEntity)
	return nil
}

// GetCrmEntity CrmEntity Getter
func (r TaobaorhinocrmreviewdeliveryAPIRequest) GetCrmEntity() *CrmEntity {
	return r._crmEntity
}
