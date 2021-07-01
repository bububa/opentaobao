package mos

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaMjOcOfflineMaxticketnoGetAPIRequest
pos机获取线下最大小票号 API请求
alibaba.mj.oc.offline.maxticketno.get

给pos机提供线下最大小票号查询 */
type AlibabaMjOcOfflineMaxticketnoGetAPIRequest struct {
	model.Params
	// 收银机号
	_posNo string
	// 外部门店号
	_storeNo string
	// 日期
	_datetime string
}

// NewAlibabaMjOcOfflineMaxticketnoGetRequest 初始化AlibabaMjOcOfflineMaxticketnoGetAPIRequest对象
func NewAlibabaMjOcOfflineMaxticketnoGetRequest() *AlibabaMjOcOfflineMaxticketnoGetAPIRequest {
	return &AlibabaMjOcOfflineMaxticketnoGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaMjOcOfflineMaxticketnoGetAPIRequest) GetApiMethodName() string {
	return "alibaba.mj.oc.offline.maxticketno.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaMjOcOfflineMaxticketnoGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is PosNo Setter
// 收银机号
func (r *AlibabaMjOcOfflineMaxticketnoGetAPIRequest) SetPosNo(_posNo string) error {
	r._posNo = _posNo
	r.Set("pos_no", _posNo)
	return nil
}

// Get PosNo Getter
func (r AlibabaMjOcOfflineMaxticketnoGetAPIRequest) GetPosNo() string {
	return r._posNo
}

// Set is StoreNo Setter
// 外部门店号
func (r *AlibabaMjOcOfflineMaxticketnoGetAPIRequest) SetStoreNo(_storeNo string) error {
	r._storeNo = _storeNo
	r.Set("store_no", _storeNo)
	return nil
}

// Get StoreNo Getter
func (r AlibabaMjOcOfflineMaxticketnoGetAPIRequest) GetStoreNo() string {
	return r._storeNo
}

// Set is Datetime Setter
// 日期
func (r *AlibabaMjOcOfflineMaxticketnoGetAPIRequest) SetDatetime(_datetime string) error {
	r._datetime = _datetime
	r.Set("datetime", _datetime)
	return nil
}

// Get Datetime Getter
func (r AlibabaMjOcOfflineMaxticketnoGetAPIRequest) GetDatetime() string {
	return r._datetime
}
