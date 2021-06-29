package vaccin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
疫苗预约地市信息查询 APIRequest
alibaba.alihealth.ms.area.district.list

微信小程序疫苗预约地市信息查询
*/
type AlibabaAlihealthMsAreaDistrictListRequest struct {
    model.Params

    // 省份ID
    divisionId   int64 

}

func NewAlibabaAlihealthMsAreaDistrictListRequest() *AlibabaAlihealthMsAreaDistrictListRequest{
    return &AlibabaAlihealthMsAreaDistrictListRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthMsAreaDistrictListRequest) GetApiMethodName() string {
    return "alibaba.alihealth.ms.area.district.list"
}

func (r AlibabaAlihealthMsAreaDistrictListRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthMsAreaDistrictListRequest) SetDivisionId(divisionId int64) error {
    r.divisionId = divisionId
    r.Set("division_id", divisionId)
    return nil
}

func (r AlibabaAlihealthMsAreaDistrictListRequest) GetDivisionId() int64 {
    return r.divisionId
}

