package logistic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaologisticsexpresspackageweightsyncAPIRequest TMS包裹重量回传 API请求
// taobao.logistics.express.packageweight.sync
//
// TMS包裹重量回传
type TaobaologisticsexpresspackageweightsyncAPIRequest struct {
	model.Params
	// 包裹重量信息
	_tmsPackageWeightRequest *TmsPackageWeightRequest
}

// NewTaobaologisticsexpresspackageweightsyncRequest 初始化TaobaologisticsexpresspackageweightsyncAPIRequest对象
func NewTaobaologisticsexpresspackageweightsyncRequest() *TaobaologisticsexpresspackageweightsyncAPIRequest {
	return &TaobaologisticsexpresspackageweightsyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaologisticsexpresspackageweightsyncAPIRequest) GetApiMethodName() string {
	return "taobao.logistics.express.packageweight.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaologisticsexpresspackageweightsyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaologisticsexpresspackageweightsyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTmsPackageWeightRequest is TmsPackageWeightRequest Setter
// 包裹重量信息
func (r *TaobaologisticsexpresspackageweightsyncAPIRequest) SetTmsPackageWeightRequest(_tmsPackageWeightRequest *TmsPackageWeightRequest) error {
	r._tmsPackageWeightRequest = _tmsPackageWeightRequest
	r.Set("tms_package_weight_request", _tmsPackageWeightRequest)
	return nil
}

// GetTmsPackageWeightRequest TmsPackageWeightRequest Getter
func (r TaobaologisticsexpresspackageweightsyncAPIRequest) GetTmsPackageWeightRequest() *TmsPackageWeightRequest {
	return r._tmsPackageWeightRequest
}
