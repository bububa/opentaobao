package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
疫苗运输温度上传 API请求
alibaba.alihealth.drug.kyt.dr.transportupload

疫苗运输温度上传
*/
type AlibabaAlihealthDrugKytDrTransportuploadRequest struct {
    model.Params
    // 企业RefEntid，是指该单据的所有者。         如企业A上传了一个出库单，您为A的单据上传运输温度，那么此时RefEntid即为A的ID。         若您本企业上传了一个单据，您为这个单据上传温度，此时RefEntid即为您本企业的ID。
    _refEntId   string
    // 单据编号
    _billCode   string
    // 运输企业RefEntId
    _transportRefEntId   string
    // 运输企业名称
    _transportRefEntName   string
    // 设备编号
    _equipmentCode   string
    // 设备名称
    _equipmentName   string
    // 温度信息
    _content   string
    // 代上传的企业agent_ref_ent_id，特指“代上传温度”的企业。         如企业A上传了一个出库单，您代为A的单据上传运输温度，那么此时Agent_RefEntid即为您的ID。         若您本企业上传了一个单据，您本企业为这个单据上传温度，不存在“代上传”概念，此时agent_ref_ent_id即为空。
    _agentRefEntId   string
}

// 初始化AlibabaAlihealthDrugKytDrTransportuploadRequest对象
func NewAlibabaAlihealthDrugKytDrTransportuploadRequest() *AlibabaAlihealthDrugKytDrTransportuploadRequest{
    return &AlibabaAlihealthDrugKytDrTransportuploadRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytDrTransportuploadRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.kyt.dr.transportupload"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytDrTransportuploadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RefEntId Setter
// 企业RefEntid，是指该单据的所有者。         如企业A上传了一个出库单，您为A的单据上传运输温度，那么此时RefEntid即为A的ID。         若您本企业上传了一个单据，您为这个单据上传温度，此时RefEntid即为您本企业的ID。
func (r *AlibabaAlihealthDrugKytDrTransportuploadRequest) SetRefEntId(_refEntId string) error {
    r._refEntId = _refEntId
    r.Set("ref_ent_id", _refEntId)
    return nil
}

// RefEntId Getter
func (r AlibabaAlihealthDrugKytDrTransportuploadRequest) GetRefEntId() string {
    return r._refEntId
}
// BillCode Setter
// 单据编号
func (r *AlibabaAlihealthDrugKytDrTransportuploadRequest) SetBillCode(_billCode string) error {
    r._billCode = _billCode
    r.Set("bill_code", _billCode)
    return nil
}

// BillCode Getter
func (r AlibabaAlihealthDrugKytDrTransportuploadRequest) GetBillCode() string {
    return r._billCode
}
// TransportRefEntId Setter
// 运输企业RefEntId
func (r *AlibabaAlihealthDrugKytDrTransportuploadRequest) SetTransportRefEntId(_transportRefEntId string) error {
    r._transportRefEntId = _transportRefEntId
    r.Set("transport_ref_ent_id", _transportRefEntId)
    return nil
}

// TransportRefEntId Getter
func (r AlibabaAlihealthDrugKytDrTransportuploadRequest) GetTransportRefEntId() string {
    return r._transportRefEntId
}
// TransportRefEntName Setter
// 运输企业名称
func (r *AlibabaAlihealthDrugKytDrTransportuploadRequest) SetTransportRefEntName(_transportRefEntName string) error {
    r._transportRefEntName = _transportRefEntName
    r.Set("transport_ref_ent_name", _transportRefEntName)
    return nil
}

// TransportRefEntName Getter
func (r AlibabaAlihealthDrugKytDrTransportuploadRequest) GetTransportRefEntName() string {
    return r._transportRefEntName
}
// EquipmentCode Setter
// 设备编号
func (r *AlibabaAlihealthDrugKytDrTransportuploadRequest) SetEquipmentCode(_equipmentCode string) error {
    r._equipmentCode = _equipmentCode
    r.Set("equipment_code", _equipmentCode)
    return nil
}

// EquipmentCode Getter
func (r AlibabaAlihealthDrugKytDrTransportuploadRequest) GetEquipmentCode() string {
    return r._equipmentCode
}
// EquipmentName Setter
// 设备名称
func (r *AlibabaAlihealthDrugKytDrTransportuploadRequest) SetEquipmentName(_equipmentName string) error {
    r._equipmentName = _equipmentName
    r.Set("equipment_name", _equipmentName)
    return nil
}

// EquipmentName Getter
func (r AlibabaAlihealthDrugKytDrTransportuploadRequest) GetEquipmentName() string {
    return r._equipmentName
}
// Content Setter
// 温度信息
func (r *AlibabaAlihealthDrugKytDrTransportuploadRequest) SetContent(_content string) error {
    r._content = _content
    r.Set("content", _content)
    return nil
}

// Content Getter
func (r AlibabaAlihealthDrugKytDrTransportuploadRequest) GetContent() string {
    return r._content
}
// AgentRefEntId Setter
// 代上传的企业agent_ref_ent_id，特指“代上传温度”的企业。         如企业A上传了一个出库单，您代为A的单据上传运输温度，那么此时Agent_RefEntid即为您的ID。         若您本企业上传了一个单据，您本企业为这个单据上传温度，不存在“代上传”概念，此时agent_ref_ent_id即为空。
func (r *AlibabaAlihealthDrugKytDrTransportuploadRequest) SetAgentRefEntId(_agentRefEntId string) error {
    r._agentRefEntId = _agentRefEntId
    r.Set("agent_ref_ent_id", _agentRefEntId)
    return nil
}

// AgentRefEntId Getter
func (r AlibabaAlihealthDrugKytDrTransportuploadRequest) GetAgentRefEntId() string {
    return r._agentRefEntId
}
