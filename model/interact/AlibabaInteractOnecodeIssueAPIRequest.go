package interact

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabainteractonecodeissueAPIRequest onecode代码通用鉴权 API请求
// alibaba.interact.onecode.issue
//
// 手淘开放鉴权接口，仅用于tida接口鉴权，无输入输出。
type AlibabainteractonecodeissueAPIRequest struct {
	model.Params
}

// NewAlibabainteractonecodeissueRequest 初始化AlibabainteractonecodeissueAPIRequest对象
func NewAlibabainteractonecodeissueRequest() *AlibabainteractonecodeissueAPIRequest {
	return &AlibabainteractonecodeissueAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabainteractonecodeissueAPIRequest) GetApiMethodName() string {
	return "alibaba.interact.onecode.issue"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabainteractonecodeissueAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabainteractonecodeissueAPIRequest) GetRawParams() model.Params {
	return r.Params
}
