package tblogistics

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAreasGetAPIRequest 查询地址区域 API请求
// taobao.areas.get
//
// 查询标准地址区域代码信息。可以直接参考最新的行政区域代码：
// &lt;a href=&#34;http://www.stats.gov.cn/tjsj/tjbz/tjyqhdmhcxhfdm/&#34; target=&#34;_blank&#34;&gt; http://www.stats.gov.cn/tjsj/tjbz/tjyqhdmhcxhfdm/&lt;/a&gt;
type TaobaoAreasGetAPIRequest struct {
	model.Params
	// 需返回的字段列表.可选值:Area 结构中的所有字段;多个字段之间用","分隔.如:id,type,name,parent_id,zip.
	_fields string
}

// NewTaobaoAreasGetRequest 初始化TaobaoAreasGetAPIRequest对象
func NewTaobaoAreasGetRequest() *TaobaoAreasGetAPIRequest {
	return &TaobaoAreasGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAreasGetAPIRequest) GetApiMethodName() string {
	return "taobao.areas.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAreasGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoAreasGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFields is Fields Setter
// 需返回的字段列表.可选值:Area 结构中的所有字段;多个字段之间用&#34;,&#34;分隔.如:id,type,name,parent_id,zip.
func (r *TaobaoAreasGetAPIRequest) SetFields(_fields string) error {
	r._fields = _fields
	r.Set("fields", _fields)
	return nil
}

// GetFields Fields Getter
func (r TaobaoAreasGetAPIRequest) GetFields() string {
	return r._fields
}
