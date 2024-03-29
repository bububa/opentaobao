package promotion

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUmpDetailDeleteAPIRequest 删除活动详情 API请求
// taobao.ump.detail.delete
//
// 删除活动详情
type TaobaoUmpDetailDeleteAPIRequest struct {
	model.Params
	// 活动详情id
	_detailId int64
}

// NewTaobaoUmpDetailDeleteRequest 初始化TaobaoUmpDetailDeleteAPIRequest对象
func NewTaobaoUmpDetailDeleteRequest() *TaobaoUmpDetailDeleteAPIRequest {
	return &TaobaoUmpDetailDeleteAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoUmpDetailDeleteAPIRequest) Reset() {
	r._detailId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoUmpDetailDeleteAPIRequest) GetApiMethodName() string {
	return "taobao.ump.detail.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoUmpDetailDeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoUmpDetailDeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDetailId is DetailId Setter
// 活动详情id
func (r *TaobaoUmpDetailDeleteAPIRequest) SetDetailId(_detailId int64) error {
	r._detailId = _detailId
	r.Set("detail_id", _detailId)
	return nil
}

// GetDetailId DetailId Getter
func (r TaobaoUmpDetailDeleteAPIRequest) GetDetailId() int64 {
	return r._detailId
}

var poolTaobaoUmpDetailDeleteAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoUmpDetailDeleteRequest()
	},
}

// GetTaobaoUmpDetailDeleteRequest 从 sync.Pool 获取 TaobaoUmpDetailDeleteAPIRequest
func GetTaobaoUmpDetailDeleteAPIRequest() *TaobaoUmpDetailDeleteAPIRequest {
	return poolTaobaoUmpDetailDeleteAPIRequest.Get().(*TaobaoUmpDetailDeleteAPIRequest)
}

// ReleaseTaobaoUmpDetailDeleteAPIRequest 将 TaobaoUmpDetailDeleteAPIRequest 放入 sync.Pool
func ReleaseTaobaoUmpDetailDeleteAPIRequest(v *TaobaoUmpDetailDeleteAPIRequest) {
	v.Reset()
	poolTaobaoUmpDetailDeleteAPIRequest.Put(v)
}
