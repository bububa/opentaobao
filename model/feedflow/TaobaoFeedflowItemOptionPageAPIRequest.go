package feedflow

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFeedflowItemOptionPageAPIRequest 分页查询定向标签列表 API请求
// taobao.feedflow.item.option.page
//
// 分页查询定向标签列表
type TaobaoFeedflowItemOptionPageAPIRequest struct {
	model.Params
	// 标签查询条件
	_labelQuery *LabelQueryDto
}

// NewTaobaoFeedflowItemOptionPageRequest 初始化TaobaoFeedflowItemOptionPageAPIRequest对象
func NewTaobaoFeedflowItemOptionPageRequest() *TaobaoFeedflowItemOptionPageAPIRequest {
	return &TaobaoFeedflowItemOptionPageAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoFeedflowItemOptionPageAPIRequest) GetApiMethodName() string {
	return "taobao.feedflow.item.option.page"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoFeedflowItemOptionPageAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetLabelQuery is LabelQuery Setter
// 标签查询条件
func (r *TaobaoFeedflowItemOptionPageAPIRequest) SetLabelQuery(_labelQuery *LabelQueryDto) error {
	r._labelQuery = _labelQuery
	r.Set("label_query", _labelQuery)
	return nil
}

// GetLabelQuery LabelQuery Getter
func (r TaobaoFeedflowItemOptionPageAPIRequest) GetLabelQuery() *LabelQueryDto {
	return r._labelQuery
}
