package tmallsc

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallServicecenterReservecondCreateAPIRequest 创建主动预约开通条件 API请求
// tmall.servicecenter.reservecond.create
//
// 1、设置主动预约开通条件
type TmallServicecenterReservecondCreateAPIRequest struct {
	model.Params
	// 主动预约开通条件
	_rocList []ReserveOpenConditionDto
}

// NewTmallServicecenterReservecondCreateRequest 初始化TmallServicecenterReservecondCreateAPIRequest对象
func NewTmallServicecenterReservecondCreateRequest() *TmallServicecenterReservecondCreateAPIRequest {
	return &TmallServicecenterReservecondCreateAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallServicecenterReservecondCreateAPIRequest) Reset() {
	r._rocList = r._rocList[:0]
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallServicecenterReservecondCreateAPIRequest) GetApiMethodName() string {
	return "tmall.servicecenter.reservecond.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallServicecenterReservecondCreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallServicecenterReservecondCreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRocList is RocList Setter
// 主动预约开通条件
func (r *TmallServicecenterReservecondCreateAPIRequest) SetRocList(_rocList []ReserveOpenConditionDto) error {
	r._rocList = _rocList
	r.Set("roc_list", _rocList)
	return nil
}

// GetRocList RocList Getter
func (r TmallServicecenterReservecondCreateAPIRequest) GetRocList() []ReserveOpenConditionDto {
	return r._rocList
}

var poolTmallServicecenterReservecondCreateAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallServicecenterReservecondCreateRequest()
	},
}

// GetTmallServicecenterReservecondCreateRequest 从 sync.Pool 获取 TmallServicecenterReservecondCreateAPIRequest
func GetTmallServicecenterReservecondCreateAPIRequest() *TmallServicecenterReservecondCreateAPIRequest {
	return poolTmallServicecenterReservecondCreateAPIRequest.Get().(*TmallServicecenterReservecondCreateAPIRequest)
}

// ReleaseTmallServicecenterReservecondCreateAPIRequest 将 TmallServicecenterReservecondCreateAPIRequest 放入 sync.Pool
func ReleaseTmallServicecenterReservecondCreateAPIRequest(v *TmallServicecenterReservecondCreateAPIRequest) {
	v.Reset()
	poolTmallServicecenterReservecondCreateAPIRequest.Put(v)
}
