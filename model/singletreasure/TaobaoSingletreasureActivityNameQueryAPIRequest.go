package singletreasure

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaosingletreasureactivitynamequeryAPIRequest 查询官方的活动名称接口 API请求
// taobao.singletreasure.activity.name.query
//
// 查询官方的活动名称列表接口
type TaobaosingletreasureactivitynamequeryAPIRequest struct {
	model.Params
}

// NewTaobaosingletreasureactivitynamequeryRequest 初始化TaobaosingletreasureactivitynamequeryAPIRequest对象
func NewTaobaosingletreasureactivitynamequeryRequest() *TaobaosingletreasureactivitynamequeryAPIRequest {
	return &TaobaosingletreasureactivitynamequeryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaosingletreasureactivitynamequeryAPIRequest) GetApiMethodName() string {
	return "taobao.singletreasure.activity.name.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaosingletreasureactivitynamequeryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaosingletreasureactivitynamequeryAPIRequest) GetRawParams() model.Params {
	return r.Params
}
