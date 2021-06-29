package jst

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
标签查询订单 API请求
taobao.oc.trades.bytag.get

根据标签查询订单编号
*/
type TaobaoOcTradesBytagGetRequest struct {
    model.Params
    // 分页大小
    pageSize   int64
    // 标签类型，1官方，2自定义
    tagType   int64
    // 当前页
    page   int64
    // 标签名称
    tagName   string
}

// 初始化TaobaoOcTradesBytagGetRequest对象
func NewTaobaoOcTradesBytagGetRequest() *TaobaoOcTradesBytagGetRequest{
    return &TaobaoOcTradesBytagGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoOcTradesBytagGetRequest) GetApiMethodName() string {
    return "taobao.oc.trades.bytag.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoOcTradesBytagGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// PageSize Setter
// 分页大小
func (r *TaobaoOcTradesBytagGetRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoOcTradesBytagGetRequest) GetPageSize() int64 {
    return r.pageSize
}
// TagType Setter
// 标签类型，1官方，2自定义
func (r *TaobaoOcTradesBytagGetRequest) SetTagType(tagType int64) error {
    r.tagType = tagType
    r.Set("tag_type", tagType)
    return nil
}

// TagType Getter
func (r TaobaoOcTradesBytagGetRequest) GetTagType() int64 {
    return r.tagType
}
// Page Setter
// 当前页
func (r *TaobaoOcTradesBytagGetRequest) SetPage(page int64) error {
    r.page = page
    r.Set("page", page)
    return nil
}

// Page Getter
func (r TaobaoOcTradesBytagGetRequest) GetPage() int64 {
    return r.page
}
// TagName Setter
// 标签名称
func (r *TaobaoOcTradesBytagGetRequest) SetTagName(tagName string) error {
    r.tagName = tagName
    r.Set("tag_name", tagName)
    return nil
}

// TagName Getter
func (r TaobaoOcTradesBytagGetRequest) GetTagName() string {
    return r.tagName
}
