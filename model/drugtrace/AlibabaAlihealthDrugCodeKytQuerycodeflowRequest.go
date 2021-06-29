package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
码流向查询 API请求
alibaba.alihealth.drug.code.kyt.querycodeflow

追溯码流向查询
*/
type AlibabaAlihealthDrugCodeKytQuerycodeflowRequest struct {
    model.Params
    // 企业ID
    _refEntId   string
    // 追溯码
    _code   string
    // 经度
    _longitude   string
    // 纬度
    _latitude   string
    // 查询地所在省
    _queryProv   string
    // 查询地所在市
    _queryCity   string
    // 查询地所在区
    _queryArea   string
    // 查询地所在区域代码
    _queryRegionCode   string
    // 详细地址
    _detail   string
}

// 初始化AlibabaAlihealthDrugCodeKytQuerycodeflowRequest对象
func NewAlibabaAlihealthDrugCodeKytQuerycodeflowRequest() *AlibabaAlihealthDrugCodeKytQuerycodeflowRequest{
    return &AlibabaAlihealthDrugCodeKytQuerycodeflowRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugCodeKytQuerycodeflowRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.code.kyt.querycodeflow"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugCodeKytQuerycodeflowRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RefEntId Setter
// 企业ID
func (r *AlibabaAlihealthDrugCodeKytQuerycodeflowRequest) SetRefEntId(_refEntId string) error {
    r._refEntId = _refEntId
    r.Set("ref_ent_id", _refEntId)
    return nil
}

// RefEntId Getter
func (r AlibabaAlihealthDrugCodeKytQuerycodeflowRequest) GetRefEntId() string {
    return r._refEntId
}
// Code Setter
// 追溯码
func (r *AlibabaAlihealthDrugCodeKytQuerycodeflowRequest) SetCode(_code string) error {
    r._code = _code
    r.Set("code", _code)
    return nil
}

// Code Getter
func (r AlibabaAlihealthDrugCodeKytQuerycodeflowRequest) GetCode() string {
    return r._code
}
// Longitude Setter
// 经度
func (r *AlibabaAlihealthDrugCodeKytQuerycodeflowRequest) SetLongitude(_longitude string) error {
    r._longitude = _longitude
    r.Set("longitude", _longitude)
    return nil
}

// Longitude Getter
func (r AlibabaAlihealthDrugCodeKytQuerycodeflowRequest) GetLongitude() string {
    return r._longitude
}
// Latitude Setter
// 纬度
func (r *AlibabaAlihealthDrugCodeKytQuerycodeflowRequest) SetLatitude(_latitude string) error {
    r._latitude = _latitude
    r.Set("latitude", _latitude)
    return nil
}

// Latitude Getter
func (r AlibabaAlihealthDrugCodeKytQuerycodeflowRequest) GetLatitude() string {
    return r._latitude
}
// QueryProv Setter
// 查询地所在省
func (r *AlibabaAlihealthDrugCodeKytQuerycodeflowRequest) SetQueryProv(_queryProv string) error {
    r._queryProv = _queryProv
    r.Set("query_prov", _queryProv)
    return nil
}

// QueryProv Getter
func (r AlibabaAlihealthDrugCodeKytQuerycodeflowRequest) GetQueryProv() string {
    return r._queryProv
}
// QueryCity Setter
// 查询地所在市
func (r *AlibabaAlihealthDrugCodeKytQuerycodeflowRequest) SetQueryCity(_queryCity string) error {
    r._queryCity = _queryCity
    r.Set("query_city", _queryCity)
    return nil
}

// QueryCity Getter
func (r AlibabaAlihealthDrugCodeKytQuerycodeflowRequest) GetQueryCity() string {
    return r._queryCity
}
// QueryArea Setter
// 查询地所在区
func (r *AlibabaAlihealthDrugCodeKytQuerycodeflowRequest) SetQueryArea(_queryArea string) error {
    r._queryArea = _queryArea
    r.Set("query_area", _queryArea)
    return nil
}

// QueryArea Getter
func (r AlibabaAlihealthDrugCodeKytQuerycodeflowRequest) GetQueryArea() string {
    return r._queryArea
}
// QueryRegionCode Setter
// 查询地所在区域代码
func (r *AlibabaAlihealthDrugCodeKytQuerycodeflowRequest) SetQueryRegionCode(_queryRegionCode string) error {
    r._queryRegionCode = _queryRegionCode
    r.Set("query_region_code", _queryRegionCode)
    return nil
}

// QueryRegionCode Getter
func (r AlibabaAlihealthDrugCodeKytQuerycodeflowRequest) GetQueryRegionCode() string {
    return r._queryRegionCode
}
// Detail Setter
// 详细地址
func (r *AlibabaAlihealthDrugCodeKytQuerycodeflowRequest) SetDetail(_detail string) error {
    r._detail = _detail
    r.Set("detail", _detail)
    return nil
}

// Detail Getter
func (r AlibabaAlihealthDrugCodeKytQuerycodeflowRequest) GetDetail() string {
    return r._detail
}
