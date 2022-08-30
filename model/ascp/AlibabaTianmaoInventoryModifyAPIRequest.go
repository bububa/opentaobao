package ascp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaTianmaoInventoryModifyAPIRequest 阿里巴巴.天猫.aic库存.修改 API请求
// alibaba.tianmao.inventory.modify
//
// 阿里巴巴.天猫.aic库存.修改
type AlibabaTianmaoInventoryModifyAPIRequest struct {
	model.Params
	// 修改库存模型
	_hiErpModifyInventoryReq *HiErpModifyInventoryReq
}

// NewAlibabaTianmaoInventoryModifyRequest 初始化AlibabaTianmaoInventoryModifyAPIRequest对象
func NewAlibabaTianmaoInventoryModifyRequest() *AlibabaTianmaoInventoryModifyAPIRequest {
	return &AlibabaTianmaoInventoryModifyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaTianmaoInventoryModifyAPIRequest) GetApiMethodName() string {
	return "alibaba.tianmao.inventory.modify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaTianmaoInventoryModifyAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetHiErpModifyInventoryReq is HiErpModifyInventoryReq Setter
// 修改库存模型
func (r *AlibabaTianmaoInventoryModifyAPIRequest) SetHiErpModifyInventoryReq(_hiErpModifyInventoryReq *HiErpModifyInventoryReq) error {
	r._hiErpModifyInventoryReq = _hiErpModifyInventoryReq
	r.Set("hi_erp_modify_inventory_req", _hiErpModifyInventoryReq)
	return nil
}

// GetHiErpModifyInventoryReq HiErpModifyInventoryReq Getter
func (r AlibabaTianmaoInventoryModifyAPIRequest) GetHiErpModifyInventoryReq() *HiErpModifyInventoryReq {
	return r._hiErpModifyInventoryReq
}
