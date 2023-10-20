package alidoc

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthAlidocDrugStoreAddAPIRequest gsk新增药店 API请求
// alibaba.alihealth.alidoc.drug.store.add
//
// GSK上传药店信息
type AlibabaAlihealthAlidocDrugStoreAddAPIRequest struct {
	model.Params
	// 新增药店
	_drugStoreAddTopRequest *DrugStoreAddTopRequest
}

// NewAlibabaAlihealthAlidocDrugStoreAddRequest 初始化AlibabaAlihealthAlidocDrugStoreAddAPIRequest对象
func NewAlibabaAlihealthAlidocDrugStoreAddRequest() *AlibabaAlihealthAlidocDrugStoreAddAPIRequest {
	return &AlibabaAlihealthAlidocDrugStoreAddAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthAlidocDrugStoreAddAPIRequest) Reset() {
	r._drugStoreAddTopRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthAlidocDrugStoreAddAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.alidoc.drug.store.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthAlidocDrugStoreAddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthAlidocDrugStoreAddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDrugStoreAddTopRequest is DrugStoreAddTopRequest Setter
// 新增药店
func (r *AlibabaAlihealthAlidocDrugStoreAddAPIRequest) SetDrugStoreAddTopRequest(_drugStoreAddTopRequest *DrugStoreAddTopRequest) error {
	r._drugStoreAddTopRequest = _drugStoreAddTopRequest
	r.Set("drug_store_add_top_request", _drugStoreAddTopRequest)
	return nil
}

// GetDrugStoreAddTopRequest DrugStoreAddTopRequest Getter
func (r AlibabaAlihealthAlidocDrugStoreAddAPIRequest) GetDrugStoreAddTopRequest() *DrugStoreAddTopRequest {
	return r._drugStoreAddTopRequest
}

var poolAlibabaAlihealthAlidocDrugStoreAddAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthAlidocDrugStoreAddRequest()
	},
}

// GetAlibabaAlihealthAlidocDrugStoreAddRequest 从 sync.Pool 获取 AlibabaAlihealthAlidocDrugStoreAddAPIRequest
func GetAlibabaAlihealthAlidocDrugStoreAddAPIRequest() *AlibabaAlihealthAlidocDrugStoreAddAPIRequest {
	return poolAlibabaAlihealthAlidocDrugStoreAddAPIRequest.Get().(*AlibabaAlihealthAlidocDrugStoreAddAPIRequest)
}

// ReleaseAlibabaAlihealthAlidocDrugStoreAddAPIRequest 将 AlibabaAlihealthAlidocDrugStoreAddAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthAlidocDrugStoreAddAPIRequest(v *AlibabaAlihealthAlidocDrugStoreAddAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthAlidocDrugStoreAddAPIRequest.Put(v)
}
