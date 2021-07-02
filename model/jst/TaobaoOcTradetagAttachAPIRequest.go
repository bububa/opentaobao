package jst

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOcTradetagAttachAPIRequest 订单打标或者订单标签更新 API请求
// taobao.oc.tradetag.attach
//
// 对订单添加标签和更新标签。标签分为官方标签和自定义标签。<br/>官方标签目前有:赠品,电子发票,收货地址变更,预售。具体格式说明请看http://open.taobao.com/doc/detail.htm?id=102731<br/>自定义标签有2个通用属性:<br/>    `show_str:给消费者显示的字符串（如果可以显示的话）<br/>    `pic_urls:图片url,地址必须是图片空间的url,最多5张
type TaobaoOcTradetagAttachAPIRequest struct {
	model.Params
	// 标签类型       1：官方标签      2：自定义标签
	_tagType int64
	// 订单id
	_tid int64
	// 标签名称
	_tagName string
	// 标签值，json格式
	_tagValue string
	// 该标签在消费者端是否显示,0:不显示,1：显示
	_visible int64
}

// NewTaobaoOcTradetagAttachRequest 初始化TaobaoOcTradetagAttachAPIRequest对象
func NewTaobaoOcTradetagAttachRequest() *TaobaoOcTradetagAttachAPIRequest {
	return &TaobaoOcTradetagAttachAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOcTradetagAttachAPIRequest) GetApiMethodName() string {
	return "taobao.oc.tradetag.attach"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOcTradetagAttachAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetTagType is TagType Setter
// 标签类型       1：官方标签      2：自定义标签
func (r *TaobaoOcTradetagAttachAPIRequest) SetTagType(_tagType int64) error {
	r._tagType = _tagType
	r.Set("tag_type", _tagType)
	return nil
}

// GetTagType TagType Getter
func (r TaobaoOcTradetagAttachAPIRequest) GetTagType() int64 {
	return r._tagType
}

// SetTid is Tid Setter
// 订单id
func (r *TaobaoOcTradetagAttachAPIRequest) SetTid(_tid int64) error {
	r._tid = _tid
	r.Set("tid", _tid)
	return nil
}

// GetTid Tid Getter
func (r TaobaoOcTradetagAttachAPIRequest) GetTid() int64 {
	return r._tid
}

// SetTagName is TagName Setter
// 标签名称
func (r *TaobaoOcTradetagAttachAPIRequest) SetTagName(_tagName string) error {
	r._tagName = _tagName
	r.Set("tag_name", _tagName)
	return nil
}

// GetTagName TagName Getter
func (r TaobaoOcTradetagAttachAPIRequest) GetTagName() string {
	return r._tagName
}

// SetTagValue is TagValue Setter
// 标签值，json格式
func (r *TaobaoOcTradetagAttachAPIRequest) SetTagValue(_tagValue string) error {
	r._tagValue = _tagValue
	r.Set("tag_value", _tagValue)
	return nil
}

// GetTagValue TagValue Getter
func (r TaobaoOcTradetagAttachAPIRequest) GetTagValue() string {
	return r._tagValue
}

// SetVisible is Visible Setter
// 该标签在消费者端是否显示,0:不显示,1：显示
func (r *TaobaoOcTradetagAttachAPIRequest) SetVisible(_visible int64) error {
	r._visible = _visible
	r.Set("visible", _visible)
	return nil
}

// GetVisible Visible Getter
func (r TaobaoOcTradetagAttachAPIRequest) GetVisible() int64 {
	return r._visible
}
