package traveltrade

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripTravelBookinfoQueryAPIRequest 飞猪度假-订单二次预约查询接口 API请求
// alitrip.travel.bookinfo.query
//
// 飞猪度假订单二次预约详情查询接口
type AlitripTravelBookinfoQueryAPIRequest struct {
	model.Params
	// 预定信息id
	_bookinfoId int64
}

// NewAlitripTravelBookinfoQueryRequest 初始化AlitripTravelBookinfoQueryAPIRequest对象
func NewAlitripTravelBookinfoQueryRequest() *AlitripTravelBookinfoQueryAPIRequest {
	return &AlitripTravelBookinfoQueryAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripTravelBookinfoQueryAPIRequest) Reset() {
	r._bookinfoId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripTravelBookinfoQueryAPIRequest) GetApiMethodName() string {
	return "alitrip.travel.bookinfo.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripTravelBookinfoQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripTravelBookinfoQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBookinfoId is BookinfoId Setter
// 预定信息id
func (r *AlitripTravelBookinfoQueryAPIRequest) SetBookinfoId(_bookinfoId int64) error {
	r._bookinfoId = _bookinfoId
	r.Set("bookinfo_id", _bookinfoId)
	return nil
}

// GetBookinfoId BookinfoId Getter
func (r AlitripTravelBookinfoQueryAPIRequest) GetBookinfoId() int64 {
	return r._bookinfoId
}

var poolAlitripTravelBookinfoQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripTravelBookinfoQueryRequest()
	},
}

// GetAlitripTravelBookinfoQueryRequest 从 sync.Pool 获取 AlitripTravelBookinfoQueryAPIRequest
func GetAlitripTravelBookinfoQueryAPIRequest() *AlitripTravelBookinfoQueryAPIRequest {
	return poolAlitripTravelBookinfoQueryAPIRequest.Get().(*AlitripTravelBookinfoQueryAPIRequest)
}

// ReleaseAlitripTravelBookinfoQueryAPIRequest 将 AlitripTravelBookinfoQueryAPIRequest 放入 sync.Pool
func ReleaseAlitripTravelBookinfoQueryAPIRequest(v *AlitripTravelBookinfoQueryAPIRequest) {
	v.Reset()
	poolAlitripTravelBookinfoQueryAPIRequest.Put(v)
}
