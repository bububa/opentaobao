package smartstore

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoSmartstoreDeviceIashelfBatchFeedbackAPIRequest
智能硬件互动云货架批量数据回流 API请求
taobao.smartstore.device.iashelf.batch.feedback

智慧门店互动云货架设备批量回流，
只能回流单个设备的批量业务数据
规则：
1. 回流的设备属于当前授权的用户
2. 回流的设备属于当前应用添加 */
type TaobaoSmartstoreDeviceIashelfBatchFeedbackAPIRequest struct {
	model.Params
	// 硬件CODE
	_deviceCode string
	// 回流数据数组，一次最多100条
	_datas []DeviceBizDataDo
}

// NewTaobaoSmartstoreDeviceIashelfBatchFeedbackRequest 初始化TaobaoSmartstoreDeviceIashelfBatchFeedbackAPIRequest对象
func NewTaobaoSmartstoreDeviceIashelfBatchFeedbackRequest() *TaobaoSmartstoreDeviceIashelfBatchFeedbackAPIRequest {
	return &TaobaoSmartstoreDeviceIashelfBatchFeedbackAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSmartstoreDeviceIashelfBatchFeedbackAPIRequest) GetApiMethodName() string {
	return "taobao.smartstore.device.iashelf.batch.feedback"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSmartstoreDeviceIashelfBatchFeedbackAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is DeviceCode Setter
// 硬件CODE
func (r *TaobaoSmartstoreDeviceIashelfBatchFeedbackAPIRequest) SetDeviceCode(_deviceCode string) error {
	r._deviceCode = _deviceCode
	r.Set("device_code", _deviceCode)
	return nil
}

// Get DeviceCode Getter
func (r TaobaoSmartstoreDeviceIashelfBatchFeedbackAPIRequest) GetDeviceCode() string {
	return r._deviceCode
}

// Set is Datas Setter
// 回流数据数组，一次最多100条
func (r *TaobaoSmartstoreDeviceIashelfBatchFeedbackAPIRequest) SetDatas(_datas []DeviceBizDataDo) error {
	r._datas = _datas
	r.Set("datas", _datas)
	return nil
}

// Get Datas Getter
func (r TaobaoSmartstoreDeviceIashelfBatchFeedbackAPIRequest) GetDatas() []DeviceBizDataDo {
	return r._datas
}
