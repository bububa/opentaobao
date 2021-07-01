package iot

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAilabsIotBusinessRecipeGetpageAPIRequest
分页查询食谱 API请求
alibaba.ailabs.iot.business.recipe.getpage

分页查询食谱数据 */
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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAilabsIotBusinessRecipeGetpageAPIRequest) GetApiMethodName() string {
	return "alibaba.ailabs.iot.business.recipe.getpage"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAilabsIotBusinessRecipeGetpageAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is OpenAccountId Setter
// 开放账号id
func (r *AlibabaAilabsIotBusinessRecipeGetpageAPIRequest) SetOpenAccountId(_openAccountId string) error {
	r._openAccountId = _openAccountId
	r.Set("open_account_id", _openAccountId)
	return nil
}

// Get OpenAccountId Getter
func (r AlibabaAilabsIotBusinessRecipeGetpageAPIRequest) GetOpenAccountId() string {
	return r._openAccountId
}

// Set is PageNum Setter
// 分页页码
func (r *AlibabaAilabsIotBusinessRecipeGetpageAPIRequest) SetPageNum(_pageNum int64) error {
	r._pageNum = _pageNum
	r.Set("page_num", _pageNum)
	return nil
}

// Get PageNum Getter
func (r AlibabaAilabsIotBusinessRecipeGetpageAPIRequest) GetPageNum() int64 {
	return r._pageNum
}

// Set is PageSize Setter
// 分页大小
func (r *AlibabaAilabsIotBusinessRecipeGetpageAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// Get PageSize Getter
func (r AlibabaAilabsIotBusinessRecipeGetpageAPIRequest) GetPageSize() int64 {
	return r._pageSize
}
