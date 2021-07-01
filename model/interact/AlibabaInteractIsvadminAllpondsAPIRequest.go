package interact

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaInteractIsvadminAllpondsAPIRequest
获取天猫互动奖池列表 API请求
alibaba.interact.isvadmin.allponds

获取天猫互动奖池列表 */
type AlibabaInteractIsvadminAllpondsAPIRequest struct {
	model.Params
}

// NewAlibabaInteractIsvadminAllpondsRequest 初始化AlibabaInteractIsvadminAllpondsAPIRequest对象
func NewAlibabaInteractIsvadminAllpondsRequest() *AlibabaInteractIsvadminAllpondsAPIRequest {
	return &AlibabaInteractIsvadminAllpondsAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaInteractIsvadminAllpondsAPIRequest) GetApiMethodName() string {
	return "alibaba.interact.isvadmin.allponds"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaInteractIsvadminAllpondsAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}
