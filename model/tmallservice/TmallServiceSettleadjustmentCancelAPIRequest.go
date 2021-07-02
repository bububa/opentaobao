package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallServiceSettleadjustmentCancelAPIRequest 取消结算调整单 API请求
// tmall.service.settleadjustment.cancel
//
// 提供给服务商在对取消已经发起的结算调整单。
// 通过说明调整单ID进行结算调整单取消。
type TmallServiceSettleadjustmentCancelAPIRequest struct {
	model.Params
	// 结算调整单ID
	_id int64
	// 取消原因说明
	_comments string
}

// NewTmallServiceSettleadjustmentCancelRequest 初始化TmallServiceSettleadjustmentCancelAPIRequest对象
func NewTmallServiceSettleadjustmentCancelRequest() *TmallServiceSettleadjustmentCancelAPIRequest {
	return &TmallServiceSettleadjustmentCancelAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallServiceSettleadjustmentCancelAPIRequest) GetApiMethodName() string {
	return "tmall.service.settleadjustment.cancel"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallServiceSettleadjustmentCancelAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Id Setter
// 结算调整单ID
func (r *TmallServiceSettleadjustmentCancelAPIRequest) SetId(_id int64) error {
	r._id = _id
	r.Set("id", _id)
	return nil
}

// Get Id Getter
func (r TmallServiceSettleadjustmentCancelAPIRequest) GetId() int64 {
	return r._id
}

// Set is Comments Setter
// 取消原因说明
func (r *TmallServiceSettleadjustmentCancelAPIRequest) SetComments(_comments string) error {
	r._comments = _comments
	r.Set("comments", _comments)
	return nil
}

// Get Comments Getter
func (r TmallServiceSettleadjustmentCancelAPIRequest) GetComments() string {
	return r._comments
}
