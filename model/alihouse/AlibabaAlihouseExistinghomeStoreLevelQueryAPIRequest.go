package alihouse

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseExistinghomeStoreLevelQueryAPIRequest 门店等级评分查询 API请求
// alibaba.alihouse.existinghome.store.level.query
//
// 门店等级评分查询
type AlibabaAlihouseExistinghomeStoreLevelQueryAPIRequest struct {
	model.Params
	// 外部门店id列表，不能超过200个
	_outerStoreIds string
	// 行业属性  2-二手房  3-租房
	_businessType int64
}

// NewAlibabaAlihouseExistinghomeStoreLevelQueryRequest 初始化AlibabaAlihouseExistinghomeStoreLevelQueryAPIRequest对象
func NewAlibabaAlihouseExistinghomeStoreLevelQueryRequest() *AlibabaAlihouseExistinghomeStoreLevelQueryAPIRequest {
	return &AlibabaAlihouseExistinghomeStoreLevelQueryAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihouseExistinghomeStoreLevelQueryAPIRequest) Reset() {
	r._outerStoreIds = ""
	r._businessType = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseExistinghomeStoreLevelQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.existinghome.store.level.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseExistinghomeStoreLevelQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihouseExistinghomeStoreLevelQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOuterStoreIds is OuterStoreIds Setter
// 外部门店id列表，不能超过200个
func (r *AlibabaAlihouseExistinghomeStoreLevelQueryAPIRequest) SetOuterStoreIds(_outerStoreIds string) error {
	r._outerStoreIds = _outerStoreIds
	r.Set("outer_store_ids", _outerStoreIds)
	return nil
}

// GetOuterStoreIds OuterStoreIds Getter
func (r AlibabaAlihouseExistinghomeStoreLevelQueryAPIRequest) GetOuterStoreIds() string {
	return r._outerStoreIds
}

// SetBusinessType is BusinessType Setter
// 行业属性  2-二手房  3-租房
func (r *AlibabaAlihouseExistinghomeStoreLevelQueryAPIRequest) SetBusinessType(_businessType int64) error {
	r._businessType = _businessType
	r.Set("business_type", _businessType)
	return nil
}

// GetBusinessType BusinessType Getter
func (r AlibabaAlihouseExistinghomeStoreLevelQueryAPIRequest) GetBusinessType() int64 {
	return r._businessType
}

var poolAlibabaAlihouseExistinghomeStoreLevelQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihouseExistinghomeStoreLevelQueryRequest()
	},
}

// GetAlibabaAlihouseExistinghomeStoreLevelQueryRequest 从 sync.Pool 获取 AlibabaAlihouseExistinghomeStoreLevelQueryAPIRequest
func GetAlibabaAlihouseExistinghomeStoreLevelQueryAPIRequest() *AlibabaAlihouseExistinghomeStoreLevelQueryAPIRequest {
	return poolAlibabaAlihouseExistinghomeStoreLevelQueryAPIRequest.Get().(*AlibabaAlihouseExistinghomeStoreLevelQueryAPIRequest)
}

// ReleaseAlibabaAlihouseExistinghomeStoreLevelQueryAPIRequest 将 AlibabaAlihouseExistinghomeStoreLevelQueryAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihouseExistinghomeStoreLevelQueryAPIRequest(v *AlibabaAlihouseExistinghomeStoreLevelQueryAPIRequest) {
	v.Reset()
	poolAlibabaAlihouseExistinghomeStoreLevelQueryAPIRequest.Put(v)
}
