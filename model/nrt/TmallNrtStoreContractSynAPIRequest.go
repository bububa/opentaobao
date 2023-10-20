package nrt

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallnrtstorecontractsynAPIRequest 喵零合同同步 API请求
// tmall.nrt.store.contract.syn
//
// 喵零合同同步
type TmallnrtstorecontractsynAPIRequest struct {
	model.Params
	// 合同参数对象
	_eaStoreContractDto *NrtStoreContractDto
}

// NewTmallnrtstorecontractsynRequest 初始化TmallnrtstorecontractsynAPIRequest对象
func NewTmallnrtstorecontractsynRequest() *TmallnrtstorecontractsynAPIRequest {
	return &TmallnrtstorecontractsynAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallnrtstorecontractsynAPIRequest) GetApiMethodName() string {
	return "tmall.nrt.store.contract.syn"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallnrtstorecontractsynAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallnrtstorecontractsynAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetEaStoreContractDto is EaStoreContractDto Setter
// 合同参数对象
func (r *TmallnrtstorecontractsynAPIRequest) SetEaStoreContractDto(_eaStoreContractDto *NrtStoreContractDto) error {
	r._eaStoreContractDto = _eaStoreContractDto
	r.Set("ea_store_contract_dto", _eaStoreContractDto)
	return nil
}

// GetEaStoreContractDto EaStoreContractDto Getter
func (r TmallnrtstorecontractsynAPIRequest) GetEaStoreContractDto() *NrtStoreContractDto {
	return r._eaStoreContractDto
}
