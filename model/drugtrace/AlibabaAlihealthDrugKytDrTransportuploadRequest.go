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
    refEntId   string
    // 单据编号
    billCode   string
    // 运输企业RefEntId
    transportRefEntId   string
    // 运输企业名称
    transportRefEntName   string
    // 设备编号
    equipmentCode   string
    // 设备名称
    equipmentName   string
    // 温度信息
    content   string
    // 代上传的企业agent_ref_ent_id，特指“代上传温度”的企业。         如企业A上传了一个出库单，您代为A的单据上传运输温度，那么此时Agent_RefEntid即为您的ID。         若您本企业上传了一个单据，您本企业为这个单据上传温度，不存在“代上传”概念，此时agent_ref_ent_id即为空。
    agentRefEntId   string
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
func (r *AlibabaAlihealthDrugKytDrTransportuploadRequest) SetRefEntId(refEntId string) error {
    r.refEntId = refEntId
    r.Set("ref_ent_id", refEntId)
    return nil
}

// RefEntId Getter
func (r AlibabaAlihealthDrugKytDrTransportuploadRequest) GetRefEntId() string {
    return r.refEntId
}
// BillCode Setter
// 单据编号
func (r *AlibabaAlihealthDrugKytDrTransportuploadRequest) SetBillCode(billCode string) error {
    r.billCode = billCode
    r.Set("bill_code", billCode)
    return nil
}

// BillCode Getter
func (r AlibabaAlihealthDrugKytDrTransportuploadRequest) GetBillCode() string {
    return r.billCode
}
// TransportRefEntId Setter
// 运输企业RefEntId
func (r *AlibabaAlihealthDrugKytDrTransportuploadRequest) SetTransportRefEntId(transportRefEntId string) error {
    r.transportRefEntId = transportRefEntId
    r.Set("transport_ref_ent_id", transportRefEntId)
    return nil
}

// TransportRefEntId Getter
func (r AlibabaAlihealthDrugKytDrTransportuploadRequest) GetTransportRefEntId() string {
    return r.transportRefEntId
}
// TransportRefEntName Setter
// 运输企业名称
func (r *AlibabaAlihealthDrugKytDrTransportuploadRequest) SetTransportRefEntName(transportRefEntName string) error {
    r.transportRefEntName = transportRefEntName
    r.Set("transport_ref_ent_name", transportRefEntName)
    return nil
}

// TransportRefEntName Getter
func (r AlibabaAlihealthDrugKytDrTransportuploadRequest) GetTransportRefEntName() string {
    return r.transportRefEntName
}
// EquipmentCode Setter
// 设备编号
func (r *AlibabaAlihealthDrugKytDrTransportuploadRequest) SetEquipmentCode(equipmentCode string) error {
    r.equipmentCode = equipmentCode
    r.Set("equipment_code", equipmentCode)
    return nil
}

// EquipmentCode Getter
func (r AlibabaAlihealthDrugKytDrTransportuploadRequest) GetEquipmentCode() string {
    return r.equipmentCode
}
// EquipmentName Setter
// 设备名称
func (r *AlibabaAlihealthDrugKytDrTransportuploadRequest) SetEquipmentName(equipmentName string) error {
    r.equipmentName = equipmentName
    r.Set("equipment_name", equipmentName)
    return nil
}

// EquipmentName Getter
func (r AlibabaAlihealthDrugKytDrTransportuploadRequest) GetEquipmentName() string {
    return r.equipmentName
}
// Content Setter
// 温度信息
func (r *AlibabaAlihealthDrugKytDrTransportuploadRequest) SetContent(content string) error {
    r.content = content
    r.Set("content", content)
    return nil
}

// Content Getter
func (r AlibabaAlihealthDrugKytDrTransportuploadRequest) GetContent() string {
    return r.content
}
// AgentRefEntId Setter
// 代上传的企业agent_ref_ent_id，特指“代上传温度”的企业。         如企业A上传了一个出库单，您代为A的单据上传运输温度，那么此时Agent_RefEntid即为您的ID。         若您本企业上传了一个单据，您本企业为这个单据上传温度，不存在“代上传”概念，此时agent_ref_ent_id即为空。
func (r *AlibabaAlihealthDrugKytDrTransportuploadRequest) SetAgentRefEntId(agentRefEntId string) error {
    r.agentRefEntId = agentRefEntId
    r.Set("agent_ref_ent_id", agentRefEntId)
    return nil
}

// AgentRefEntId Getter
func (r AlibabaAlihealthDrugKytDrTransportuploadRequest) GetAgentRefEntId() string {
    return r.agentRefEntId
}
