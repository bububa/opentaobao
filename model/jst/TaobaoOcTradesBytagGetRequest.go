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
type TaobaoOcTradesBytagGetAPIRequest struct {
    model.Params
    // 分页大小
    _pageSize   int64
    // 标签类型，1官方，2自定义
    _tagType   int64
    // 当前页
    _page   int64
    // 标签名称
    _tagName   string
}

// 初始化TaobaoOcTradesBytagGetAPIRequest对象
func NewTaobaoOcTradesBytagGetRequest() *TaobaoOcTradesBytagGetAPIRequest{
    return &TaobaoOcTradesBytagGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoOcTradesBytagGetAPIRequest) GetApiMethodName() string {
    return "taobao.oc.trades.bytag.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoOcTradesBytagGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// PageSize Setter
// 分页大小
func (r *TaobaoOcTradesBytagGetAPIRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoOcTradesBytagGetAPIRequest) GetPageSize() int64 {
    return r._pageSize
}
// TagType Setter
// 标签类型，1官方，2自定义
func (r *TaobaoOcTradesBytagGetAPIRequest) SetTagType(_tagType int64) error {
    r._tagType = _tagType
    r.Set("tag_type", _tagType)
    return nil
}

// TagType Getter
func (r TaobaoOcTradesBytagGetAPIRequest) GetTagType() int64 {
    return r._tagType
}
// Page Setter
// 当前页
func (r *TaobaoOcTradesBytagGetAPIRequest) SetPage(_page int64) error {
    r._page = _page
    r.Set("page", _page)
    return nil
}

// Page Getter
func (r TaobaoOcTradesBytagGetAPIRequest) GetPage() int64 {
    return r._page
}
// TagName Setter
// 标签名称
func (r *TaobaoOcTradesBytagGetAPIRequest) SetTagName(_tagName string) error {
    r._tagName = _tagName
    r.Set("tag_name", _tagName)
    return nil
}

// TagName Getter
func (r TaobaoOcTradesBytagGetAPIRequest) GetTagName() string {
    return r._tagName
}
