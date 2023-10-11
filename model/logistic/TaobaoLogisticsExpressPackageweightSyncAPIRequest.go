package logistic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoLogisticsExpressPackageweightSyncAPIRequest TMS包裹重量回传 API请求
// taobao.logistics.express.packageweight.sync
//
// TMS包裹重量回传
type TaobaoLogisticsExpressPackageweightSyncAPIRequest struct {
	model.Params
	// 包裹重量信息
	_tmsPackageWeightRequest *TmsPackageWeightRequest
}

// NewTaobaoLogisticsExpressPackageweightSyncRequest 初始化TaobaoLogisticsExpressPackageweightSyncAPIRequest对象
func NewTaobaoLogisticsExpressPackageweightSyncRequest() *TaobaoLogisticsExpressPackageweightSyncAPIRequest {
	return &TaobaoLogisticsExpressPackageweightSyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoLogisticsExpressPackageweightSyncAPIRequest) GetApiMethodName() string {
	return "taobao.logistics.express.packageweight.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoLogisticsExpressPackageweightSyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoLogisticsExpressPackageweightSyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTmsPackageWeightRequest is TmsPackageWeightRequest Setter
// 包裹重量信息
func (r *TaobaoLogisticsExpressPackageweightSyncAPIRequest) SetTmsPackageWeightRequest(_tmsPackageWeightRequest *TmsPackageWeightRequest) error {
	r._tmsPackageWeightRequest = _tmsPackageWeightRequest
	r.Set("tms_package_weight_request", _tmsPackageWeightRequest)
	return nil
}

// GetTmsPackageWeightRequest TmsPackageWeightRequest Getter
func (r TaobaoLogisticsExpressPackageweightSyncAPIRequest) GetTmsPackageWeightRequest() *TmsPackageWeightRequest {
	return r._tmsPackageWeightRequest
}
