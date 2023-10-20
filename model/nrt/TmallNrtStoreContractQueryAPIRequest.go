package nrt

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallNrtStoreContractQueryAPIRequest 摊位合同查询接口 API请求
// tmall.nrt.store.contract.query
//
// 摊位合同查询接口
type TmallNrtStoreContractQueryAPIRequest struct {
	model.Params
	// 参数对象
	_eaStoreContractQuery *NrtStoreQueryDto
}

// NewTmallNrtStoreContractQueryRequest 初始化TmallNrtStoreContractQueryAPIRequest对象
func NewTmallNrtStoreContractQueryRequest() *TmallNrtStoreContractQueryAPIRequest {
	return &TmallNrtStoreContractQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallNrtStoreContractQueryAPIRequest) GetApiMethodName() string {
	return "tmall.nrt.store.contract.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallNrtStoreContractQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallNrtStoreContractQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetEaStoreContractQuery is EaStoreContractQuery Setter
// 参数对象
func (r *TmallNrtStoreContractQueryAPIRequest) SetEaStoreContractQuery(_eaStoreContractQuery *NrtStoreQueryDto) error {
	r._eaStoreContractQuery = _eaStoreContractQuery
	r.Set("ea_store_contract_query", _eaStoreContractQuery)
	return nil
}

// GetEaStoreContractQuery EaStoreContractQuery Getter
func (r TmallNrtStoreContractQueryAPIRequest) GetEaStoreContractQuery() *NrtStoreQueryDto {
	return r._eaStoreContractQuery
}
