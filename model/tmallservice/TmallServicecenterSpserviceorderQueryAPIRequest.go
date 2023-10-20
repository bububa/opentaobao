package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallservicecenterspserviceorderqueryAPIRequest 服务单列表查询 API请求
// tmall.servicecenter.spserviceorder.query
//
// 查询服务单列表
type TmallservicecenterspserviceorderqueryAPIRequest struct {
	model.Params
	// 交易主订单id
	_parentBizOrderId int64
}

// NewTmallservicecenterspserviceorderqueryRequest 初始化TmallservicecenterspserviceorderqueryAPIRequest对象
func NewTmallservicecenterspserviceorderqueryRequest() *TmallservicecenterspserviceorderqueryAPIRequest {
	return &TmallservicecenterspserviceorderqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallservicecenterspserviceorderqueryAPIRequest) GetApiMethodName() string {
	return "tmall.servicecenter.spserviceorder.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallservicecenterspserviceorderqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallservicecenterspserviceorderqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParentBizOrderId is ParentBizOrderId Setter
// 交易主订单id
func (r *TmallservicecenterspserviceorderqueryAPIRequest) SetParentBizOrderId(_parentBizOrderId int64) error {
	r._parentBizOrderId = _parentBizOrderId
	r.Set("parent_biz_order_id", _parentBizOrderId)
	return nil
}

// GetParentBizOrderId ParentBizOrderId Getter
func (r TmallservicecenterspserviceorderqueryAPIRequest) GetParentBizOrderId() int64 {
	return r._parentBizOrderId
}
