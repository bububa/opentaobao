package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaofenxiaogradesgetAPIRequest 分销商等级查询 API请求
// taobao.fenxiao.grades.get
//
// 根据供应商ID，查询他的分销商等级信息
type TaobaofenxiaogradesgetAPIRequest struct {
	model.Params
}

// NewTaobaofenxiaogradesgetRequest 初始化TaobaofenxiaogradesgetAPIRequest对象
func NewTaobaofenxiaogradesgetRequest() *TaobaofenxiaogradesgetAPIRequest {
	return &TaobaofenxiaogradesgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaofenxiaogradesgetAPIRequest) GetApiMethodName() string {
	return "taobao.fenxiao.grades.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaofenxiaogradesgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaofenxiaogradesgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}
