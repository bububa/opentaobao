package logistic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAreasGetAPIRequest 查询地址区域 API请求
// taobao.areas.get
//
// 查询标准地址区域代码信息。可以直接参考最新的行政区域代码：
// <a href="http://www.stats.gov.cn/tjsj/tjbz/tjyqhdmhcxhfdm/" target="_blank"> http://www.stats.gov.cn/tjsj/tjbz/tjyqhdmhcxhfdm/</a>
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
func (r TaobaoAreasGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetFields is Fields Setter
// 需返回的字段列表.可选值:Area 结构中的所有字段;多个字段之间用","分隔.如:id,type,name,parent_id,zip.
func (r *TaobaoAreasGetAPIRequest) SetFields(_fields string) error {
	r._fields = _fields
	r.Set("fields", _fields)
	return nil
}

// GetFields Fields Getter
func (r TaobaoAreasGetAPIRequest) GetFields() string {
	return r._fields
}
