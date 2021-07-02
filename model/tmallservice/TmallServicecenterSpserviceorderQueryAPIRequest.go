package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallServicecenterSpserviceorderQueryAPIRequest 服务单列表查询 API请求
// tmall.servicecenter.spserviceorder.query
//
// 查询服务单列表
type TmallServicecenterSpserviceorderQueryAPIRequest struct {
	model.Params
	// 交易主订单id
	_parentBizOrderId int64
}

// NewTmallServicecenterSpserviceorderQueryRequest 初始化TmallServicecenterSpserviceorderQueryAPIRequest对象
func NewTmallServicecenterSpserviceorderQueryRequest() *TmallServicecenterSpserviceorderQueryAPIRequest {
	return &TmallServicecenterSpserviceorderQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallServicecenterSpserviceorderQueryAPIRequest) GetApiMethodName() string {
	return "tmall.servicecenter.spserviceorder.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallServicecenterSpserviceorderQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ParentBizOrderId Setter
// 交易主订单id
func (r *TmallServicecenterSpserviceorderQueryAPIRequest) SetParentBizOrderId(_parentBizOrderId int64) error {
	r._parentBizOrderId = _parentBizOrderId
	r.Set("parent_biz_order_id", _parentBizOrderId)
	return nil
}

// Get ParentBizOrderId Getter
func (r TmallServicecenterSpserviceorderQueryAPIRequest) GetParentBizOrderId() int64 {
	return r._parentBizOrderId
}
