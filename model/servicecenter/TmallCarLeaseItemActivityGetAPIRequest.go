package servicecenter

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallCarLeaseItemActivityGetAPIRequest 查询汽车租赁活动信息 API请求
// tmall.car.lease.item.activity.get
//
// 查询汽车租赁活动信息
type TmallCarLeaseItemActivityGetAPIRequest struct {
	model.Params
}

// NewTmallCarLeaseItemActivityGetRequest 初始化TmallCarLeaseItemActivityGetAPIRequest对象
func NewTmallCarLeaseItemActivityGetRequest() *TmallCarLeaseItemActivityGetAPIRequest {
	return &TmallCarLeaseItemActivityGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallCarLeaseItemActivityGetAPIRequest) GetApiMethodName() string {
	return "tmall.car.lease.item.activity.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallCarLeaseItemActivityGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}
