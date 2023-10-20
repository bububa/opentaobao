package alihealthlab

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthLabItemStoreRelationSyncAPIRequest 检验检测业务，isv项目门店关系同步 API请求
// alibaba.alihealth.lab.item.store.relation.sync
//
// 阿里健康检验检测业务，isv检验检测项目门店关系同步到健康，支持检验检测项目门店关系的增加和删除
type AlibabaAlihealthLabItemStoreRelationSyncAPIRequest struct {
	model.Params
	// isv门店编码
	_isvStoreCodes []string
	// EFFECTIVE 有效，INVALID 无效
	_isvRelationStatus string
	// 检验检测项目isv侧编码
	_isvItemCode string
}

// NewAlibabaAlihealthLabItemStoreRelationSyncRequest 初始化AlibabaAlihealthLabItemStoreRelationSyncAPIRequest对象
func NewAlibabaAlihealthLabItemStoreRelationSyncRequest() *AlibabaAlihealthLabItemStoreRelationSyncAPIRequest {
	return &AlibabaAlihealthLabItemStoreRelationSyncAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthLabItemStoreRelationSyncAPIRequest) Reset() {
	r._isvStoreCodes = r._isvStoreCodes[:0]
	r._isvRelationStatus = ""
	r._isvItemCode = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthLabItemStoreRelationSyncAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.lab.item.store.relation.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthLabItemStoreRelationSyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthLabItemStoreRelationSyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetIsvStoreCodes is IsvStoreCodes Setter
// isv门店编码
func (r *AlibabaAlihealthLabItemStoreRelationSyncAPIRequest) SetIsvStoreCodes(_isvStoreCodes []string) error {
	r._isvStoreCodes = _isvStoreCodes
	r.Set("isv_store_codes", _isvStoreCodes)
	return nil
}

// GetIsvStoreCodes IsvStoreCodes Getter
func (r AlibabaAlihealthLabItemStoreRelationSyncAPIRequest) GetIsvStoreCodes() []string {
	return r._isvStoreCodes
}

// SetIsvRelationStatus is IsvRelationStatus Setter
// EFFECTIVE 有效，INVALID 无效
func (r *AlibabaAlihealthLabItemStoreRelationSyncAPIRequest) SetIsvRelationStatus(_isvRelationStatus string) error {
	r._isvRelationStatus = _isvRelationStatus
	r.Set("isv_relation_status", _isvRelationStatus)
	return nil
}

// GetIsvRelationStatus IsvRelationStatus Getter
func (r AlibabaAlihealthLabItemStoreRelationSyncAPIRequest) GetIsvRelationStatus() string {
	return r._isvRelationStatus
}

// SetIsvItemCode is IsvItemCode Setter
// 检验检测项目isv侧编码
func (r *AlibabaAlihealthLabItemStoreRelationSyncAPIRequest) SetIsvItemCode(_isvItemCode string) error {
	r._isvItemCode = _isvItemCode
	r.Set("isv_item_code", _isvItemCode)
	return nil
}

// GetIsvItemCode IsvItemCode Getter
func (r AlibabaAlihealthLabItemStoreRelationSyncAPIRequest) GetIsvItemCode() string {
	return r._isvItemCode
}

var poolAlibabaAlihealthLabItemStoreRelationSyncAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthLabItemStoreRelationSyncRequest()
	},
}

// GetAlibabaAlihealthLabItemStoreRelationSyncRequest 从 sync.Pool 获取 AlibabaAlihealthLabItemStoreRelationSyncAPIRequest
func GetAlibabaAlihealthLabItemStoreRelationSyncAPIRequest() *AlibabaAlihealthLabItemStoreRelationSyncAPIRequest {
	return poolAlibabaAlihealthLabItemStoreRelationSyncAPIRequest.Get().(*AlibabaAlihealthLabItemStoreRelationSyncAPIRequest)
}

// ReleaseAlibabaAlihealthLabItemStoreRelationSyncAPIRequest 将 AlibabaAlihealthLabItemStoreRelationSyncAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthLabItemStoreRelationSyncAPIRequest(v *AlibabaAlihealthLabItemStoreRelationSyncAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthLabItemStoreRelationSyncAPIRequest.Put(v)
}
