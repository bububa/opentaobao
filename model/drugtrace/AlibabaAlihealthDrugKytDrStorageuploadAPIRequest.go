package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthdrugkytdrstorageuploadAPIRequest 疫苗存储温度上传 API请求
// alibaba.alihealth.drug.kyt.dr.storageupload
//
// 疫苗存储温度上传
type AlibabaalihealthdrugkytdrstorageuploadAPIRequest struct {
	model.Params
	// 企业RefEntid，是指该单据的所有者。         如企业A上传了一个出库单，您为A的单据上传运输温度，那么此时RefEntid即为A的ID。         若您本企业上传了一个单据，您为这个单据上传温度，此时RefEntid即为您本企业的ID。
	_refEntId string
	// 设备编号
	_equipmentCode string
	// 设备名称
	_equipmentName string
	// 温度信息
	_content string
	// 代上传的企业agent_ref_ent_id，特指“代上传温度”的企业。         如企业A上传了一个出库单，您代为A的单据上传运输温度，那么此时Agent_RefEntid即为您的ID。         若您本企业上传了一个单据，您本企业为这个单据上传温度，不存在“代上传”概念，此时agent_ref_ent_id即为空。
	_agentRefEntId string
}

// NewAlibabaalihealthdrugkytdrstorageuploadRequest 初始化AlibabaalihealthdrugkytdrstorageuploadAPIRequest对象
func NewAlibabaalihealthdrugkytdrstorageuploadRequest() *AlibabaalihealthdrugkytdrstorageuploadAPIRequest {
	return &AlibabaalihealthdrugkytdrstorageuploadAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthdrugkytdrstorageuploadAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.kyt.dr.storageupload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthdrugkytdrstorageuploadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthdrugkytdrstorageuploadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefEntId is RefEntId Setter
// 企业RefEntid，是指该单据的所有者。         如企业A上传了一个出库单，您为A的单据上传运输温度，那么此时RefEntid即为A的ID。         若您本企业上传了一个单据，您为这个单据上传温度，此时RefEntid即为您本企业的ID。
func (r *AlibabaalihealthdrugkytdrstorageuploadAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaalihealthdrugkytdrstorageuploadAPIRequest) GetRefEntId() string {
	return r._refEntId
}

// SetEquipmentCode is EquipmentCode Setter
// 设备编号
func (r *AlibabaalihealthdrugkytdrstorageuploadAPIRequest) SetEquipmentCode(_equipmentCode string) error {
	r._equipmentCode = _equipmentCode
	r.Set("equipment_code", _equipmentCode)
	return nil
}

// GetEquipmentCode EquipmentCode Getter
func (r AlibabaalihealthdrugkytdrstorageuploadAPIRequest) GetEquipmentCode() string {
	return r._equipmentCode
}

// SetEquipmentName is EquipmentName Setter
// 设备名称
func (r *AlibabaalihealthdrugkytdrstorageuploadAPIRequest) SetEquipmentName(_equipmentName string) error {
	r._equipmentName = _equipmentName
	r.Set("equipment_name", _equipmentName)
	return nil
}

// GetEquipmentName EquipmentName Getter
func (r AlibabaalihealthdrugkytdrstorageuploadAPIRequest) GetEquipmentName() string {
	return r._equipmentName
}

// SetContent is Content Setter
// 温度信息
func (r *AlibabaalihealthdrugkytdrstorageuploadAPIRequest) SetContent(_content string) error {
	r._content = _content
	r.Set("content", _content)
	return nil
}

// GetContent Content Getter
func (r AlibabaalihealthdrugkytdrstorageuploadAPIRequest) GetContent() string {
	return r._content
}

// SetAgentRefEntId is AgentRefEntId Setter
// 代上传的企业agent_ref_ent_id，特指“代上传温度”的企业。         如企业A上传了一个出库单，您代为A的单据上传运输温度，那么此时Agent_RefEntid即为您的ID。         若您本企业上传了一个单据，您本企业为这个单据上传温度，不存在“代上传”概念，此时agent_ref_ent_id即为空。
func (r *AlibabaalihealthdrugkytdrstorageuploadAPIRequest) SetAgentRefEntId(_agentRefEntId string) error {
	r._agentRefEntId = _agentRefEntId
	r.Set("agent_ref_ent_id", _agentRefEntId)
	return nil
}

// GetAgentRefEntId AgentRefEntId Getter
func (r AlibabaalihealthdrugkytdrstorageuploadAPIRequest) GetAgentRefEntId() string {
	return r._agentRefEntId
}
