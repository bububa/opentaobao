package openim

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOpenimSnfilterwordSetfilterAPIRequest
关键词过滤 API请求
taobao.openim.snfilterword.setfilter

设置openim关键词过滤 */
type TaobaoOpenimSnfilterwordSetfilterAPIRequest struct {
	model.Params
	// 上传者身份信息，区分不同上传者;只是记录，没有身份校验功能
	_creator string
	// 需要过滤的关键词
	_filterword string
	// 过滤原因描述
	_desc string
}

// NewTaobaoOpenimSnfilterwordSetfilterRequest 初始化TaobaoOpenimSnfilterwordSetfilterAPIRequest对象
func NewTaobaoOpenimSnfilterwordSetfilterRequest() *TaobaoOpenimSnfilterwordSetfilterAPIRequest {
	return &TaobaoOpenimSnfilterwordSetfilterAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOpenimSnfilterwordSetfilterAPIRequest) GetApiMethodName() string {
	return "taobao.openim.snfilterword.setfilter"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOpenimSnfilterwordSetfilterAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Creator Setter
// 上传者身份信息，区分不同上传者;只是记录，没有身份校验功能
func (r *TaobaoOpenimSnfilterwordSetfilterAPIRequest) SetCreator(_creator string) error {
	r._creator = _creator
	r.Set("creator", _creator)
	return nil
}

// Get Creator Getter
func (r TaobaoOpenimSnfilterwordSetfilterAPIRequest) GetCreator() string {
	return r._creator
}

// Set is Filterword Setter
// 需要过滤的关键词
func (r *TaobaoOpenimSnfilterwordSetfilterAPIRequest) SetFilterword(_filterword string) error {
	r._filterword = _filterword
	r.Set("filterword", _filterword)
	return nil
}

// Get Filterword Getter
func (r TaobaoOpenimSnfilterwordSetfilterAPIRequest) GetFilterword() string {
	return r._filterword
}

// Set is Desc Setter
// 过滤原因描述
func (r *TaobaoOpenimSnfilterwordSetfilterAPIRequest) SetDesc(_desc string) error {
	r._desc = _desc
	r.Set("desc", _desc)
	return nil
}

// Get Desc Getter
func (r TaobaoOpenimSnfilterwordSetfilterAPIRequest) GetDesc() string {
	return r._desc
}
