package alihealth2

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询渠道商api API请求
alibaba.alihealth.tracecodeseller.channel.search

查询渠道商api
*/
type AlibabaAlihealthTracecodesellerChannelSearchAPIRequest struct {
    model.Params
    // 身份认证
    _skeyCode   string
    // 商家id
    _entInfoId   int64
    // 第几页
    _page   int64
    // 每页几条
    _pageSize   int64
    // 0 出库 2 入库
    _outInType   int64
}

// 初始化AlibabaAlihealthTracecodesellerChannelSearchAPIRequest对象
func NewAlibabaAlihealthTracecodesellerChannelSearchRequest() *AlibabaAlihealthTracecodesellerChannelSearchAPIRequest{
    return &AlibabaAlihealthTracecodesellerChannelSearchAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthTracecodesellerChannelSearchAPIRequest) GetApiMethodName() string {
    return "alibaba.alihealth.tracecodeseller.channel.search"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthTracecodesellerChannelSearchAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SkeyCode Setter
// 身份认证
func (r *AlibabaAlihealthTracecodesellerChannelSearchAPIRequest) SetSkeyCode(_skeyCode string) error {
    r._skeyCode = _skeyCode
    r.Set("skey_code", _skeyCode)
    return nil
}

// SkeyCode Getter
func (r AlibabaAlihealthTracecodesellerChannelSearchAPIRequest) GetSkeyCode() string {
    return r._skeyCode
}
// EntInfoId Setter
// 商家id
func (r *AlibabaAlihealthTracecodesellerChannelSearchAPIRequest) SetEntInfoId(_entInfoId int64) error {
    r._entInfoId = _entInfoId
    r.Set("ent_info_id", _entInfoId)
    return nil
}

// EntInfoId Getter
func (r AlibabaAlihealthTracecodesellerChannelSearchAPIRequest) GetEntInfoId() int64 {
    return r._entInfoId
}
// Page Setter
// 第几页
func (r *AlibabaAlihealthTracecodesellerChannelSearchAPIRequest) SetPage(_page int64) error {
    r._page = _page
    r.Set("page", _page)
    return nil
}

// Page Getter
func (r AlibabaAlihealthTracecodesellerChannelSearchAPIRequest) GetPage() int64 {
    return r._page
}
// PageSize Setter
// 每页几条
func (r *AlibabaAlihealthTracecodesellerChannelSearchAPIRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r AlibabaAlihealthTracecodesellerChannelSearchAPIRequest) GetPageSize() int64 {
    return r._pageSize
}
// OutInType Setter
// 0 出库 2 入库
func (r *AlibabaAlihealthTracecodesellerChannelSearchAPIRequest) SetOutInType(_outInType int64) error {
    r._outInType = _outInType
    r.Set("out_in_type", _outInType)
    return nil
}

// OutInType Getter
func (r AlibabaAlihealthTracecodesellerChannelSearchAPIRequest) GetOutInType() int64 {
    return r._outInType
}
