package nrt

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallNrtStoreContractSynAPIRequest 喵零合同同步 API请求
// tmall.nrt.store.contract.syn
//
// 喵零合同同步
type TmallNrtStoreContractSynAPIRequest struct {
	model.Params
	// 合同参数对象
	_eaStoreContractDto *NrtStoreContractDto
}

// NewTmallNrtStoreContractSynRequest 初始化TmallNrtStoreContractSynAPIRequest对象
func NewTmallNrtStoreContractSynRequest() *TmallNrtStoreContractSynAPIRequest {
	return &TmallNrtStoreContractSynAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallNrtStoreContractSynAPIRequest) GetApiMethodName() string {
	return "tmall.nrt.store.contract.syn"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallNrtStoreContractSynAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetEaStoreContractDto is EaStoreContractDto Setter
// 合同参数对象
func (r *TmallNrtStoreContractSynAPIRequest) SetEaStoreContractDto(_eaStoreContractDto *NrtStoreContractDto) error {
	r._eaStoreContractDto = _eaStoreContractDto
	r.Set("ea_store_contract_dto", _eaStoreContractDto)
	return nil
}

// GetEaStoreContractDto EaStoreContractDto Getter
func (r TmallNrtStoreContractSynAPIRequest) GetEaStoreContractDto() *NrtStoreContractDto {
	return r._eaStoreContractDto
}
