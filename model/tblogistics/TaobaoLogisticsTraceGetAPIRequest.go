package tblogistics

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaologisticstracegetAPIRequest 用户根据交易号查询物流流转信息 API请求
// taobao.logistics.trace.get
//
// 用户根据交易号查询物流流转信息
type TaobaologisticstracegetAPIRequest struct {
	model.Params
	// 交易号
	_tid int64
}

// NewTaobaologisticstracegetRequest 初始化TaobaologisticstracegetAPIRequest对象
func NewTaobaologisticstracegetRequest() *TaobaologisticstracegetAPIRequest {
	return &TaobaologisticstracegetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaologisticstracegetAPIRequest) GetApiMethodName() string {
	return "taobao.logistics.trace.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaologisticstracegetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaologisticstracegetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTid is Tid Setter
// 交易号
func (r *TaobaologisticstracegetAPIRequest) SetTid(_tid int64) error {
	r._tid = _tid
	r.Set("tid", _tid)
	return nil
}

// GetTid Tid Getter
func (r TaobaologisticstracegetAPIRequest) GetTid() int64 {
	return r._tid
}
