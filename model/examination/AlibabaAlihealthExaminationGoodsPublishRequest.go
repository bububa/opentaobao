package examination

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
体检机构对接_商品发布／更新 APIRequest
alibaba.alihealth.examination.goods.publish

体检机构对接_商品发布／更新
*/
type AlibabaAlihealthExaminationGoodsPublishRequest struct {
    model.Params

    // 商品id，机构保证全局唯一
    groupId   string 

    // 商品名称
    groupName   string 

    // 套餐列表
    packageList   []Package 

    // 操作类型: publish=发布，update=更新
    type   string 

    // 最多200个字，界面对应商品详情页描述
    goodsDesc   string 

    // 最多256个字，界面对应列表文字
    targetGroup   string 

    // 联调中正式上线前标签给B；联调后正式上线后标签给C
    label   string 

    // 商品类目，1：体检 ，2：核酸，4 ：健康证
    categoryId   string 

    // 0自营商品，1平台商品
    mode   string 

    // 类目ID，填入叶子类目ID，儿童体检: 20210204000004, 中青年体检: 20210204000005, 老年体检: 20210204000006, 证件体检（含入职）: 20210204000007, 核酸检测（到店服务）: 20210204000008, 专科服务（不包含核酸检测）: 20210204000009, 上门检测: 202102040000010, 上门护理: 202102040000011, 上门体检 202102040000012
    backendCategoryId   int64 

}

func NewAlibabaAlihealthExaminationGoodsPublishRequest() *AlibabaAlihealthExaminationGoodsPublishRequest{
    return &AlibabaAlihealthExaminationGoodsPublishRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthExaminationGoodsPublishRequest) GetApiMethodName() string {
    return "alibaba.alihealth.examination.goods.publish"
}

func (r AlibabaAlihealthExaminationGoodsPublishRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthExaminationGoodsPublishRequest) SetGroupId(groupId string) error {
    r.groupId = groupId
    r.Set("group_id", groupId)
    return nil
}

func (r AlibabaAlihealthExaminationGoodsPublishRequest) GetGroupId() string {
    return r.groupId
}

func (r *AlibabaAlihealthExaminationGoodsPublishRequest) SetGroupName(groupName string) error {
    r.groupName = groupName
    r.Set("group_name", groupName)
    return nil
}

func (r AlibabaAlihealthExaminationGoodsPublishRequest) GetGroupName() string {
    return r.groupName
}

func (r *AlibabaAlihealthExaminationGoodsPublishRequest) SetPackageList(packageList []Package) error {
    r.packageList = packageList
    r.Set("package_list", packageList)
    return nil
}

func (r AlibabaAlihealthExaminationGoodsPublishRequest) GetPackageList() []Package {
    return r.packageList
}

func (r *AlibabaAlihealthExaminationGoodsPublishRequest) SetType(type string) error {
    r.type = type
    r.Set("type", type)
    return nil
}

func (r AlibabaAlihealthExaminationGoodsPublishRequest) GetType() string {
    return r.type
}

func (r *AlibabaAlihealthExaminationGoodsPublishRequest) SetGoodsDesc(goodsDesc string) error {
    r.goodsDesc = goodsDesc
    r.Set("goods_desc", goodsDesc)
    return nil
}

func (r AlibabaAlihealthExaminationGoodsPublishRequest) GetGoodsDesc() string {
    return r.goodsDesc
}

func (r *AlibabaAlihealthExaminationGoodsPublishRequest) SetTargetGroup(targetGroup string) error {
    r.targetGroup = targetGroup
    r.Set("target_group", targetGroup)
    return nil
}

func (r AlibabaAlihealthExaminationGoodsPublishRequest) GetTargetGroup() string {
    return r.targetGroup
}

func (r *AlibabaAlihealthExaminationGoodsPublishRequest) SetLabel(label string) error {
    r.label = label
    r.Set("label", label)
    return nil
}

func (r AlibabaAlihealthExaminationGoodsPublishRequest) GetLabel() string {
    return r.label
}

func (r *AlibabaAlihealthExaminationGoodsPublishRequest) SetCategoryId(categoryId string) error {
    r.categoryId = categoryId
    r.Set("category_id", categoryId)
    return nil
}

func (r AlibabaAlihealthExaminationGoodsPublishRequest) GetCategoryId() string {
    return r.categoryId
}

func (r *AlibabaAlihealthExaminationGoodsPublishRequest) SetMode(mode string) error {
    r.mode = mode
    r.Set("mode", mode)
    return nil
}

func (r AlibabaAlihealthExaminationGoodsPublishRequest) GetMode() string {
    return r.mode
}

func (r *AlibabaAlihealthExaminationGoodsPublishRequest) SetBackendCategoryId(backendCategoryId int64) error {
    r.backendCategoryId = backendCategoryId
    r.Set("backend_category_id", backendCategoryId)
    return nil
}

func (r AlibabaAlihealthExaminationGoodsPublishRequest) GetBackendCategoryId() int64 {
    return r.backendCategoryId
}

