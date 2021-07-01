package examination

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
单项/加项包信息同步 API请求
alibaba.alihealth.examination.items.publish

体检机构对接_单项/加项包信息发布／更新
*/
type AlibabaAlihealthExaminationItemsPublishAPIRequest struct {
    model.Params
    // 商品id，机构保证全局唯一
    _groupId   string
    // 套餐列表
    _isvPackages   []IsvPackage
    // 单项之间关系
    _isvItemRelationDTOS   []IsvItemRelationDTO
    // 体检机构标识
    _hospitalCodes   []string
    // 加项包列表
    _isvItemPackDTOS   []IsvItemPackDTO
    // 单项信息列表
    _isvItemDTOS   []IsvItemDTO
    // 加项包关系列表
    _isvPackRelationDTOS   []IsvPackRelationDTO
}

// 初始化AlibabaAlihealthExaminationItemsPublishAPIRequest对象
func NewAlibabaAlihealthExaminationItemsPublishRequest() *AlibabaAlihealthExaminationItemsPublishAPIRequest{
    return &AlibabaAlihealthExaminationItemsPublishAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthExaminationItemsPublishAPIRequest) GetApiMethodName() string {
    return "alibaba.alihealth.examination.items.publish"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthExaminationItemsPublishAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// GroupId Setter
// 商品id，机构保证全局唯一
func (r *AlibabaAlihealthExaminationItemsPublishAPIRequest) SetGroupId(_groupId string) error {
    r._groupId = _groupId
    r.Set("group_id", _groupId)
    return nil
}

// GroupId Getter
func (r AlibabaAlihealthExaminationItemsPublishAPIRequest) GetGroupId() string {
    return r._groupId
}
// IsvPackages Setter
// 套餐列表
func (r *AlibabaAlihealthExaminationItemsPublishAPIRequest) SetIsvPackages(_isvPackages []IsvPackage) error {
    r._isvPackages = _isvPackages
    r.Set("isv_packages", _isvPackages)
    return nil
}

// IsvPackages Getter
func (r AlibabaAlihealthExaminationItemsPublishAPIRequest) GetIsvPackages() []IsvPackage {
    return r._isvPackages
}
// IsvItemRelationDTOS Setter
// 单项之间关系
func (r *AlibabaAlihealthExaminationItemsPublishAPIRequest) SetIsvItemRelationDTOS(_isvItemRelationDTOS []IsvItemRelationDTO) error {
    r._isvItemRelationDTOS = _isvItemRelationDTOS
    r.Set("isv_item_relation_d_t_o_s", _isvItemRelationDTOS)
    return nil
}

// IsvItemRelationDTOS Getter
func (r AlibabaAlihealthExaminationItemsPublishAPIRequest) GetIsvItemRelationDTOS() []IsvItemRelationDTO {
    return r._isvItemRelationDTOS
}
// HospitalCodes Setter
// 体检机构标识
func (r *AlibabaAlihealthExaminationItemsPublishAPIRequest) SetHospitalCodes(_hospitalCodes []string) error {
    r._hospitalCodes = _hospitalCodes
    r.Set("hospital_codes", _hospitalCodes)
    return nil
}

// HospitalCodes Getter
func (r AlibabaAlihealthExaminationItemsPublishAPIRequest) GetHospitalCodes() []string {
    return r._hospitalCodes
}
// IsvItemPackDTOS Setter
// 加项包列表
func (r *AlibabaAlihealthExaminationItemsPublishAPIRequest) SetIsvItemPackDTOS(_isvItemPackDTOS []IsvItemPackDTO) error {
    r._isvItemPackDTOS = _isvItemPackDTOS
    r.Set("isv_item_pack_d_t_o_s", _isvItemPackDTOS)
    return nil
}

// IsvItemPackDTOS Getter
func (r AlibabaAlihealthExaminationItemsPublishAPIRequest) GetIsvItemPackDTOS() []IsvItemPackDTO {
    return r._isvItemPackDTOS
}
// IsvItemDTOS Setter
// 单项信息列表
func (r *AlibabaAlihealthExaminationItemsPublishAPIRequest) SetIsvItemDTOS(_isvItemDTOS []IsvItemDTO) error {
    r._isvItemDTOS = _isvItemDTOS
    r.Set("isv_item_d_t_o_s", _isvItemDTOS)
    return nil
}

// IsvItemDTOS Getter
func (r AlibabaAlihealthExaminationItemsPublishAPIRequest) GetIsvItemDTOS() []IsvItemDTO {
    return r._isvItemDTOS
}
// IsvPackRelationDTOS Setter
// 加项包关系列表
func (r *AlibabaAlihealthExaminationItemsPublishAPIRequest) SetIsvPackRelationDTOS(_isvPackRelationDTOS []IsvPackRelationDTO) error {
    r._isvPackRelationDTOS = _isvPackRelationDTOS
    r.Set("isv_pack_relation_d_t_o_s", _isvPackRelationDTOS)
    return nil
}

// IsvPackRelationDTOS Getter
func (r AlibabaAlihealthExaminationItemsPublishAPIRequest) GetIsvPackRelationDTOS() []IsvPackRelationDTO {
    return r._isvPackRelationDTOS
}
