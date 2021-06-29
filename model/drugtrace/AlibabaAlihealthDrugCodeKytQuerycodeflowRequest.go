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
    refEntId   string
    // 追溯码
    code   string
    // 经度
    longitude   string
    // 纬度
    latitude   string
    // 查询地所在省
    queryProv   string
    // 查询地所在市
    queryCity   string
    // 查询地所在区
    queryArea   string
    // 查询地所在区域代码
    queryRegionCode   string
    // 详细地址
    detail   string
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
func (r *AlibabaAlihealthDrugCodeKytQuerycodeflowRequest) SetRefEntId(refEntId string) error {
    r.refEntId = refEntId
    r.Set("ref_ent_id", refEntId)
    return nil
}

// RefEntId Getter
func (r AlibabaAlihealthDrugCodeKytQuerycodeflowRequest) GetRefEntId() string {
    return r.refEntId
}
// Code Setter
// 追溯码
func (r *AlibabaAlihealthDrugCodeKytQuerycodeflowRequest) SetCode(code string) error {
    r.code = code
    r.Set("code", code)
    return nil
}

// Code Getter
func (r AlibabaAlihealthDrugCodeKytQuerycodeflowRequest) GetCode() string {
    return r.code
}
// Longitude Setter
// 经度
func (r *AlibabaAlihealthDrugCodeKytQuerycodeflowRequest) SetLongitude(longitude string) error {
    r.longitude = longitude
    r.Set("longitude", longitude)
    return nil
}

// Longitude Getter
func (r AlibabaAlihealthDrugCodeKytQuerycodeflowRequest) GetLongitude() string {
    return r.longitude
}
// Latitude Setter
// 纬度
func (r *AlibabaAlihealthDrugCodeKytQuerycodeflowRequest) SetLatitude(latitude string) error {
    r.latitude = latitude
    r.Set("latitude", latitude)
    return nil
}

// Latitude Getter
func (r AlibabaAlihealthDrugCodeKytQuerycodeflowRequest) GetLatitude() string {
    return r.latitude
}
// QueryProv Setter
// 查询地所在省
func (r *AlibabaAlihealthDrugCodeKytQuerycodeflowRequest) SetQueryProv(queryProv string) error {
    r.queryProv = queryProv
    r.Set("query_prov", queryProv)
    return nil
}

// QueryProv Getter
func (r AlibabaAlihealthDrugCodeKytQuerycodeflowRequest) GetQueryProv() string {
    return r.queryProv
}
// QueryCity Setter
// 查询地所在市
func (r *AlibabaAlihealthDrugCodeKytQuerycodeflowRequest) SetQueryCity(queryCity string) error {
    r.queryCity = queryCity
    r.Set("query_city", queryCity)
    return nil
}

// QueryCity Getter
func (r AlibabaAlihealthDrugCodeKytQuerycodeflowRequest) GetQueryCity() string {
    return r.queryCity
}
// QueryArea Setter
// 查询地所在区
func (r *AlibabaAlihealthDrugCodeKytQuerycodeflowRequest) SetQueryArea(queryArea string) error {
    r.queryArea = queryArea
    r.Set("query_area", queryArea)
    return nil
}

// QueryArea Getter
func (r AlibabaAlihealthDrugCodeKytQuerycodeflowRequest) GetQueryArea() string {
    return r.queryArea
}
// QueryRegionCode Setter
// 查询地所在区域代码
func (r *AlibabaAlihealthDrugCodeKytQuerycodeflowRequest) SetQueryRegionCode(queryRegionCode string) error {
    r.queryRegionCode = queryRegionCode
    r.Set("query_region_code", queryRegionCode)
    return nil
}

// QueryRegionCode Getter
func (r AlibabaAlihealthDrugCodeKytQuerycodeflowRequest) GetQueryRegionCode() string {
    return r.queryRegionCode
}
// Detail Setter
// 详细地址
func (r *AlibabaAlihealthDrugCodeKytQuerycodeflowRequest) SetDetail(detail string) error {
    r.detail = detail
    r.Set("detail", detail)
    return nil
}

// Detail Getter
func (r AlibabaAlihealthDrugCodeKytQuerycodeflowRequest) GetDetail() string {
    return r.detail
}
