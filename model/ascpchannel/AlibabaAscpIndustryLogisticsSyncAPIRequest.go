package ascpchannel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAscpIndustryLogisticsSyncAPIRequest 物流状态同步 API请求
// alibaba.ascp.industry.logistics.sync
//
// 履约物流状态同步
type AlibabaAscpIndustryLogisticsSyncAPIRequest struct {
	model.Params
	// 参数
	_param *LogisticsSyncSellerRequest
}

// NewAlibabaAscpIndustryLogisticsSyncRequest 初始化AlibabaAscpIndustryLogisticsSyncAPIRequest对象
func NewAlibabaAscpIndustryLogisticsSyncRequest() *AlibabaAscpIndustryLogisticsSyncAPIRequest {
	return &AlibabaAscpIndustryLogisticsSyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAscpIndustryLogisticsSyncAPIRequest) GetApiMethodName() string {
	return "alibaba.ascp.industry.logistics.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAscpIndustryLogisticsSyncAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetParam is Param Setter
// 参数
func (r *AlibabaAscpIndustryLogisticsSyncAPIRequest) SetParam(_param *LogisticsSyncSellerRequest) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabaAscpIndustryLogisticsSyncAPIRequest) GetParam() *LogisticsSyncSellerRequest {
	return r._param
}
