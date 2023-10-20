package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallservicecentertaskqueryrefundAPIRequest 查询任务类工单是否退款 API请求
// tmall.servicecenter.task.queryrefund
//
// 查询任务类工单是否退款
type TmallservicecentertaskqueryrefundAPIRequest struct {
	model.Params
	// 工单id列表
	_workcardList []string
}

// NewTmallservicecentertaskqueryrefundRequest 初始化TmallservicecentertaskqueryrefundAPIRequest对象
func NewTmallservicecentertaskqueryrefundRequest() *TmallservicecentertaskqueryrefundAPIRequest {
	return &TmallservicecentertaskqueryrefundAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallservicecentertaskqueryrefundAPIRequest) GetApiMethodName() string {
	return "tmall.servicecenter.task.queryrefund"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallservicecentertaskqueryrefundAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallservicecentertaskqueryrefundAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetWorkcardList is WorkcardList Setter
// 工单id列表
func (r *TmallservicecentertaskqueryrefundAPIRequest) SetWorkcardList(_workcardList []string) error {
	r._workcardList = _workcardList
	r.Set("workcard_list", _workcardList)
	return nil
}

// GetWorkcardList WorkcardList Getter
func (r TmallservicecentertaskqueryrefundAPIRequest) GetWorkcardList() []string {
	return r._workcardList
}
