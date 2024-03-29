package tmallservice

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallServiceSettleadjustmentGetAPIRequest 查询结算调整单单条记录 API请求
// tmall.service.settleadjustment.get
//
// 提供给服务商通过结算调整单id获取结算调整单信息
type TmallServiceSettleadjustmentGetAPIRequest struct {
	model.Params
	// 结算调整单ID
	_id int64
}

// NewTmallServiceSettleadjustmentGetRequest 初始化TmallServiceSettleadjustmentGetAPIRequest对象
func NewTmallServiceSettleadjustmentGetRequest() *TmallServiceSettleadjustmentGetAPIRequest {
	return &TmallServiceSettleadjustmentGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallServiceSettleadjustmentGetAPIRequest) Reset() {
	r._id = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallServiceSettleadjustmentGetAPIRequest) GetApiMethodName() string {
	return "tmall.service.settleadjustment.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallServiceSettleadjustmentGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallServiceSettleadjustmentGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetId is Id Setter
// 结算调整单ID
func (r *TmallServiceSettleadjustmentGetAPIRequest) SetId(_id int64) error {
	r._id = _id
	r.Set("id", _id)
	return nil
}

// GetId Id Getter
func (r TmallServiceSettleadjustmentGetAPIRequest) GetId() int64 {
	return r._id
}

var poolTmallServiceSettleadjustmentGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallServiceSettleadjustmentGetRequest()
	},
}

// GetTmallServiceSettleadjustmentGetRequest 从 sync.Pool 获取 TmallServiceSettleadjustmentGetAPIRequest
func GetTmallServiceSettleadjustmentGetAPIRequest() *TmallServiceSettleadjustmentGetAPIRequest {
	return poolTmallServiceSettleadjustmentGetAPIRequest.Get().(*TmallServiceSettleadjustmentGetAPIRequest)
}

// ReleaseTmallServiceSettleadjustmentGetAPIRequest 将 TmallServiceSettleadjustmentGetAPIRequest 放入 sync.Pool
func ReleaseTmallServiceSettleadjustmentGetAPIRequest(v *TmallServiceSettleadjustmentGetAPIRequest) {
	v.Reset()
	poolTmallServiceSettleadjustmentGetAPIRequest.Put(v)
}
