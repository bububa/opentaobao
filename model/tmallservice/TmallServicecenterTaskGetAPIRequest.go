package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallServicecenterTaskGetAPIRequest 服务工单拉取 API请求
// tmall.servicecenter.task.get
//
// 接口供服务供应商通过交易主订单查询其未拉取的任务类工单
type TmallServicecenterTaskGetAPIRequest struct {
	model.Params
	// Taobao主交易订单ID
	_parentBizOrderId int64
}

// NewTmallServicecenterTaskGetRequest 初始化TmallServicecenterTaskGetAPIRequest对象
func NewTmallServicecenterTaskGetRequest() *TmallServicecenterTaskGetAPIRequest {
	return &TmallServicecenterTaskGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallServicecenterTaskGetAPIRequest) GetApiMethodName() string {
	return "tmall.servicecenter.task.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallServicecenterTaskGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ParentBizOrderId Setter
// Taobao主交易订单ID
func (r *TmallServicecenterTaskGetAPIRequest) SetParentBizOrderId(_parentBizOrderId int64) error {
	r._parentBizOrderId = _parentBizOrderId
	r.Set("parent_biz_order_id", _parentBizOrderId)
	return nil
}

// Get ParentBizOrderId Getter
func (r TmallServicecenterTaskGetAPIRequest) GetParentBizOrderId() int64 {
	return r._parentBizOrderId
}
