package alidoc

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthAlidocDrugStoreUpdateAPIRequest 更新药店 API请求
// alibaba.alihealth.alidoc.drug.store.update
//
// 药店信息更新接口
type AlibabaAlihealthAlidocDrugStoreUpdateAPIRequest struct {
	model.Params
	// 更新对象
	_drugStoreUpdateTopRequest *DrugStoreUpdateTopRequest
}

// NewAlibabaAlihealthAlidocDrugStoreUpdateRequest 初始化AlibabaAlihealthAlidocDrugStoreUpdateAPIRequest对象
func NewAlibabaAlihealthAlidocDrugStoreUpdateRequest() *AlibabaAlihealthAlidocDrugStoreUpdateAPIRequest {
	return &AlibabaAlihealthAlidocDrugStoreUpdateAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthAlidocDrugStoreUpdateAPIRequest) Reset() {
	r._drugStoreUpdateTopRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthAlidocDrugStoreUpdateAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.alidoc.drug.store.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthAlidocDrugStoreUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthAlidocDrugStoreUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDrugStoreUpdateTopRequest is DrugStoreUpdateTopRequest Setter
// 更新对象
func (r *AlibabaAlihealthAlidocDrugStoreUpdateAPIRequest) SetDrugStoreUpdateTopRequest(_drugStoreUpdateTopRequest *DrugStoreUpdateTopRequest) error {
	r._drugStoreUpdateTopRequest = _drugStoreUpdateTopRequest
	r.Set("drug_store_update_top_request", _drugStoreUpdateTopRequest)
	return nil
}

// GetDrugStoreUpdateTopRequest DrugStoreUpdateTopRequest Getter
func (r AlibabaAlihealthAlidocDrugStoreUpdateAPIRequest) GetDrugStoreUpdateTopRequest() *DrugStoreUpdateTopRequest {
	return r._drugStoreUpdateTopRequest
}

var poolAlibabaAlihealthAlidocDrugStoreUpdateAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthAlidocDrugStoreUpdateRequest()
	},
}

// GetAlibabaAlihealthAlidocDrugStoreUpdateRequest 从 sync.Pool 获取 AlibabaAlihealthAlidocDrugStoreUpdateAPIRequest
func GetAlibabaAlihealthAlidocDrugStoreUpdateAPIRequest() *AlibabaAlihealthAlidocDrugStoreUpdateAPIRequest {
	return poolAlibabaAlihealthAlidocDrugStoreUpdateAPIRequest.Get().(*AlibabaAlihealthAlidocDrugStoreUpdateAPIRequest)
}

// ReleaseAlibabaAlihealthAlidocDrugStoreUpdateAPIRequest 将 AlibabaAlihealthAlidocDrugStoreUpdateAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthAlidocDrugStoreUpdateAPIRequest(v *AlibabaAlihealthAlidocDrugStoreUpdateAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthAlidocDrugStoreUpdateAPIRequest.Put(v)
}
