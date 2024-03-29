package promotion

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUmpDetailGetAPIRequest 查询活动详情 API请求
// taobao.ump.detail.get
//
// 查询活动详情
type TaobaoUmpDetailGetAPIRequest struct {
	model.Params
	// 活动详情的id
	_detailId int64
}

// NewTaobaoUmpDetailGetRequest 初始化TaobaoUmpDetailGetAPIRequest对象
func NewTaobaoUmpDetailGetRequest() *TaobaoUmpDetailGetAPIRequest {
	return &TaobaoUmpDetailGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoUmpDetailGetAPIRequest) Reset() {
	r._detailId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoUmpDetailGetAPIRequest) GetApiMethodName() string {
	return "taobao.ump.detail.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoUmpDetailGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoUmpDetailGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDetailId is DetailId Setter
// 活动详情的id
func (r *TaobaoUmpDetailGetAPIRequest) SetDetailId(_detailId int64) error {
	r._detailId = _detailId
	r.Set("detail_id", _detailId)
	return nil
}

// GetDetailId DetailId Getter
func (r TaobaoUmpDetailGetAPIRequest) GetDetailId() int64 {
	return r._detailId
}

var poolTaobaoUmpDetailGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoUmpDetailGetRequest()
	},
}

// GetTaobaoUmpDetailGetRequest 从 sync.Pool 获取 TaobaoUmpDetailGetAPIRequest
func GetTaobaoUmpDetailGetAPIRequest() *TaobaoUmpDetailGetAPIRequest {
	return poolTaobaoUmpDetailGetAPIRequest.Get().(*TaobaoUmpDetailGetAPIRequest)
}

// ReleaseTaobaoUmpDetailGetAPIRequest 将 TaobaoUmpDetailGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoUmpDetailGetAPIRequest(v *TaobaoUmpDetailGetAPIRequest) {
	v.Reset()
	poolTaobaoUmpDetailGetAPIRequest.Put(v)
}
