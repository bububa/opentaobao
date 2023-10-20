package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallServicecenterTaskQueryrefundAPIRequest 查询任务类工单是否退款 API请求
// tmall.servicecenter.task.queryrefund
//
// 查询任务类工单是否退款
type TmallServicecenterTaskQueryrefundAPIRequest struct {
	model.Params
	// 工单id列表
	_workcardList []string
}

// NewTmallServicecenterTaskQueryrefundRequest 初始化TmallServicecenterTaskQueryrefundAPIRequest对象
func NewTmallServicecenterTaskQueryrefundRequest() *TmallServicecenterTaskQueryrefundAPIRequest {
	return &TmallServicecenterTaskQueryrefundAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallServicecenterTaskQueryrefundAPIRequest) GetApiMethodName() string {
	return "tmall.servicecenter.task.queryrefund"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallServicecenterTaskQueryrefundAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallServicecenterTaskQueryrefundAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetWorkcardList is WorkcardList Setter
// 工单id列表
func (r *TmallServicecenterTaskQueryrefundAPIRequest) SetWorkcardList(_workcardList []string) error {
	r._workcardList = _workcardList
	r.Set("workcard_list", _workcardList)
	return nil
}

// GetWorkcardList WorkcardList Getter
func (r TmallServicecenterTaskQueryrefundAPIRequest) GetWorkcardList() []string {
	return r._workcardList
}
