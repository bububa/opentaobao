package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthdrugkytdrtransportuploadAPIRequest 疫苗运输温度上传 API请求
// alibaba.alihealth.drug.kyt.dr.transportupload
//
// 疫苗运输温度上传
type AlibabaalihealthdrugkytdrtransportuploadAPIRequest struct {
	model.Params
	// 企业RefEntid，是指该单据的所有者。         如企业A上传了一个出库单，您为A的单据上传运输温度，那么此时RefEntid即为A的ID。         若您本企业上传了一个单据，您为这个单据上传温度，此时RefEntid即为您本企业的ID。
	_refEntId string
	// 单据编号
	_billCode string
	// 运输企业RefEntId
	_transportRefEntId string
	// 运输企业名称
	_transportRefEntName string
	// 设备编号
	_equipmentCode string
	// 设备名称
	_equipmentName string
	// 温度信息
	_content string
	// 代上传的企业agent_ref_ent_id，特指“代上传温度”的企业。         如企业A上传了一个出库单，您代为A的单据上传运输温度，那么此时Agent_RefEntid即为您的ID。         若您本企业上传了一个单据，您本企业为这个单据上传温度，不存在“代上传”概念，此时agent_ref_ent_id即为空。
	_agentRefEntId string
}

// NewAlibabaalihealthdrugkytdrtransportuploadRequest 初始化AlibabaalihealthdrugkytdrtransportuploadAPIRequest对象
func NewAlibabaalihealthdrugkytdrtransportuploadRequest() *AlibabaalihealthdrugkytdrtransportuploadAPIRequest {
	return &AlibabaalihealthdrugkytdrtransportuploadAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthdrugkytdrtransportuploadAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.kyt.dr.transportupload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthdrugkytdrtransportuploadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthdrugkytdrtransportuploadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefEntId is RefEntId Setter
// 企业RefEntid，是指该单据的所有者。         如企业A上传了一个出库单，您为A的单据上传运输温度，那么此时RefEntid即为A的ID。         若您本企业上传了一个单据，您为这个单据上传温度，此时RefEntid即为您本企业的ID。
func (r *AlibabaalihealthdrugkytdrtransportuploadAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaalihealthdrugkytdrtransportuploadAPIRequest) GetRefEntId() string {
	return r._refEntId
}

// SetBillCode is BillCode Setter
// 单据编号
func (r *AlibabaalihealthdrugkytdrtransportuploadAPIRequest) SetBillCode(_billCode string) error {
	r._billCode = _billCode
	r.Set("bill_code", _billCode)
	return nil
}

// GetBillCode BillCode Getter
func (r AlibabaalihealthdrugkytdrtransportuploadAPIRequest) GetBillCode() string {
	return r._billCode
}

// SetTransportRefEntId is TransportRefEntId Setter
// 运输企业RefEntId
func (r *AlibabaalihealthdrugkytdrtransportuploadAPIRequest) SetTransportRefEntId(_transportRefEntId string) error {
	r._transportRefEntId = _transportRefEntId
	r.Set("transport_ref_ent_id", _transportRefEntId)
	return nil
}

// GetTransportRefEntId TransportRefEntId Getter
func (r AlibabaalihealthdrugkytdrtransportuploadAPIRequest) GetTransportRefEntId() string {
	return r._transportRefEntId
}

// SetTransportRefEntName is TransportRefEntName Setter
// 运输企业名称
func (r *AlibabaalihealthdrugkytdrtransportuploadAPIRequest) SetTransportRefEntName(_transportRefEntName string) error {
	r._transportRefEntName = _transportRefEntName
	r.Set("transport_ref_ent_name", _transportRefEntName)
	return nil
}

// GetTransportRefEntName TransportRefEntName Getter
func (r AlibabaalihealthdrugkytdrtransportuploadAPIRequest) GetTransportRefEntName() string {
	return r._transportRefEntName
}

// SetEquipmentCode is EquipmentCode Setter
// 设备编号
func (r *AlibabaalihealthdrugkytdrtransportuploadAPIRequest) SetEquipmentCode(_equipmentCode string) error {
	r._equipmentCode = _equipmentCode
	r.Set("equipment_code", _equipmentCode)
	return nil
}

// GetEquipmentCode EquipmentCode Getter
func (r AlibabaalihealthdrugkytdrtransportuploadAPIRequest) GetEquipmentCode() string {
	return r._equipmentCode
}

// SetEquipmentName is EquipmentName Setter
// 设备名称
func (r *AlibabaalihealthdrugkytdrtransportuploadAPIRequest) SetEquipmentName(_equipmentName string) error {
	r._equipmentName = _equipmentName
	r.Set("equipment_name", _equipmentName)
	return nil
}

// GetEquipmentName EquipmentName Getter
func (r AlibabaalihealthdrugkytdrtransportuploadAPIRequest) GetEquipmentName() string {
	return r._equipmentName
}

// SetContent is Content Setter
// 温度信息
func (r *AlibabaalihealthdrugkytdrtransportuploadAPIRequest) SetContent(_content string) error {
	r._content = _content
	r.Set("content", _content)
	return nil
}

// GetContent Content Getter
func (r AlibabaalihealthdrugkytdrtransportuploadAPIRequest) GetContent() string {
	return r._content
}

// SetAgentRefEntId is AgentRefEntId Setter
// 代上传的企业agent_ref_ent_id，特指“代上传温度”的企业。         如企业A上传了一个出库单，您代为A的单据上传运输温度，那么此时Agent_RefEntid即为您的ID。         若您本企业上传了一个单据，您本企业为这个单据上传温度，不存在“代上传”概念，此时agent_ref_ent_id即为空。
func (r *AlibabaalihealthdrugkytdrtransportuploadAPIRequest) SetAgentRefEntId(_agentRefEntId string) error {
	r._agentRefEntId = _agentRefEntId
	r.Set("agent_ref_ent_id", _agentRefEntId)
	return nil
}

// GetAgentRefEntId AgentRefEntId Getter
func (r AlibabaalihealthdrugkytdrtransportuploadAPIRequest) GetAgentRefEntId() string {
	return r._agentRefEntId
}
