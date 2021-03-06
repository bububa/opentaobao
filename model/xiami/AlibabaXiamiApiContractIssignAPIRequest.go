package xiami

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaXiamiApiContractIssignAPIRequest 查询是否签约 API请求
// alibaba.xiami.api.contract.issign
//
// 查询是否签约
type AlibabaXiamiApiContractIssignAPIRequest struct {
	model.Params
}

// NewAlibabaXiamiApiContractIssignRequest 初始化AlibabaXiamiApiContractIssignAPIRequest对象
func NewAlibabaXiamiApiContractIssignRequest() *AlibabaXiamiApiContractIssignAPIRequest {
	return &AlibabaXiamiApiContractIssignAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaXiamiApiContractIssignAPIRequest) GetApiMethodName() string {
	return "alibaba.xiami.api.contract.issign"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaXiamiApiContractIssignAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}
