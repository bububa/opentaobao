package util

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMosFalconPosCounterQueryAPIRequest 云POS查看专柜属性 API请求
// alibaba.mos.falcon.pos.counter.query
//
// 银泰商业获取专柜是否支持小数等属性查看
type AlibabaMosFalconPosCounterQueryAPIRequest struct {
	model.Params
	// 设备序列号
	_sn string
	// 门店号
	_storeNo string
	// 专柜号
	_counterNo string
}

// NewAlibabaMosFalconPosCounterQueryRequest 初始化AlibabaMosFalconPosCounterQueryAPIRequest对象
func NewAlibabaMosFalconPosCounterQueryRequest() *AlibabaMosFalconPosCounterQueryAPIRequest {
	return &AlibabaMosFalconPosCounterQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaMosFalconPosCounterQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.mos.falcon.pos.counter.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaMosFalconPosCounterQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetSn is Sn Setter
// 设备序列号
func (r *AlibabaMosFalconPosCounterQueryAPIRequest) SetSn(_sn string) error {
	r._sn = _sn
	r.Set("sn", _sn)
	return nil
}

// GetSn Sn Getter
func (r AlibabaMosFalconPosCounterQueryAPIRequest) GetSn() string {
	return r._sn
}

// SetStoreNo is StoreNo Setter
// 门店号
func (r *AlibabaMosFalconPosCounterQueryAPIRequest) SetStoreNo(_storeNo string) error {
	r._storeNo = _storeNo
	r.Set("store_no", _storeNo)
	return nil
}

// GetStoreNo StoreNo Getter
func (r AlibabaMosFalconPosCounterQueryAPIRequest) GetStoreNo() string {
	return r._storeNo
}

// SetCounterNo is CounterNo Setter
// 专柜号
func (r *AlibabaMosFalconPosCounterQueryAPIRequest) SetCounterNo(_counterNo string) error {
	r._counterNo = _counterNo
	r.Set("counter_no", _counterNo)
	return nil
}

// GetCounterNo CounterNo Getter
func (r AlibabaMosFalconPosCounterQueryAPIRequest) GetCounterNo() string {
	return r._counterNo
}
