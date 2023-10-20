package mtopopen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaologisticsopenalibitywriteAPIRequest 为快递公司提供的物流信息通用写入接口 API请求
// taobao.logistics.openalibity.write
//
// 为快递公司提供的物流信息通用写入接口
type TaobaologisticsopenalibitywriteAPIRequest struct {
	model.Params
	// 信息写入请求对象
	_generalLogisticsDataWriteRequest *GeneralLogisticsDataWriteRequest
}

// NewTaobaologisticsopenalibitywriteRequest 初始化TaobaologisticsopenalibitywriteAPIRequest对象
func NewTaobaologisticsopenalibitywriteRequest() *TaobaologisticsopenalibitywriteAPIRequest {
	return &TaobaologisticsopenalibitywriteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaologisticsopenalibitywriteAPIRequest) GetApiMethodName() string {
	return "taobao.logistics.openalibity.write"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaologisticsopenalibitywriteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaologisticsopenalibitywriteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetGeneralLogisticsDataWriteRequest is GeneralLogisticsDataWriteRequest Setter
// 信息写入请求对象
func (r *TaobaologisticsopenalibitywriteAPIRequest) SetGeneralLogisticsDataWriteRequest(_generalLogisticsDataWriteRequest *GeneralLogisticsDataWriteRequest) error {
	r._generalLogisticsDataWriteRequest = _generalLogisticsDataWriteRequest
	r.Set("general_logistics_data_write_request", _generalLogisticsDataWriteRequest)
	return nil
}

// GetGeneralLogisticsDataWriteRequest GeneralLogisticsDataWriteRequest Getter
func (r TaobaologisticsopenalibitywriteAPIRequest) GetGeneralLogisticsDataWriteRequest() *GeneralLogisticsDataWriteRequest {
	return r._generalLogisticsDataWriteRequest
}
