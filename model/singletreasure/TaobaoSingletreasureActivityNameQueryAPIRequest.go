package singletreasure

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSingletreasureActivityNameQueryAPIRequest 查询官方的活动名称接口 API请求
// taobao.singletreasure.activity.name.query
//
// 查询官方的活动名称列表接口
type TaobaoSingletreasureActivityNameQueryAPIRequest struct {
	model.Params
}

// NewTaobaoSingletreasureActivityNameQueryRequest 初始化TaobaoSingletreasureActivityNameQueryAPIRequest对象
func NewTaobaoSingletreasureActivityNameQueryRequest() *TaobaoSingletreasureActivityNameQueryAPIRequest {
	return &TaobaoSingletreasureActivityNameQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSingletreasureActivityNameQueryAPIRequest) GetApiMethodName() string {
	return "taobao.singletreasure.activity.name.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSingletreasureActivityNameQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}
