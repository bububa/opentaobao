package alihealthlab

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthLabItemStoreRelationSyncAPIRequest 检验检测业务，isv项目门店关系同步 API请求
// alibaba.alihealth.lab.item.store.relation.sync
//
// 阿里健康检验检测业务，isv检验检测项目门店关系同步到健康，支持检验检测项目门店关系的增加和删除
type AlibabaAlihealthLabItemStoreRelationSyncAPIRequest struct {
	model.Params
	// EFFECTIVE 有效，INVALID 无效
	_isvRelationStatus string
	// isv门店编码
	_isvStoreCodes []string
	// 检验检测项目isv侧编码
	_isvItemCode string
}

// NewAlibabaAlihealthLabItemStoreRelationSyncRequest 初始化AlibabaAlihealthLabItemStoreRelationSyncAPIRequest对象
func NewAlibabaAlihealthLabItemStoreRelationSyncRequest() *AlibabaAlihealthLabItemStoreRelationSyncAPIRequest {
	return &AlibabaAlihealthLabItemStoreRelationSyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthLabItemStoreRelationSyncAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.lab.item.store.relation.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthLabItemStoreRelationSyncAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
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
