package ascp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabatianmaoinventorymodifyAPIRequest 阿里巴巴.天猫.aic库存.修改 API请求
// alibaba.tianmao.inventory.modify
//
// 阿里巴巴.天猫.aic库存.修改
type AlibabatianmaoinventorymodifyAPIRequest struct {
	model.Params
	// 修改库存模型
	_hiErpModifyInventoryReq *HiErpModifyInventoryReq
}

// NewAlibabatianmaoinventorymodifyRequest 初始化AlibabatianmaoinventorymodifyAPIRequest对象
func NewAlibabatianmaoinventorymodifyRequest() *AlibabatianmaoinventorymodifyAPIRequest {
	return &AlibabatianmaoinventorymodifyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabatianmaoinventorymodifyAPIRequest) GetApiMethodName() string {
	return "alibaba.tianmao.inventory.modify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabatianmaoinventorymodifyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabatianmaoinventorymodifyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetHiErpModifyInventoryReq is HiErpModifyInventoryReq Setter
// 修改库存模型
func (r *AlibabatianmaoinventorymodifyAPIRequest) SetHiErpModifyInventoryReq(_hiErpModifyInventoryReq *HiErpModifyInventoryReq) error {
	r._hiErpModifyInventoryReq = _hiErpModifyInventoryReq
	r.Set("hi_erp_modify_inventory_req", _hiErpModifyInventoryReq)
	return nil
}

// GetHiErpModifyInventoryReq HiErpModifyInventoryReq Getter
func (r AlibabatianmaoinventorymodifyAPIRequest) GetHiErpModifyInventoryReq() *HiErpModifyInventoryReq {
	return r._hiErpModifyInventoryReq
}
