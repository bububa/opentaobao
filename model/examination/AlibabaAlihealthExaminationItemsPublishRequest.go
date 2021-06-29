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
type AlibabaAlihealthExaminationItemsPublishRequest struct {
    model.Params
    // 商品id，机构保证全局唯一
    groupId   string
    // 套餐列表
    isvPackages   []IsvPackage
    // 单项之间关系
    isvItemRelationDTOS   []IsvItemRelationDto
    // 体检机构标识
    hospitalCodes   []string
    // 加项包列表
    isvItemPackDTOS   []IsvItemPackDto
    // 单项信息列表
    isvItemDTOS   []IsvItemDto
    // 加项包关系列表
    isvPackRelationDTOS   []IsvPackRelationDto
}

// 初始化AlibabaAlihealthExaminationItemsPublishRequest对象
func NewAlibabaAlihealthExaminationItemsPublishRequest() *AlibabaAlihealthExaminationItemsPublishRequest{
    return &AlibabaAlihealthExaminationItemsPublishRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthExaminationItemsPublishRequest) GetApiMethodName() string {
    return "alibaba.alihealth.examination.items.publish"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthExaminationItemsPublishRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// GroupId Setter
// 商品id，机构保证全局唯一
func (r *AlibabaAlihealthExaminationItemsPublishRequest) SetGroupId(groupId string) error {
    r.groupId = groupId
    r.Set("group_id", groupId)
    return nil
}

// GroupId Getter
func (r AlibabaAlihealthExaminationItemsPublishRequest) GetGroupId() string {
    return r.groupId
}
// IsvPackages Setter
// 套餐列表
func (r *AlibabaAlihealthExaminationItemsPublishRequest) SetIsvPackages(isvPackages []IsvPackage) error {
    r.isvPackages = isvPackages
    r.Set("isv_packages", isvPackages)
    return nil
}

// IsvPackages Getter
func (r AlibabaAlihealthExaminationItemsPublishRequest) GetIsvPackages() []IsvPackage {
    return r.isvPackages
}
// IsvItemRelationDTOS Setter
// 单项之间关系
func (r *AlibabaAlihealthExaminationItemsPublishRequest) SetIsvItemRelationDTOS(isvItemRelationDTOS []IsvItemRelationDto) error {
    r.isvItemRelationDTOS = isvItemRelationDTOS
    r.Set("isv_item_relation_d_t_o_s", isvItemRelationDTOS)
    return nil
}

// IsvItemRelationDTOS Getter
func (r AlibabaAlihealthExaminationItemsPublishRequest) GetIsvItemRelationDTOS() []IsvItemRelationDto {
    return r.isvItemRelationDTOS
}
// HospitalCodes Setter
// 体检机构标识
func (r *AlibabaAlihealthExaminationItemsPublishRequest) SetHospitalCodes(hospitalCodes []string) error {
    r.hospitalCodes = hospitalCodes
    r.Set("hospital_codes", hospitalCodes)
    return nil
}

// HospitalCodes Getter
func (r AlibabaAlihealthExaminationItemsPublishRequest) GetHospitalCodes() []string {
    return r.hospitalCodes
}
// IsvItemPackDTOS Setter
// 加项包列表
func (r *AlibabaAlihealthExaminationItemsPublishRequest) SetIsvItemPackDTOS(isvItemPackDTOS []IsvItemPackDto) error {
    r.isvItemPackDTOS = isvItemPackDTOS
    r.Set("isv_item_pack_d_t_o_s", isvItemPackDTOS)
    return nil
}

// IsvItemPackDTOS Getter
func (r AlibabaAlihealthExaminationItemsPublishRequest) GetIsvItemPackDTOS() []IsvItemPackDto {
    return r.isvItemPackDTOS
}
// IsvItemDTOS Setter
// 单项信息列表
func (r *AlibabaAlihealthExaminationItemsPublishRequest) SetIsvItemDTOS(isvItemDTOS []IsvItemDto) error {
    r.isvItemDTOS = isvItemDTOS
    r.Set("isv_item_d_t_o_s", isvItemDTOS)
    return nil
}

// IsvItemDTOS Getter
func (r AlibabaAlihealthExaminationItemsPublishRequest) GetIsvItemDTOS() []IsvItemDto {
    return r.isvItemDTOS
}
// IsvPackRelationDTOS Setter
// 加项包关系列表
func (r *AlibabaAlihealthExaminationItemsPublishRequest) SetIsvPackRelationDTOS(isvPackRelationDTOS []IsvPackRelationDto) error {
    r.isvPackRelationDTOS = isvPackRelationDTOS
    r.Set("isv_pack_relation_d_t_o_s", isvPackRelationDTOS)
    return nil
}

// IsvPackRelationDTOS Getter
func (r AlibabaAlihealthExaminationItemsPublishRequest) GetIsvPackRelationDTOS() []IsvPackRelationDto {
    return r.isvPackRelationDTOS
}
