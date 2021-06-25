package iot

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
分页查询食谱 APIRequest
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

func NewAlibabaAilabsIotBusinessRecipeGetpageRequest() *AlibabaAilabsIotBusinessRecipeGetpageRequest{
    return &AlibabaAilabsIotBusinessRecipeGetpageRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAilabsIotBusinessRecipeGetpageRequest) GetApiMethodName() string {
    return "alibaba.ailabs.iot.business.recipe.getpage"
}

func (r AlibabaAilabsIotBusinessRecipeGetpageRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAilabsIotBusinessRecipeGetpageRequest) SetOpenAccountId(openAccountId string) error {
    r.openAccountId = openAccountId
    r.Set("open_account_id", openAccountId)
    return nil
}

func (r AlibabaAilabsIotBusinessRecipeGetpageRequest) GetOpenAccountId() string {
    return r.openAccountId
}

func (r *AlibabaAilabsIotBusinessRecipeGetpageRequest) SetPageNum(pageNum int64) error {
    r.pageNum = pageNum
    r.Set("page_num", pageNum)
    return nil
}

func (r AlibabaAilabsIotBusinessRecipeGetpageRequest) GetPageNum() int64 {
    return r.pageNum
}

func (r *AlibabaAilabsIotBusinessRecipeGetpageRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r AlibabaAilabsIotBusinessRecipeGetpageRequest) GetPageSize() int64 {
    return r.pageSize
}

