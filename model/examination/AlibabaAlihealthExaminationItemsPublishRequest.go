package examination

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
单项/加项包信息同步 APIRequest
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

func NewAlibabaAlihealthExaminationItemsPublishRequest() *AlibabaAlihealthExaminationItemsPublishRequest{
    return &AlibabaAlihealthExaminationItemsPublishRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthExaminationItemsPublishRequest) GetApiMethodName() string {
    return "alibaba.alihealth.examination.items.publish"
}

func (r AlibabaAlihealthExaminationItemsPublishRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthExaminationItemsPublishRequest) SetGroupId(groupId string) error {
    r.groupId = groupId
    r.Set("group_id", groupId)
    return nil
}

func (r AlibabaAlihealthExaminationItemsPublishRequest) GetGroupId() string {
    return r.groupId
}

func (r *AlibabaAlihealthExaminationItemsPublishRequest) SetIsvPackages(isvPackages []IsvPackage) error {
    r.isvPackages = isvPackages
    r.Set("isv_packages", isvPackages)
    return nil
}

func (r AlibabaAlihealthExaminationItemsPublishRequest) GetIsvPackages() []IsvPackage {
    return r.isvPackages
}

func (r *AlibabaAlihealthExaminationItemsPublishRequest) SetIsvItemRelationDTOS(isvItemRelationDTOS []IsvItemRelationDto) error {
    r.isvItemRelationDTOS = isvItemRelationDTOS
    r.Set("isv_item_relation_d_t_o_s", isvItemRelationDTOS)
    return nil
}

func (r AlibabaAlihealthExaminationItemsPublishRequest) GetIsvItemRelationDTOS() []IsvItemRelationDto {
    return r.isvItemRelationDTOS
}

func (r *AlibabaAlihealthExaminationItemsPublishRequest) SetHospitalCodes(hospitalCodes []string) error {
    r.hospitalCodes = hospitalCodes
    r.Set("hospital_codes", hospitalCodes)
    return nil
}

func (r AlibabaAlihealthExaminationItemsPublishRequest) GetHospitalCodes() []string {
    return r.hospitalCodes
}

func (r *AlibabaAlihealthExaminationItemsPublishRequest) SetIsvItemPackDTOS(isvItemPackDTOS []IsvItemPackDto) error {
    r.isvItemPackDTOS = isvItemPackDTOS
    r.Set("isv_item_pack_d_t_o_s", isvItemPackDTOS)
    return nil
}

func (r AlibabaAlihealthExaminationItemsPublishRequest) GetIsvItemPackDTOS() []IsvItemPackDto {
    return r.isvItemPackDTOS
}

func (r *AlibabaAlihealthExaminationItemsPublishRequest) SetIsvItemDTOS(isvItemDTOS []IsvItemDto) error {
    r.isvItemDTOS = isvItemDTOS
    r.Set("isv_item_d_t_o_s", isvItemDTOS)
    return nil
}

func (r AlibabaAlihealthExaminationItemsPublishRequest) GetIsvItemDTOS() []IsvItemDto {
    return r.isvItemDTOS
}

func (r *AlibabaAlihealthExaminationItemsPublishRequest) SetIsvPackRelationDTOS(isvPackRelationDTOS []IsvPackRelationDto) error {
    r.isvPackRelationDTOS = isvPackRelationDTOS
    r.Set("isv_pack_relation_d_t_o_s", isvPackRelationDTOS)
    return nil
}

func (r AlibabaAlihealthExaminationItemsPublishRequest) GetIsvPackRelationDTOS() []IsvPackRelationDto {
    return r.isvPackRelationDTOS
}

