package iot

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
分页查询食谱 API请求
alibaba.ailabs.iot.business.recipe.getpage

分页查询食谱数据
*/
type AlibabaAilabsIotBusinessRecipeGetpageRequest struct {
    model.Params
    // 开放账号id
    _openAccountId   string
    // 分页页码
    _pageNum   int64
    // 分页大小
    _pageSize   int64
}

// 初始化AlibabaAilabsIotBusinessRecipeGetpageRequest对象
func NewAlibabaAilabsIotBusinessRecipeGetpageRequest() *AlibabaAilabsIotBusinessRecipeGetpageRequest{
    return &AlibabaAilabsIotBusinessRecipeGetpageRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAilabsIotBusinessRecipeGetpageRequest) GetApiMethodName() string {
    return "alibaba.ailabs.iot.business.recipe.getpage"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAilabsIotBusinessRecipeGetpageRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OpenAccountId Setter
// 开放账号id
func (r *AlibabaAilabsIotBusinessRecipeGetpageRequest) SetOpenAccountId(_openAccountId string) error {
    r._openAccountId = _openAccountId
    r.Set("open_account_id", _openAccountId)
    return nil
}

// OpenAccountId Getter
func (r AlibabaAilabsIotBusinessRecipeGetpageRequest) GetOpenAccountId() string {
    return r._openAccountId
}
// PageNum Setter
// 分页页码
func (r *AlibabaAilabsIotBusinessRecipeGetpageRequest) SetPageNum(_pageNum int64) error {
    r._pageNum = _pageNum
    r.Set("page_num", _pageNum)
    return nil
}

// PageNum Getter
func (r AlibabaAilabsIotBusinessRecipeGetpageRequest) GetPageNum() int64 {
    return r._pageNum
}
// PageSize Setter
// 分页大小
func (r *AlibabaAilabsIotBusinessRecipeGetpageRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r AlibabaAilabsIotBusinessRecipeGetpageRequest) GetPageSize() int64 {
    return r._pageSize
}
