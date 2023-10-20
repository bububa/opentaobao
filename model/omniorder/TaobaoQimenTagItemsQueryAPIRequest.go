package omniorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoqimentagitemsqueryAPIRequest 打标结果查询-标维度 API请求
// taobao.qimen.tag.items.query
//
// 调用该接口，查询打了某个标的商品列表。说明：该接口调用后，返回值的时间较长，建议不要经常调用。
type TaobaoqimentagitemsqueryAPIRequest struct {
	model.Params
	// 打标值，string（50），TBKU=同步库存标，MDZT=门店自提标，必填
	_tagType string
	// 备注，string（500）
	_remark string
}

// NewTaobaoqimentagitemsqueryRequest 初始化TaobaoqimentagitemsqueryAPIRequest对象
func NewTaobaoqimentagitemsqueryRequest() *TaobaoqimentagitemsqueryAPIRequest {
	return &TaobaoqimentagitemsqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoqimentagitemsqueryAPIRequest) GetApiMethodName() string {
	return "taobao.qimen.tag.items.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoqimentagitemsqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoqimentagitemsqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTagType is TagType Setter
// 打标值，string（50），TBKU=同步库存标，MDZT=门店自提标，必填
func (r *TaobaoqimentagitemsqueryAPIRequest) SetTagType(_tagType string) error {
	r._tagType = _tagType
	r.Set("tag_type", _tagType)
	return nil
}

// GetTagType TagType Getter
func (r TaobaoqimentagitemsqueryAPIRequest) GetTagType() string {
	return r._tagType
}

// SetRemark is Remark Setter
// 备注，string（500）
func (r *TaobaoqimentagitemsqueryAPIRequest) SetRemark(_remark string) error {
	r._remark = _remark
	r.Set("remark", _remark)
	return nil
}

// GetRemark Remark Getter
func (r TaobaoqimentagitemsqueryAPIRequest) GetRemark() string {
	return r._remark
}
