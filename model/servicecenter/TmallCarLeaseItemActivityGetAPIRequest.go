package servicecenter

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallcarleaseitemactivitygetAPIRequest 查询汽车租赁活动信息 API请求
// tmall.car.lease.item.activity.get
//
// 查询汽车租赁活动信息
type TmallcarleaseitemactivitygetAPIRequest struct {
	model.Params
}

// NewTmallcarleaseitemactivitygetRequest 初始化TmallcarleaseitemactivitygetAPIRequest对象
func NewTmallcarleaseitemactivitygetRequest() *TmallcarleaseitemactivitygetAPIRequest {
	return &TmallcarleaseitemactivitygetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallcarleaseitemactivitygetAPIRequest) GetApiMethodName() string {
	return "tmall.car.lease.item.activity.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallcarleaseitemactivitygetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallcarleaseitemactivitygetAPIRequest) GetRawParams() model.Params {
	return r.Params
}
