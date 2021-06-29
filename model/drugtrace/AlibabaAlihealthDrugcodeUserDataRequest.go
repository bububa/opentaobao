package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
西安杨森同步用户行为接口 API请求
alibaba.alihealth.drugcode.user.data

西安杨森同步用户行为接口
*/
type AlibabaAlihealthDrugcodeUserDataRequest struct {
    model.Params
    // 用户信息
    list   []HaoxinqingDataDto
    // 企业ID，用户区分 appkey下不同企业数据隔离的
    refEntId   string
}

// 初始化AlibabaAlihealthDrugcodeUserDataRequest对象
func NewAlibabaAlihealthDrugcodeUserDataRequest() *AlibabaAlihealthDrugcodeUserDataRequest{
    return &AlibabaAlihealthDrugcodeUserDataRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugcodeUserDataRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drugcode.user.data"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugcodeUserDataRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// List Setter
// 用户信息
func (r *AlibabaAlihealthDrugcodeUserDataRequest) SetList(list []HaoxinqingDataDto) error {
    r.list = list
    r.Set("list", list)
    return nil
}

// List Getter
func (r AlibabaAlihealthDrugcodeUserDataRequest) GetList() []HaoxinqingDataDto {
    return r.list
}
// RefEntId Setter
// 企业ID，用户区分 appkey下不同企业数据隔离的
func (r *AlibabaAlihealthDrugcodeUserDataRequest) SetRefEntId(refEntId string) error {
    r.refEntId = refEntId
    r.Set("ref_ent_id", refEntId)
    return nil
}

// RefEntId Getter
func (r AlibabaAlihealthDrugcodeUserDataRequest) GetRefEntId() string {
    return r.refEntId
}
