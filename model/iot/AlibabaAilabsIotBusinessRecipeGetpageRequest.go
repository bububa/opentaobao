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
    openAccountId   string
    // 分页页码
    pageNum   int64
    // 分页大小
    pageSize   int64
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
func (r *AlibabaAilabsIotBusinessRecipeGetpageRequest) SetOpenAccountId(openAccountId string) error {
    r.openAccountId = openAccountId
    r.Set("open_account_id", openAccountId)
    return nil
}

// OpenAccountId Getter
func (r AlibabaAilabsIotBusinessRecipeGetpageRequest) GetOpenAccountId() string {
    return r.openAccountId
}
// PageNum Setter
// 分页页码
func (r *AlibabaAilabsIotBusinessRecipeGetpageRequest) SetPageNum(pageNum int64) error {
    r.pageNum = pageNum
    r.Set("page_num", pageNum)
    return nil
}

// PageNum Getter
func (r AlibabaAilabsIotBusinessRecipeGetpageRequest) GetPageNum() int64 {
    return r.pageNum
}
// PageSize Setter
// 分页大小
func (r *AlibabaAilabsIotBusinessRecipeGetpageRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

// PageSize Getter
func (r AlibabaAilabsIotBusinessRecipeGetpageRequest) GetPageSize() int64 {
    return r.pageSize
}
