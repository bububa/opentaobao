package smartstore

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
智能硬件互动云货架批量数据回流 API请求
taobao.smartstore.device.iashelf.batch.feedback

智慧门店互动云货架设备批量回流，
只能回流单个设备的批量业务数据
规则：
1. 回流的设备属于当前授权的用户
2. 回流的设备属于当前应用添加
*/
type TaobaoSmartstoreDeviceIashelfBatchFeedbackRequest struct {
    model.Params
    // 硬件CODE
    deviceCode   string
    // 回流数据数组，一次最多100条
    datas   []DeviceBizDataDo
}

// 初始化TaobaoSmartstoreDeviceIashelfBatchFeedbackRequest对象
func NewTaobaoSmartstoreDeviceIashelfBatchFeedbackRequest() *TaobaoSmartstoreDeviceIashelfBatchFeedbackRequest{
    return &TaobaoSmartstoreDeviceIashelfBatchFeedbackRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoSmartstoreDeviceIashelfBatchFeedbackRequest) GetApiMethodName() string {
    return "taobao.smartstore.device.iashelf.batch.feedback"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoSmartstoreDeviceIashelfBatchFeedbackRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// DeviceCode Setter
// 硬件CODE
func (r *TaobaoSmartstoreDeviceIashelfBatchFeedbackRequest) SetDeviceCode(deviceCode string) error {
    r.deviceCode = deviceCode
    r.Set("device_code", deviceCode)
    return nil
}

// DeviceCode Getter
func (r TaobaoSmartstoreDeviceIashelfBatchFeedbackRequest) GetDeviceCode() string {
    return r.deviceCode
}
// Datas Setter
// 回流数据数组，一次最多100条
func (r *TaobaoSmartstoreDeviceIashelfBatchFeedbackRequest) SetDatas(datas []DeviceBizDataDo) error {
    r.datas = datas
    r.Set("datas", datas)
    return nil
}

// Datas Getter
func (r TaobaoSmartstoreDeviceIashelfBatchFeedbackRequest) GetDatas() []DeviceBizDataDo {
    return r.datas
}
