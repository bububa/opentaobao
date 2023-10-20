package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallservicesettleadjustmentcancelAPIRequest 取消结算调整单 API请求
// tmall.service.settleadjustment.cancel
//
// 提供给服务商在对取消已经发起的结算调整单。
// 通过说明调整单ID进行结算调整单取消。
type TmallservicesettleadjustmentcancelAPIRequest struct {
	model.Params
	// 取消原因说明
	_comments string
	// 结算调整单ID
	_id int64
}

// NewTmallservicesettleadjustmentcancelRequest 初始化TmallservicesettleadjustmentcancelAPIRequest对象
func NewTmallservicesettleadjustmentcancelRequest() *TmallservicesettleadjustmentcancelAPIRequest {
	return &TmallservicesettleadjustmentcancelAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallservicesettleadjustmentcancelAPIRequest) GetApiMethodName() string {
	return "tmall.service.settleadjustment.cancel"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallservicesettleadjustmentcancelAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallservicesettleadjustmentcancelAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetComments is Comments Setter
// 取消原因说明
func (r *TmallservicesettleadjustmentcancelAPIRequest) SetComments(_comments string) error {
	r._comments = _comments
	r.Set("comments", _comments)
	return nil
}

// GetComments Comments Getter
func (r TmallservicesettleadjustmentcancelAPIRequest) GetComments() string {
	return r._comments
}

// SetId is Id Setter
// 结算调整单ID
func (r *TmallservicesettleadjustmentcancelAPIRequest) SetId(_id int64) error {
	r._id = _id
	r.Set("id", _id)
	return nil
}

// GetId Id Getter
func (r TmallservicesettleadjustmentcancelAPIRequest) GetId() int64 {
	return r._id
}
