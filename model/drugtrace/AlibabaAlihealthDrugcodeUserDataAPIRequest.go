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
type AlibabaAlihealthDrugcodeUserDataAPIRequest struct {
    model.Params
    // 用户信息
    _list   []HaoxinqingDataDto
    // 企业ID，用户区分 appkey下不同企业数据隔离的
    _refEntId   string
}

// 初始化AlibabaAlihealthDrugcodeUserDataAPIRequest对象
func NewAlibabaAlihealthDrugcodeUserDataRequest() *AlibabaAlihealthDrugcodeUserDataAPIRequest{
    return &AlibabaAlihealthDrugcodeUserDataAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugcodeUserDataAPIRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drugcode.user.data"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugcodeUserDataAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// List Setter
// 用户信息
func (r *AlibabaAlihealthDrugcodeUserDataAPIRequest) SetList(_list []HaoxinqingDataDto) error {
    r._list = _list
    r.Set("list", _list)
    return nil
}

// List Getter
func (r AlibabaAlihealthDrugcodeUserDataAPIRequest) GetList() []HaoxinqingDataDto {
    return r._list
}
// RefEntId Setter
// 企业ID，用户区分 appkey下不同企业数据隔离的
func (r *AlibabaAlihealthDrugcodeUserDataAPIRequest) SetRefEntId(_refEntId string) error {
    r._refEntId = _refEntId
    r.Set("ref_ent_id", _refEntId)
    return nil
}

// RefEntId Getter
func (r AlibabaAlihealthDrugcodeUserDataAPIRequest) GetRefEntId() string {
    return r._refEntId
}
