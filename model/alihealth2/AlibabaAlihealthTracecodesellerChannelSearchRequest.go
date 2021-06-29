package alihealth2

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询渠道商api APIRequest
alibaba.alihealth.tracecodeseller.channel.search

查询渠道商api
*/
type AlibabaAlihealthTracecodesellerChannelSearchRequest struct {
    model.Params

    // 身份认证
    skeyCode   string 

    // 商家id
    entInfoId   int64 

    // 第几页
    page   int64 

    // 每页几条
    pageSize   int64 

    // 0 出库 2 入库
    outInType   int64 

}

func NewAlibabaAlihealthTracecodesellerChannelSearchRequest() *AlibabaAlihealthTracecodesellerChannelSearchRequest{
    return &AlibabaAlihealthTracecodesellerChannelSearchRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthTracecodesellerChannelSearchRequest) GetApiMethodName() string {
    return "alibaba.alihealth.tracecodeseller.channel.search"
}

func (r AlibabaAlihealthTracecodesellerChannelSearchRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthTracecodesellerChannelSearchRequest) SetSkeyCode(skeyCode string) error {
    r.skeyCode = skeyCode
    r.Set("skey_code", skeyCode)
    return nil
}

func (r AlibabaAlihealthTracecodesellerChannelSearchRequest) GetSkeyCode() string {
    return r.skeyCode
}

func (r *AlibabaAlihealthTracecodesellerChannelSearchRequest) SetEntInfoId(entInfoId int64) error {
    r.entInfoId = entInfoId
    r.Set("ent_info_id", entInfoId)
    return nil
}

func (r AlibabaAlihealthTracecodesellerChannelSearchRequest) GetEntInfoId() int64 {
    return r.entInfoId
}

func (r *AlibabaAlihealthTracecodesellerChannelSearchRequest) SetPage(page int64) error {
    r.page = page
    r.Set("page", page)
    return nil
}

func (r AlibabaAlihealthTracecodesellerChannelSearchRequest) GetPage() int64 {
    return r.page
}

func (r *AlibabaAlihealthTracecodesellerChannelSearchRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r AlibabaAlihealthTracecodesellerChannelSearchRequest) GetPageSize() int64 {
    return r.pageSize
}

func (r *AlibabaAlihealthTracecodesellerChannelSearchRequest) SetOutInType(outInType int64) error {
    r.outInType = outInType
    r.Set("out_in_type", outInType)
    return nil
}

func (r AlibabaAlihealthTracecodesellerChannelSearchRequest) GetOutInType() int64 {
    return r.outInType
}

