package util

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabamosfalconposcounterqueryAPIRequest 云POS查看专柜属性 API请求
// alibaba.mos.falcon.pos.counter.query
//
// 银泰商业获取专柜是否支持小数等属性查看
type AlibabamosfalconposcounterqueryAPIRequest struct {
	model.Params
	// 设备序列号
	_sn string
	// 门店号
	_storeNo string
	// 专柜号
	_counterNo string
}

// NewAlibabamosfalconposcounterqueryRequest 初始化AlibabamosfalconposcounterqueryAPIRequest对象
func NewAlibabamosfalconposcounterqueryRequest() *AlibabamosfalconposcounterqueryAPIRequest {
	return &AlibabamosfalconposcounterqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabamosfalconposcounterqueryAPIRequest) GetApiMethodName() string {
	return "alibaba.mos.falcon.pos.counter.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabamosfalconposcounterqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabamosfalconposcounterqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSn is Sn Setter
// 设备序列号
func (r *AlibabamosfalconposcounterqueryAPIRequest) SetSn(_sn string) error {
	r._sn = _sn
	r.Set("sn", _sn)
	return nil
}

// GetSn Sn Getter
func (r AlibabamosfalconposcounterqueryAPIRequest) GetSn() string {
	return r._sn
}

// SetStoreNo is StoreNo Setter
// 门店号
func (r *AlibabamosfalconposcounterqueryAPIRequest) SetStoreNo(_storeNo string) error {
	r._storeNo = _storeNo
	r.Set("store_no", _storeNo)
	return nil
}

// GetStoreNo StoreNo Getter
func (r AlibabamosfalconposcounterqueryAPIRequest) GetStoreNo() string {
	return r._storeNo
}

// SetCounterNo is CounterNo Setter
// 专柜号
func (r *AlibabamosfalconposcounterqueryAPIRequest) SetCounterNo(_counterNo string) error {
	r._counterNo = _counterNo
	r.Set("counter_no", _counterNo)
	return nil
}

// GetCounterNo CounterNo Getter
func (r AlibabamosfalconposcounterqueryAPIRequest) GetCounterNo() string {
	return r._counterNo
}
