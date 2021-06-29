package vaccin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
疫苗预约地市信息查询 API请求
alibaba.alihealth.ms.area.district.list

微信小程序疫苗预约地市信息查询
*/
type AlibabaAlihealthMsAreaDistrictListRequest struct {
    model.Params
    // 省份ID
    _divisionId   int64
}

// 初始化AlibabaAlihealthMsAreaDistrictListRequest对象
func NewAlibabaAlihealthMsAreaDistrictListRequest() *AlibabaAlihealthMsAreaDistrictListRequest{
    return &AlibabaAlihealthMsAreaDistrictListRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthMsAreaDistrictListRequest) GetApiMethodName() string {
    return "alibaba.alihealth.ms.area.district.list"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthMsAreaDistrictListRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// DivisionId Setter
// 省份ID
func (r *AlibabaAlihealthMsAreaDistrictListRequest) SetDivisionId(_divisionId int64) error {
    r._divisionId = _divisionId
    r.Set("division_id", _divisionId)
    return nil
}

// DivisionId Getter
func (r AlibabaAlihealthMsAreaDistrictListRequest) GetDivisionId() int64 {
    return r._divisionId
}
