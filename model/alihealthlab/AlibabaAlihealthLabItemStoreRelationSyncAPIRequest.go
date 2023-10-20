package alihealthlab

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthlabitemstorerelationsyncAPIRequest 检验检测业务，isv项目门店关系同步 API请求
// alibaba.alihealth.lab.item.store.relation.sync
//
// 阿里健康检验检测业务，isv检验检测项目门店关系同步到健康，支持检验检测项目门店关系的增加和删除
type AlibabaalihealthlabitemstorerelationsyncAPIRequest struct {
	model.Params
	// isv门店编码
	_isvStoreCodes []string
	// EFFECTIVE 有效，INVALID 无效
	_isvRelationStatus string
	// 检验检测项目isv侧编码
	_isvItemCode string
}

// NewAlibabaalihealthlabitemstorerelationsyncRequest 初始化AlibabaalihealthlabitemstorerelationsyncAPIRequest对象
func NewAlibabaalihealthlabitemstorerelationsyncRequest() *AlibabaalihealthlabitemstorerelationsyncAPIRequest {
	return &AlibabaalihealthlabitemstorerelationsyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthlabitemstorerelationsyncAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.lab.item.store.relation.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthlabitemstorerelationsyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthlabitemstorerelationsyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetIsvStoreCodes is IsvStoreCodes Setter
// isv门店编码
func (r *AlibabaalihealthlabitemstorerelationsyncAPIRequest) SetIsvStoreCodes(_isvStoreCodes []string) error {
	r._isvStoreCodes = _isvStoreCodes
	r.Set("isv_store_codes", _isvStoreCodes)
	return nil
}

// GetIsvStoreCodes IsvStoreCodes Getter
func (r AlibabaalihealthlabitemstorerelationsyncAPIRequest) GetIsvStoreCodes() []string {
	return r._isvStoreCodes
}

// SetIsvRelationStatus is IsvRelationStatus Setter
// EFFECTIVE 有效，INVALID 无效
func (r *AlibabaalihealthlabitemstorerelationsyncAPIRequest) SetIsvRelationStatus(_isvRelationStatus string) error {
	r._isvRelationStatus = _isvRelationStatus
	r.Set("isv_relation_status", _isvRelationStatus)
	return nil
}

// GetIsvRelationStatus IsvRelationStatus Getter
func (r AlibabaalihealthlabitemstorerelationsyncAPIRequest) GetIsvRelationStatus() string {
	return r._isvRelationStatus
}

// SetIsvItemCode is IsvItemCode Setter
// 检验检测项目isv侧编码
func (r *AlibabaalihealthlabitemstorerelationsyncAPIRequest) SetIsvItemCode(_isvItemCode string) error {
	r._isvItemCode = _isvItemCode
	r.Set("isv_item_code", _isvItemCode)
	return nil
}

// GetIsvItemCode IsvItemCode Getter
func (r AlibabaalihealthlabitemstorerelationsyncAPIRequest) GetIsvItemCode() string {
	return r._isvItemCode
}
