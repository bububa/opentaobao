package openim

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoopenimsnfilterwordsetfilterAPIRequest 关键词过滤 API请求
// taobao.openim.snfilterword.setfilter
//
// 设置openim关键词过滤
type TaobaoopenimsnfilterwordsetfilterAPIRequest struct {
	model.Params
	// 上传者身份信息，区分不同上传者;只是记录，没有身份校验功能
	_creator string
	// 需要过滤的关键词
	_filterword string
	// 过滤原因描述
	_desc string
}

// NewTaobaoopenimsnfilterwordsetfilterRequest 初始化TaobaoopenimsnfilterwordsetfilterAPIRequest对象
func NewTaobaoopenimsnfilterwordsetfilterRequest() *TaobaoopenimsnfilterwordsetfilterAPIRequest {
	return &TaobaoopenimsnfilterwordsetfilterAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoopenimsnfilterwordsetfilterAPIRequest) GetApiMethodName() string {
	return "taobao.openim.snfilterword.setfilter"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoopenimsnfilterwordsetfilterAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoopenimsnfilterwordsetfilterAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCreator is Creator Setter
// 上传者身份信息，区分不同上传者;只是记录，没有身份校验功能
func (r *TaobaoopenimsnfilterwordsetfilterAPIRequest) SetCreator(_creator string) error {
	r._creator = _creator
	r.Set("creator", _creator)
	return nil
}

// GetCreator Creator Getter
func (r TaobaoopenimsnfilterwordsetfilterAPIRequest) GetCreator() string {
	return r._creator
}

// SetFilterword is Filterword Setter
// 需要过滤的关键词
func (r *TaobaoopenimsnfilterwordsetfilterAPIRequest) SetFilterword(_filterword string) error {
	r._filterword = _filterword
	r.Set("filterword", _filterword)
	return nil
}

// GetFilterword Filterword Getter
func (r TaobaoopenimsnfilterwordsetfilterAPIRequest) GetFilterword() string {
	return r._filterword
}

// SetDesc is Desc Setter
// 过滤原因描述
func (r *TaobaoopenimsnfilterwordsetfilterAPIRequest) SetDesc(_desc string) error {
	r._desc = _desc
	r.Set("desc", _desc)
	return nil
}

// GetDesc Desc Getter
func (r TaobaoopenimsnfilterwordsetfilterAPIRequest) GetDesc() string {
	return r._desc
}
