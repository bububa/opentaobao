package ascp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabatianmaoinventoryqueryAPIRequest 阿里巴巴.天猫.aic库存.查询 API请求
// alibaba.tianmao.inventory.query
//
// 阿里巴巴.天猫.aic库存.查询
type AlibabatianmaoinventoryqueryAPIRequest struct {
	model.Params
	// 库存查询模型
	_hiErpQueryInventoryReq *HiErpQueryInventoryReq
}

// NewAlibabatianmaoinventoryqueryRequest 初始化AlibabatianmaoinventoryqueryAPIRequest对象
func NewAlibabatianmaoinventoryqueryRequest() *AlibabatianmaoinventoryqueryAPIRequest {
	return &AlibabatianmaoinventoryqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabatianmaoinventoryqueryAPIRequest) GetApiMethodName() string {
	return "alibaba.tianmao.inventory.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabatianmaoinventoryqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabatianmaoinventoryqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetHiErpQueryInventoryReq is HiErpQueryInventoryReq Setter
// 库存查询模型
func (r *AlibabatianmaoinventoryqueryAPIRequest) SetHiErpQueryInventoryReq(_hiErpQueryInventoryReq *HiErpQueryInventoryReq) error {
	r._hiErpQueryInventoryReq = _hiErpQueryInventoryReq
	r.Set("hi_erp_query_inventory_req", _hiErpQueryInventoryReq)
	return nil
}

// GetHiErpQueryInventoryReq HiErpQueryInventoryReq Getter
func (r AlibabatianmaoinventoryqueryAPIRequest) GetHiErpQueryInventoryReq() *HiErpQueryInventoryReq {
	return r._hiErpQueryInventoryReq
}
