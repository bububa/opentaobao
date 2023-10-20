package nrt

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallNrtStoreContractSynAPIRequest) Reset() {
	r._eaStoreContractDto = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallNrtStoreContractSynAPIRequest) GetApiMethodName() string {
	return "tmall.nrt.store.contract.syn"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallNrtStoreContractSynAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallNrtStoreContractSynAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolTmallNrtStoreContractSynAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallNrtStoreContractSynRequest()
	},
}

// GetTmallNrtStoreContractSynRequest 从 sync.Pool 获取 TmallNrtStoreContractSynAPIRequest
func GetTmallNrtStoreContractSynAPIRequest() *TmallNrtStoreContractSynAPIRequest {
	return poolTmallNrtStoreContractSynAPIRequest.Get().(*TmallNrtStoreContractSynAPIRequest)
}

// ReleaseTmallNrtStoreContractSynAPIRequest 将 TmallNrtStoreContractSynAPIRequest 放入 sync.Pool
func ReleaseTmallNrtStoreContractSynAPIRequest(v *TmallNrtStoreContractSynAPIRequest) {
	v.Reset()
	poolTmallNrtStoreContractSynAPIRequest.Put(v)
}
