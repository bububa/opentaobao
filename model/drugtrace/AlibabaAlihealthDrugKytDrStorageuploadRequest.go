package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
疫苗存储温度上传 API请求
alibaba.alihealth.drug.kyt.dr.storageupload

疫苗存储温度上传
*/
type AlibabaAlihealthDrugKytDrStorageuploadRequest struct {
    model.Params
    // 企业RefEntid，是指该单据的所有者。         如企业A上传了一个出库单，您为A的单据上传运输温度，那么此时RefEntid即为A的ID。         若您本企业上传了一个单据，您为这个单据上传温度，此时RefEntid即为您本企业的ID。
    _refEntId   string
    // 设备编号
    _equipmentCode   string
    // 设备名称
    _equipmentName   string
    // 温度信息
    _content   string
    // 代上传的企业agent_ref_ent_id，特指“代上传温度”的企业。         如企业A上传了一个出库单，您代为A的单据上传运输温度，那么此时Agent_RefEntid即为您的ID。         若您本企业上传了一个单据，您本企业为这个单据上传温度，不存在“代上传”概念，此时agent_ref_ent_id即为空。
    _agentRefEntId   string
}

// 初始化AlibabaAlihealthDrugKytDrStorageuploadRequest对象
func NewAlibabaAlihealthDrugKytDrStorageuploadRequest() *AlibabaAlihealthDrugKytDrStorageuploadRequest{
    return &AlibabaAlihealthDrugKytDrStorageuploadRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytDrStorageuploadRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.kyt.dr.storageupload"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytDrStorageuploadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RefEntId Setter
// 企业RefEntid，是指该单据的所有者。         如企业A上传了一个出库单，您为A的单据上传运输温度，那么此时RefEntid即为A的ID。         若您本企业上传了一个单据，您为这个单据上传温度，此时RefEntid即为您本企业的ID。
func (r *AlibabaAlihealthDrugKytDrStorageuploadRequest) SetRefEntId(_refEntId string) error {
    r._refEntId = _refEntId
    r.Set("ref_ent_id", _refEntId)
    return nil
}

// RefEntId Getter
func (r AlibabaAlihealthDrugKytDrStorageuploadRequest) GetRefEntId() string {
    return r._refEntId
}
// EquipmentCode Setter
// 设备编号
func (r *AlibabaAlihealthDrugKytDrStorageuploadRequest) SetEquipmentCode(_equipmentCode string) error {
    r._equipmentCode = _equipmentCode
    r.Set("equipment_code", _equipmentCode)
    return nil
}

// EquipmentCode Getter
func (r AlibabaAlihealthDrugKytDrStorageuploadRequest) GetEquipmentCode() string {
    return r._equipmentCode
}
// EquipmentName Setter
// 设备名称
func (r *AlibabaAlihealthDrugKytDrStorageuploadRequest) SetEquipmentName(_equipmentName string) error {
    r._equipmentName = _equipmentName
    r.Set("equipment_name", _equipmentName)
    return nil
}

// EquipmentName Getter
func (r AlibabaAlihealthDrugKytDrStorageuploadRequest) GetEquipmentName() string {
    return r._equipmentName
}
// Content Setter
// 温度信息
func (r *AlibabaAlihealthDrugKytDrStorageuploadRequest) SetContent(_content string) error {
    r._content = _content
    r.Set("content", _content)
    return nil
}

// Content Getter
func (r AlibabaAlihealthDrugKytDrStorageuploadRequest) GetContent() string {
    return r._content
}
// AgentRefEntId Setter
// 代上传的企业agent_ref_ent_id，特指“代上传温度”的企业。         如企业A上传了一个出库单，您代为A的单据上传运输温度，那么此时Agent_RefEntid即为您的ID。         若您本企业上传了一个单据，您本企业为这个单据上传温度，不存在“代上传”概念，此时agent_ref_ent_id即为空。
func (r *AlibabaAlihealthDrugKytDrStorageuploadRequest) SetAgentRefEntId(_agentRefEntId string) error {
    r._agentRefEntId = _agentRefEntId
    r.Set("agent_ref_ent_id", _agentRefEntId)
    return nil
}

// AgentRefEntId Getter
func (r AlibabaAlihealthDrugKytDrStorageuploadRequest) GetAgentRefEntId() string {
    return r._agentRefEntId
}
