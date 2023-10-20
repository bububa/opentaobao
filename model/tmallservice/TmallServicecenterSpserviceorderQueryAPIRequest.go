package tmallservice

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallServicecenterSpserviceorderQueryAPIRequest) Reset() {
	r._parentBizOrderId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallServicecenterSpserviceorderQueryAPIRequest) GetApiMethodName() string {
	return "tmall.servicecenter.spserviceorder.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallServicecenterSpserviceorderQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallServicecenterSpserviceorderQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParentBizOrderId is ParentBizOrderId Setter
// 交易主订单id
func (r *TmallServicecenterSpserviceorderQueryAPIRequest) SetParentBizOrderId(_parentBizOrderId int64) error {
	r._parentBizOrderId = _parentBizOrderId
	r.Set("parent_biz_order_id", _parentBizOrderId)
	return nil
}

// GetParentBizOrderId ParentBizOrderId Getter
func (r TmallServicecenterSpserviceorderQueryAPIRequest) GetParentBizOrderId() int64 {
	return r._parentBizOrderId
}

var poolTmallServicecenterSpserviceorderQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallServicecenterSpserviceorderQueryRequest()
	},
}

// GetTmallServicecenterSpserviceorderQueryRequest 从 sync.Pool 获取 TmallServicecenterSpserviceorderQueryAPIRequest
func GetTmallServicecenterSpserviceorderQueryAPIRequest() *TmallServicecenterSpserviceorderQueryAPIRequest {
	return poolTmallServicecenterSpserviceorderQueryAPIRequest.Get().(*TmallServicecenterSpserviceorderQueryAPIRequest)
}

// ReleaseTmallServicecenterSpserviceorderQueryAPIRequest 将 TmallServicecenterSpserviceorderQueryAPIRequest 放入 sync.Pool
func ReleaseTmallServicecenterSpserviceorderQueryAPIRequest(v *TmallServicecenterSpserviceorderQueryAPIRequest) {
	v.Reset()
	poolTmallServicecenterSpserviceorderQueryAPIRequest.Put(v)
}
