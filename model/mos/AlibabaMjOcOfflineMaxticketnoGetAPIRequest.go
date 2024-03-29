package mos

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMjOcOfflineMaxticketnoGetAPIRequest pos机获取线下最大小票号 API请求
// alibaba.mj.oc.offline.maxticketno.get
//
// 给pos机提供线下最大小票号查询
type AlibabaMjOcOfflineMaxticketnoGetAPIRequest struct {
	model.Params
	// 日期
	_datetime string
	// 收银机号
	_posNo string
	// 外部门店号
	_storeNo string
}

// NewAlibabaMjOcOfflineMaxticketnoGetRequest 初始化AlibabaMjOcOfflineMaxticketnoGetAPIRequest对象
func NewAlibabaMjOcOfflineMaxticketnoGetRequest() *AlibabaMjOcOfflineMaxticketnoGetAPIRequest {
	return &AlibabaMjOcOfflineMaxticketnoGetAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaMjOcOfflineMaxticketnoGetAPIRequest) Reset() {
	r._datetime = ""
	r._posNo = ""
	r._storeNo = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaMjOcOfflineMaxticketnoGetAPIRequest) GetApiMethodName() string {
	return "alibaba.mj.oc.offline.maxticketno.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaMjOcOfflineMaxticketnoGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaMjOcOfflineMaxticketnoGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDatetime is Datetime Setter
// 日期
func (r *AlibabaMjOcOfflineMaxticketnoGetAPIRequest) SetDatetime(_datetime string) error {
	r._datetime = _datetime
	r.Set("datetime", _datetime)
	return nil
}

// GetDatetime Datetime Getter
func (r AlibabaMjOcOfflineMaxticketnoGetAPIRequest) GetDatetime() string {
	return r._datetime
}

// SetPosNo is PosNo Setter
// 收银机号
func (r *AlibabaMjOcOfflineMaxticketnoGetAPIRequest) SetPosNo(_posNo string) error {
	r._posNo = _posNo
	r.Set("pos_no", _posNo)
	return nil
}

// GetPosNo PosNo Getter
func (r AlibabaMjOcOfflineMaxticketnoGetAPIRequest) GetPosNo() string {
	return r._posNo
}

// SetStoreNo is StoreNo Setter
// 外部门店号
func (r *AlibabaMjOcOfflineMaxticketnoGetAPIRequest) SetStoreNo(_storeNo string) error {
	r._storeNo = _storeNo
	r.Set("store_no", _storeNo)
	return nil
}

// GetStoreNo StoreNo Getter
func (r AlibabaMjOcOfflineMaxticketnoGetAPIRequest) GetStoreNo() string {
	return r._storeNo
}

var poolAlibabaMjOcOfflineMaxticketnoGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaMjOcOfflineMaxticketnoGetRequest()
	},
}

// GetAlibabaMjOcOfflineMaxticketnoGetRequest 从 sync.Pool 获取 AlibabaMjOcOfflineMaxticketnoGetAPIRequest
func GetAlibabaMjOcOfflineMaxticketnoGetAPIRequest() *AlibabaMjOcOfflineMaxticketnoGetAPIRequest {
	return poolAlibabaMjOcOfflineMaxticketnoGetAPIRequest.Get().(*AlibabaMjOcOfflineMaxticketnoGetAPIRequest)
}

// ReleaseAlibabaMjOcOfflineMaxticketnoGetAPIRequest 将 AlibabaMjOcOfflineMaxticketnoGetAPIRequest 放入 sync.Pool
func ReleaseAlibabaMjOcOfflineMaxticketnoGetAPIRequest(v *AlibabaMjOcOfflineMaxticketnoGetAPIRequest) {
	v.Reset()
	poolAlibabaMjOcOfflineMaxticketnoGetAPIRequest.Put(v)
}
