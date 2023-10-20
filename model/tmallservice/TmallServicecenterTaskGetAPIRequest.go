package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallservicecentertaskgetAPIRequest 服务工单拉取 API请求
// tmall.servicecenter.task.get
//
// 接口供服务供应商通过交易主订单查询其未拉取的任务类工单
type TmallservicecentertaskgetAPIRequest struct {
	model.Params
	// Taobao主交易订单ID
	_parentBizOrderId int64
}

// NewTmallservicecentertaskgetRequest 初始化TmallservicecentertaskgetAPIRequest对象
func NewTmallservicecentertaskgetRequest() *TmallservicecentertaskgetAPIRequest {
	return &TmallservicecentertaskgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallservicecentertaskgetAPIRequest) GetApiMethodName() string {
	return "tmall.servicecenter.task.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallservicecentertaskgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallservicecentertaskgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParentBizOrderId is ParentBizOrderId Setter
// Taobao主交易订单ID
func (r *TmallservicecentertaskgetAPIRequest) SetParentBizOrderId(_parentBizOrderId int64) error {
	r._parentBizOrderId = _parentBizOrderId
	r.Set("parent_biz_order_id", _parentBizOrderId)
	return nil
}

// GetParentBizOrderId ParentBizOrderId Getter
func (r TmallservicecentertaskgetAPIRequest) GetParentBizOrderId() int64 {
	return r._parentBizOrderId
}
