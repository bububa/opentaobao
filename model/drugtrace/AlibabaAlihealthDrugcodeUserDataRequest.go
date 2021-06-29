package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
西安杨森同步用户行为接口 APIRequest
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

func NewAlibabaAlihealthDrugcodeUserDataRequest() *AlibabaAlihealthDrugcodeUserDataRequest{
    return &AlibabaAlihealthDrugcodeUserDataRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthDrugcodeUserDataRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drugcode.user.data"
}

func (r AlibabaAlihealthDrugcodeUserDataRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthDrugcodeUserDataRequest) SetList(list []HaoxinqingDataDto) error {
    r.list = list
    r.Set("list", list)
    return nil
}

func (r AlibabaAlihealthDrugcodeUserDataRequest) GetList() []HaoxinqingDataDto {
    return r.list
}

func (r *AlibabaAlihealthDrugcodeUserDataRequest) SetRefEntId(refEntId string) error {
    r.refEntId = refEntId
    r.Set("ref_ent_id", refEntId)
    return nil
}

func (r AlibabaAlihealthDrugcodeUserDataRequest) GetRefEntId() string {
    return r.refEntId
}

