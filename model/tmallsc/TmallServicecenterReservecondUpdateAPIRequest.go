package tmallsc

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallServicecenterReservecondUpdateAPIRequest 主动预约条件更新 API请求
// tmall.servicecenter.reservecond.update
//
// 1、设置主动预约开通条件
type TmallServicecenterReservecondUpdateAPIRequest struct {
	model.Params
	// 主动预约开通条件
	_rocList []ReserveOpenConditionDto
}

// NewTmallServicecenterReservecondUpdateRequest 初始化TmallServicecenterReservecondUpdateAPIRequest对象
func NewTmallServicecenterReservecondUpdateRequest() *TmallServicecenterReservecondUpdateAPIRequest {
	return &TmallServicecenterReservecondUpdateAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallServicecenterReservecondUpdateAPIRequest) Reset() {
	r._rocList = r._rocList[:0]
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallServicecenterReservecondUpdateAPIRequest) GetApiMethodName() string {
	return "tmall.servicecenter.reservecond.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallServicecenterReservecondUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallServicecenterReservecondUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRocList is RocList Setter
// 主动预约开通条件
func (r *TmallServicecenterReservecondUpdateAPIRequest) SetRocList(_rocList []ReserveOpenConditionDto) error {
	r._rocList = _rocList
	r.Set("roc_list", _rocList)
	return nil
}

// GetRocList RocList Getter
func (r TmallServicecenterReservecondUpdateAPIRequest) GetRocList() []ReserveOpenConditionDto {
	return r._rocList
}

var poolTmallServicecenterReservecondUpdateAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallServicecenterReservecondUpdateRequest()
	},
}

// GetTmallServicecenterReservecondUpdateRequest 从 sync.Pool 获取 TmallServicecenterReservecondUpdateAPIRequest
func GetTmallServicecenterReservecondUpdateAPIRequest() *TmallServicecenterReservecondUpdateAPIRequest {
	return poolTmallServicecenterReservecondUpdateAPIRequest.Get().(*TmallServicecenterReservecondUpdateAPIRequest)
}

// ReleaseTmallServicecenterReservecondUpdateAPIRequest 将 TmallServicecenterReservecondUpdateAPIRequest 放入 sync.Pool
func ReleaseTmallServicecenterReservecondUpdateAPIRequest(v *TmallServicecenterReservecondUpdateAPIRequest) {
	v.Reset()
	poolTmallServicecenterReservecondUpdateAPIRequest.Put(v)
}
