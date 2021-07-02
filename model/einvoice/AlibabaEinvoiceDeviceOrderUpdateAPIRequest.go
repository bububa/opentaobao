package einvoice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaEinvoiceDeviceOrderUpdateAPIRequest 回传/更新设备订购单 API请求
// alibaba.einvoice.device.order.update
//
// 更新设备订购单，同步税控设备信息
type AlibabaEinvoiceDeviceOrderUpdateAPIRequest struct {
	model.Params
	// 订购单工单事件：  deploy_finish: 设备就绪，部署完成  isv_reject: 服务商驳回订购单
	_action string
	// 税控设备ID
	_deviceId string
	// 拓展字段。  ①当action=deploy_finish时，拓展字段中必须包含：  serv_start_time: 服务有效周期-起始时间  serv_end_time: 服务有效周期-结束时间  时间格式：yyyy-MM-dd HH:mm:ss  ②当action=isv_reject时，拓展字段中必须包含：  message: 驳回原因
	_extJson string
	// 订购开通单ID
	_flowId string
	// 税号
	_payeeRegisterNo string
}

// NewAlibabaEinvoiceDeviceOrderUpdateRequest 初始化AlibabaEinvoiceDeviceOrderUpdateAPIRequest对象
func NewAlibabaEinvoiceDeviceOrderUpdateRequest() *AlibabaEinvoiceDeviceOrderUpdateAPIRequest {
	return &AlibabaEinvoiceDeviceOrderUpdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaEinvoiceDeviceOrderUpdateAPIRequest) GetApiMethodName() string {
	return "alibaba.einvoice.device.order.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaEinvoiceDeviceOrderUpdateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetAction is Action Setter
// 订购单工单事件：  deploy_finish: 设备就绪，部署完成  isv_reject: 服务商驳回订购单
func (r *AlibabaEinvoiceDeviceOrderUpdateAPIRequest) SetAction(_action string) error {
	r._action = _action
	r.Set("action", _action)
	return nil
}

// GetAction Action Getter
func (r AlibabaEinvoiceDeviceOrderUpdateAPIRequest) GetAction() string {
	return r._action
}

// SetDeviceId is DeviceId Setter
// 税控设备ID
func (r *AlibabaEinvoiceDeviceOrderUpdateAPIRequest) SetDeviceId(_deviceId string) error {
	r._deviceId = _deviceId
	r.Set("device_id", _deviceId)
	return nil
}

// GetDeviceId DeviceId Getter
func (r AlibabaEinvoiceDeviceOrderUpdateAPIRequest) GetDeviceId() string {
	return r._deviceId
}

// SetExtJson is ExtJson Setter
// 拓展字段。  ①当action=deploy_finish时，拓展字段中必须包含：  serv_start_time: 服务有效周期-起始时间  serv_end_time: 服务有效周期-结束时间  时间格式：yyyy-MM-dd HH:mm:ss  ②当action=isv_reject时，拓展字段中必须包含：  message: 驳回原因
func (r *AlibabaEinvoiceDeviceOrderUpdateAPIRequest) SetExtJson(_extJson string) error {
	r._extJson = _extJson
	r.Set("ext_json", _extJson)
	return nil
}

// GetExtJson ExtJson Getter
func (r AlibabaEinvoiceDeviceOrderUpdateAPIRequest) GetExtJson() string {
	return r._extJson
}

// SetFlowId is FlowId Setter
// 订购开通单ID
func (r *AlibabaEinvoiceDeviceOrderUpdateAPIRequest) SetFlowId(_flowId string) error {
	r._flowId = _flowId
	r.Set("flow_id", _flowId)
	return nil
}

// GetFlowId FlowId Getter
func (r AlibabaEinvoiceDeviceOrderUpdateAPIRequest) GetFlowId() string {
	return r._flowId
}

// SetPayeeRegisterNo is PayeeRegisterNo Setter
// 税号
func (r *AlibabaEinvoiceDeviceOrderUpdateAPIRequest) SetPayeeRegisterNo(_payeeRegisterNo string) error {
	r._payeeRegisterNo = _payeeRegisterNo
	r.Set("payee_register_no", _payeeRegisterNo)
	return nil
}

// GetPayeeRegisterNo PayeeRegisterNo Getter
func (r AlibabaEinvoiceDeviceOrderUpdateAPIRequest) GetPayeeRegisterNo() string {
	return r._payeeRegisterNo
}
