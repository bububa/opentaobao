package feedflow

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaofeedflowitemoptionpageAPIRequest 分页查询定向标签列表 API请求
// taobao.feedflow.item.option.page
//
// 分页查询定向标签列表
type TaobaofeedflowitemoptionpageAPIRequest struct {
	model.Params
	// 标签查询条件
	_labelQuery *LabelQueryDto
}

// NewTaobaofeedflowitemoptionpageRequest 初始化TaobaofeedflowitemoptionpageAPIRequest对象
func NewTaobaofeedflowitemoptionpageRequest() *TaobaofeedflowitemoptionpageAPIRequest {
	return &TaobaofeedflowitemoptionpageAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaofeedflowitemoptionpageAPIRequest) GetApiMethodName() string {
	return "taobao.feedflow.item.option.page"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaofeedflowitemoptionpageAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaofeedflowitemoptionpageAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetLabelQuery is LabelQuery Setter
// 标签查询条件
func (r *TaobaofeedflowitemoptionpageAPIRequest) SetLabelQuery(_labelQuery *LabelQueryDto) error {
	r._labelQuery = _labelQuery
	r.Set("label_query", _labelQuery)
	return nil
}

// GetLabelQuery LabelQuery Getter
func (r TaobaofeedflowitemoptionpageAPIRequest) GetLabelQuery() *LabelQueryDto {
	return r._labelQuery
}
