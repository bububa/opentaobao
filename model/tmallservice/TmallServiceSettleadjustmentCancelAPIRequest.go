package tmallservice

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallServiceSettleadjustmentCancelAPIRequest 取消结算调整单 API请求
// tmall.service.settleadjustment.cancel
//
// 提供给服务商在对取消已经发起的结算调整单。
// 通过说明调整单ID进行结算调整单取消。
type TmallServiceSettleadjustmentCancelAPIRequest struct {
	model.Params
	// 取消原因说明
	_comments string
	// 结算调整单ID
	_id int64
}

// NewTmallServiceSettleadjustmentCancelRequest 初始化TmallServiceSettleadjustmentCancelAPIRequest对象
func NewTmallServiceSettleadjustmentCancelRequest() *TmallServiceSettleadjustmentCancelAPIRequest {
	return &TmallServiceSettleadjustmentCancelAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallServiceSettleadjustmentCancelAPIRequest) Reset() {
	r._comments = ""
	r._id = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallServiceSettleadjustmentCancelAPIRequest) GetApiMethodName() string {
	return "tmall.service.settleadjustment.cancel"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallServiceSettleadjustmentCancelAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallServiceSettleadjustmentCancelAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetComments is Comments Setter
// 取消原因说明
func (r *TmallServiceSettleadjustmentCancelAPIRequest) SetComments(_comments string) error {
	r._comments = _comments
	r.Set("comments", _comments)
	return nil
}

// GetComments Comments Getter
func (r TmallServiceSettleadjustmentCancelAPIRequest) GetComments() string {
	return r._comments
}

// SetId is Id Setter
// 结算调整单ID
func (r *TmallServiceSettleadjustmentCancelAPIRequest) SetId(_id int64) error {
	r._id = _id
	r.Set("id", _id)
	return nil
}

// GetId Id Getter
func (r TmallServiceSettleadjustmentCancelAPIRequest) GetId() int64 {
	return r._id
}

var poolTmallServiceSettleadjustmentCancelAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallServiceSettleadjustmentCancelRequest()
	},
}

// GetTmallServiceSettleadjustmentCancelRequest 从 sync.Pool 获取 TmallServiceSettleadjustmentCancelAPIRequest
func GetTmallServiceSettleadjustmentCancelAPIRequest() *TmallServiceSettleadjustmentCancelAPIRequest {
	return poolTmallServiceSettleadjustmentCancelAPIRequest.Get().(*TmallServiceSettleadjustmentCancelAPIRequest)
}

// ReleaseTmallServiceSettleadjustmentCancelAPIRequest 将 TmallServiceSettleadjustmentCancelAPIRequest 放入 sync.Pool
func ReleaseTmallServiceSettleadjustmentCancelAPIRequest(v *TmallServiceSettleadjustmentCancelAPIRequest) {
	v.Reset()
	poolTmallServiceSettleadjustmentCancelAPIRequest.Put(v)
}
