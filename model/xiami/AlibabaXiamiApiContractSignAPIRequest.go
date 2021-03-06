package xiami

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaXiamiApiContractSignAPIRequest 提供签约链接 API请求
// alibaba.xiami.api.contract.sign
//
// 提供签约链接。in：商家id；out：签约url
type AlibabaXiamiApiContractSignAPIRequest struct {
	model.Params
}

// NewAlibabaXiamiApiContractSignRequest 初始化AlibabaXiamiApiContractSignAPIRequest对象
func NewAlibabaXiamiApiContractSignRequest() *AlibabaXiamiApiContractSignAPIRequest {
	return &AlibabaXiamiApiContractSignAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaXiamiApiContractSignAPIRequest) GetApiMethodName() string {
	return "alibaba.xiami.api.contract.sign"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaXiamiApiContractSignAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}
