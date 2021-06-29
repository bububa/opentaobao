package jst

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
订单打标或者订单标签更新 API请求
taobao.oc.tradetag.attach

对订单添加标签和更新标签。标签分为官方标签和自定义标签。<br/>官方标签目前有:赠品,电子发票,收货地址变更,预售。具体格式说明请看http://open.taobao.com/doc/detail.htm?id=102731<br/>自定义标签有2个通用属性:<br/>    `show_str:给消费者显示的字符串（如果可以显示的话）<br/>    `pic_urls:图片url,地址必须是图片空间的url,最多5张
*/
type TaobaoOcTradetagAttachRequest struct {
    model.Params
    // 标签类型       1：官方标签      2：自定义标签
    tagType   int64
    // 订单id
    tid   int64
    // 标签名称
    tagName   string
    // 标签值，json格式
    tagValue   string
    // 该标签在消费者端是否显示,0:不显示,1：显示
    visible   int64
}

// 初始化TaobaoOcTradetagAttachRequest对象
func NewTaobaoOcTradetagAttachRequest() *TaobaoOcTradetagAttachRequest{
    return &TaobaoOcTradetagAttachRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoOcTradetagAttachRequest) GetApiMethodName() string {
    return "taobao.oc.tradetag.attach"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoOcTradetagAttachRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TagType Setter
// 标签类型       1：官方标签      2：自定义标签
func (r *TaobaoOcTradetagAttachRequest) SetTagType(tagType int64) error {
    r.tagType = tagType
    r.Set("tag_type", tagType)
    return nil
}

// TagType Getter
func (r TaobaoOcTradetagAttachRequest) GetTagType() int64 {
    return r.tagType
}
// Tid Setter
// 订单id
func (r *TaobaoOcTradetagAttachRequest) SetTid(tid int64) error {
    r.tid = tid
    r.Set("tid", tid)
    return nil
}

// Tid Getter
func (r TaobaoOcTradetagAttachRequest) GetTid() int64 {
    return r.tid
}
// TagName Setter
// 标签名称
func (r *TaobaoOcTradetagAttachRequest) SetTagName(tagName string) error {
    r.tagName = tagName
    r.Set("tag_name", tagName)
    return nil
}

// TagName Getter
func (r TaobaoOcTradetagAttachRequest) GetTagName() string {
    return r.tagName
}
// TagValue Setter
// 标签值，json格式
func (r *TaobaoOcTradetagAttachRequest) SetTagValue(tagValue string) error {
    r.tagValue = tagValue
    r.Set("tag_value", tagValue)
    return nil
}

// TagValue Getter
func (r TaobaoOcTradetagAttachRequest) GetTagValue() string {
    return r.tagValue
}
// Visible Setter
// 该标签在消费者端是否显示,0:不显示,1：显示
func (r *TaobaoOcTradetagAttachRequest) SetVisible(visible int64) error {
    r.visible = visible
    r.Set("visible", visible)
    return nil
}

// Visible Getter
func (r TaobaoOcTradetagAttachRequest) GetVisible() int64 {
    return r.visible
}
