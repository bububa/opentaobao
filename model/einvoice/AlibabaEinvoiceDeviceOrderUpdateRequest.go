package einvoice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
回传/更新设备订购单 API请求
alibaba.einvoice.device.order.update

更新设备订购单，同步税控设备信息
*/
type AlibabaEinvoiceDeviceOrderUpdateRequest struct {
    model.Params
    // 订购单工单事件：  deploy_finish: 设备就绪，部署完成  isv_reject: 服务商驳回订购单
    _action   string
    // 税控设备ID
    _deviceId   string
    // 拓展字段。  ①当action=deploy_finish时，拓展字段中必须包含：  serv_start_time: 服务有效周期-起始时间  serv_end_time: 服务有效周期-结束时间  时间格式：yyyy-MM-dd HH:mm:ss  ②当action=isv_reject时，拓展字段中必须包含：  message: 驳回原因
    _extJson   string
    // 订购开通单ID
    _flowId   string
    // 税号
    _payeeRegisterNo   string
}

// 初始化AlibabaEinvoiceDeviceOrderUpdateRequest对象
func NewAlibabaEinvoiceDeviceOrderUpdateRequest() *AlibabaEinvoiceDeviceOrderUpdateRequest{
    return &AlibabaEinvoiceDeviceOrderUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaEinvoiceDeviceOrderUpdateRequest) GetApiMethodName() string {
    return "alibaba.einvoice.device.order.update"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaEinvoiceDeviceOrderUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Action Setter
// 订购单工单事件：  deploy_finish: 设备就绪，部署完成  isv_reject: 服务商驳回订购单
func (r *AlibabaEinvoiceDeviceOrderUpdateRequest) SetAction(_action string) error {
    r._action = _action
    r.Set("action", _action)
    return nil
}

// Action Getter
func (r AlibabaEinvoiceDeviceOrderUpdateRequest) GetAction() string {
    return r._action
}
// DeviceId Setter
// 税控设备ID
func (r *AlibabaEinvoiceDeviceOrderUpdateRequest) SetDeviceId(_deviceId string) error {
    r._deviceId = _deviceId
    r.Set("device_id", _deviceId)
    return nil
}

// DeviceId Getter
func (r AlibabaEinvoiceDeviceOrderUpdateRequest) GetDeviceId() string {
    return r._deviceId
}
// ExtJson Setter
// 拓展字段。  ①当action=deploy_finish时，拓展字段中必须包含：  serv_start_time: 服务有效周期-起始时间  serv_end_time: 服务有效周期-结束时间  时间格式：yyyy-MM-dd HH:mm:ss  ②当action=isv_reject时，拓展字段中必须包含：  message: 驳回原因
func (r *AlibabaEinvoiceDeviceOrderUpdateRequest) SetExtJson(_extJson string) error {
    r._extJson = _extJson
    r.Set("ext_json", _extJson)
    return nil
}

// ExtJson Getter
func (r AlibabaEinvoiceDeviceOrderUpdateRequest) GetExtJson() string {
    return r._extJson
}
// FlowId Setter
// 订购开通单ID
func (r *AlibabaEinvoiceDeviceOrderUpdateRequest) SetFlowId(_flowId string) error {
    r._flowId = _flowId
    r.Set("flow_id", _flowId)
    return nil
}

// FlowId Getter
func (r AlibabaEinvoiceDeviceOrderUpdateRequest) GetFlowId() string {
    return r._flowId
}
// PayeeRegisterNo Setter
// 税号
func (r *AlibabaEinvoiceDeviceOrderUpdateRequest) SetPayeeRegisterNo(_payeeRegisterNo string) error {
    r._payeeRegisterNo = _payeeRegisterNo
    r.Set("payee_register_no", _payeeRegisterNo)
    return nil
}

// PayeeRegisterNo Getter
func (r AlibabaEinvoiceDeviceOrderUpdateRequest) GetPayeeRegisterNo() string {
    return r._payeeRegisterNo
}
