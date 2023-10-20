package caipiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaocaipiaosignstatuscheckAPIRequest 检查用户是否签署支付宝代购协议 API请求
// taobao.caipiao.signstatus.check
//
// 检查用户是否签署了支付宝代扣协议。如果签署了，返回true; 如果没签署，返回false, 同时返回签署代扣协议的Url。
type TaobaocaipiaosignstatuscheckAPIRequest struct {
	model.Params
}

// NewTaobaocaipiaosignstatuscheckRequest 初始化TaobaocaipiaosignstatuscheckAPIRequest对象
func NewTaobaocaipiaosignstatuscheckRequest() *TaobaocaipiaosignstatuscheckAPIRequest {
	return &TaobaocaipiaosignstatuscheckAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaocaipiaosignstatuscheckAPIRequest) GetApiMethodName() string {
	return "taobao.caipiao.signstatus.check"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaocaipiaosignstatuscheckAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaocaipiaosignstatuscheckAPIRequest) GetRawParams() model.Params {
	return r.Params
}
