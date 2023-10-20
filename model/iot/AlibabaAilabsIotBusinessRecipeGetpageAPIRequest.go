package iot

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAilabsIotBusinessRecipeGetpageAPIRequest 分页查询食谱 API请求
// alibaba.ailabs.iot.business.recipe.getpage
//
// 分页查询食谱数据
type AlibabaAilabsIotBusinessRecipeGetpageAPIRequest struct {
	model.Params
	// 开放账号id
	_openAccountId string
	// 分页页码
	_pageNum int64
	// 分页大小
	_pageSize int64
}

// NewAlibabaAilabsIotBusinessRecipeGetpageRequest 初始化AlibabaAilabsIotBusinessRecipeGetpageAPIRequest对象
func NewAlibabaAilabsIotBusinessRecipeGetpageRequest() *AlibabaAilabsIotBusinessRecipeGetpageAPIRequest {
	return &AlibabaAilabsIotBusinessRecipeGetpageAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAilabsIotBusinessRecipeGetpageAPIRequest) Reset() {
	r._openAccountId = ""
	r._pageNum = 0
	r._pageSize = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAilabsIotBusinessRecipeGetpageAPIRequest) GetApiMethodName() string {
	return "alibaba.ailabs.iot.business.recipe.getpage"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAilabsIotBusinessRecipeGetpageAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAilabsIotBusinessRecipeGetpageAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOpenAccountId is OpenAccountId Setter
// 开放账号id
func (r *AlibabaAilabsIotBusinessRecipeGetpageAPIRequest) SetOpenAccountId(_openAccountId string) error {
	r._openAccountId = _openAccountId
	r.Set("open_account_id", _openAccountId)
	return nil
}

// GetOpenAccountId OpenAccountId Getter
func (r AlibabaAilabsIotBusinessRecipeGetpageAPIRequest) GetOpenAccountId() string {
	return r._openAccountId
}

// SetPageNum is PageNum Setter
// 分页页码
func (r *AlibabaAilabsIotBusinessRecipeGetpageAPIRequest) SetPageNum(_pageNum int64) error {
	r._pageNum = _pageNum
	r.Set("page_num", _pageNum)
	return nil
}

// GetPageNum PageNum Getter
func (r AlibabaAilabsIotBusinessRecipeGetpageAPIRequest) GetPageNum() int64 {
	return r._pageNum
}

// SetPageSize is PageSize Setter
// 分页大小
func (r *AlibabaAilabsIotBusinessRecipeGetpageAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r AlibabaAilabsIotBusinessRecipeGetpageAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

var poolAlibabaAilabsIotBusinessRecipeGetpageAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAilabsIotBusinessRecipeGetpageRequest()
	},
}

// GetAlibabaAilabsIotBusinessRecipeGetpageRequest 从 sync.Pool 获取 AlibabaAilabsIotBusinessRecipeGetpageAPIRequest
func GetAlibabaAilabsIotBusinessRecipeGetpageAPIRequest() *AlibabaAilabsIotBusinessRecipeGetpageAPIRequest {
	return poolAlibabaAilabsIotBusinessRecipeGetpageAPIRequest.Get().(*AlibabaAilabsIotBusinessRecipeGetpageAPIRequest)
}

// ReleaseAlibabaAilabsIotBusinessRecipeGetpageAPIRequest 将 AlibabaAilabsIotBusinessRecipeGetpageAPIRequest 放入 sync.Pool
func ReleaseAlibabaAilabsIotBusinessRecipeGetpageAPIRequest(v *AlibabaAilabsIotBusinessRecipeGetpageAPIRequest) {
	v.Reset()
	poolAlibabaAilabsIotBusinessRecipeGetpageAPIRequest.Put(v)
}
