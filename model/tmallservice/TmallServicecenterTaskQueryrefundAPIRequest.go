package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallServicecenterTaskQueryrefundAPIRequest
查询任务类工单是否退款 API请求
tmall.servicecenter.task.queryrefund

查询任务类工单是否退款 */
type TmallServicecenterTaskQueryrefundAPIRequest struct {
	model.Params
	// 工单id列表
	_workcardList []int64
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
func (r TmallServicecenterTaskQueryrefundAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is WorkcardList Setter
// 工单id列表
func (r *TmallServicecenterTaskQueryrefundAPIRequest) SetWorkcardList(_workcardList []int64) error {
	r._workcardList = _workcardList
	r.Set("workcard_list", _workcardList)
	return nil
}

// Get WorkcardList Getter
func (r TmallServicecenterTaskQueryrefundAPIRequest) GetWorkcardList() []int64 {
	return r._workcardList
}
