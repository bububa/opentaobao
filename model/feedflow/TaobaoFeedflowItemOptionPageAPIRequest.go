package feedflow

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoFeedflowItemOptionPageAPIRequest) Reset() {
	r._labelQuery = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoFeedflowItemOptionPageAPIRequest) GetApiMethodName() string {
	return "taobao.feedflow.item.option.page"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoFeedflowItemOptionPageAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoFeedflowItemOptionPageAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolTaobaoFeedflowItemOptionPageAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoFeedflowItemOptionPageRequest()
	},
}

// GetTaobaoFeedflowItemOptionPageRequest 从 sync.Pool 获取 TaobaoFeedflowItemOptionPageAPIRequest
func GetTaobaoFeedflowItemOptionPageAPIRequest() *TaobaoFeedflowItemOptionPageAPIRequest {
	return poolTaobaoFeedflowItemOptionPageAPIRequest.Get().(*TaobaoFeedflowItemOptionPageAPIRequest)
}

// ReleaseTaobaoFeedflowItemOptionPageAPIRequest 将 TaobaoFeedflowItemOptionPageAPIRequest 放入 sync.Pool
func ReleaseTaobaoFeedflowItemOptionPageAPIRequest(v *TaobaoFeedflowItemOptionPageAPIRequest) {
	v.Reset()
	poolTaobaoFeedflowItemOptionPageAPIRequest.Put(v)
}
