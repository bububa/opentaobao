package examination

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
上线/下线 体检产品 API请求
alibaba.alihealth.examination.goods.onoffline

第三方体检机构对接钉钉体检中的产品 上线／下线
*/
type AlibabaAlihealthExaminationGoodsOnofflineRequest struct {
    model.Params
    // 商品组code，机构保证唯一
    groupId   string
    // 操作类型: online=上线，offline=下线
    type   string
    // 门店code列表
    hospitalCodes   string
}

// 初始化AlibabaAlihealthExaminationGoodsOnofflineRequest对象
func NewAlibabaAlihealthExaminationGoodsOnofflineRequest() *AlibabaAlihealthExaminationGoodsOnofflineRequest{
    return &AlibabaAlihealthExaminationGoodsOnofflineRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthExaminationGoodsOnofflineRequest) GetApiMethodName() string {
    return "alibaba.alihealth.examination.goods.onoffline"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthExaminationGoodsOnofflineRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// GroupId Setter
// 商品组code，机构保证唯一
func (r *AlibabaAlihealthExaminationGoodsOnofflineRequest) SetGroupId(groupId string) error {
    r.groupId = groupId
    r.Set("group_id", groupId)
    return nil
}

// GroupId Getter
func (r AlibabaAlihealthExaminationGoodsOnofflineRequest) GetGroupId() string {
    return r.groupId
}
// Type Setter
// 操作类型: online=上线，offline=下线
func (r *AlibabaAlihealthExaminationGoodsOnofflineRequest) SetType(type string) error {
    r.type = type
    r.Set("type", type)
    return nil
}

// Type Getter
func (r AlibabaAlihealthExaminationGoodsOnofflineRequest) GetType() string {
    return r.type
}
// HospitalCodes Setter
// 门店code列表
func (r *AlibabaAlihealthExaminationGoodsOnofflineRequest) SetHospitalCodes(hospitalCodes string) error {
    r.hospitalCodes = hospitalCodes
    r.Set("hospital_codes", hospitalCodes)
    return nil
}

// HospitalCodes Getter
func (r AlibabaAlihealthExaminationGoodsOnofflineRequest) GetHospitalCodes() string {
    return r.hospitalCodes
}
