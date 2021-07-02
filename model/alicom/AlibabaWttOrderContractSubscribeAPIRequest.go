package alicom

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWttOrderContractSubscribeAPIRequest 分销商合约生产 API请求
// alibaba.wtt.order.contract.subscribe
//
// 分销商合约生产
type AlibabaWttOrderContractSubscribeAPIRequest struct {
	model.Params
	// 分销商合约生产
	_distributionOrderModel *DistributionOrderModel
}

// NewAlibabaWttOrderContractSubscribeRequest 初始化AlibabaWttOrderContractSubscribeAPIRequest对象
func NewAlibabaWttOrderContractSubscribeRequest() *AlibabaWttOrderContractSubscribeAPIRequest {
	return &AlibabaWttOrderContractSubscribeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWttOrderContractSubscribeAPIRequest) GetApiMethodName() string {
	return "alibaba.wtt.order.contract.subscribe"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWttOrderContractSubscribeAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetDistributionOrderModel is DistributionOrderModel Setter
// 分销商合约生产
func (r *AlibabaWttOrderContractSubscribeAPIRequest) SetDistributionOrderModel(_distributionOrderModel *DistributionOrderModel) error {
	r._distributionOrderModel = _distributionOrderModel
	r.Set("distribution_order_model", _distributionOrderModel)
	return nil
}

// GetDistributionOrderModel DistributionOrderModel Getter
func (r AlibabaWttOrderContractSubscribeAPIRequest) GetDistributionOrderModel() *DistributionOrderModel {
	return r._distributionOrderModel
}
