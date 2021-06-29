package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
码流向查询 APIRequest
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

func NewAlibabaAlihealthDrugCodeKytQuerycodeflowRequest() *AlibabaAlihealthDrugCodeKytQuerycodeflowRequest{
    return &AlibabaAlihealthDrugCodeKytQuerycodeflowRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthDrugCodeKytQuerycodeflowRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.code.kyt.querycodeflow"
}

func (r AlibabaAlihealthDrugCodeKytQuerycodeflowRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthDrugCodeKytQuerycodeflowRequest) SetRefEntId(refEntId string) error {
    r.refEntId = refEntId
    r.Set("ref_ent_id", refEntId)
    return nil
}

func (r AlibabaAlihealthDrugCodeKytQuerycodeflowRequest) GetRefEntId() string {
    return r.refEntId
}

func (r *AlibabaAlihealthDrugCodeKytQuerycodeflowRequest) SetCode(code string) error {
    r.code = code
    r.Set("code", code)
    return nil
}

func (r AlibabaAlihealthDrugCodeKytQuerycodeflowRequest) GetCode() string {
    return r.code
}

func (r *AlibabaAlihealthDrugCodeKytQuerycodeflowRequest) SetLongitude(longitude string) error {
    r.longitude = longitude
    r.Set("longitude", longitude)
    return nil
}

func (r AlibabaAlihealthDrugCodeKytQuerycodeflowRequest) GetLongitude() string {
    return r.longitude
}

func (r *AlibabaAlihealthDrugCodeKytQuerycodeflowRequest) SetLatitude(latitude string) error {
    r.latitude = latitude
    r.Set("latitude", latitude)
    return nil
}

func (r AlibabaAlihealthDrugCodeKytQuerycodeflowRequest) GetLatitude() string {
    return r.latitude
}

func (r *AlibabaAlihealthDrugCodeKytQuerycodeflowRequest) SetQueryProv(queryProv string) error {
    r.queryProv = queryProv
    r.Set("query_prov", queryProv)
    return nil
}

func (r AlibabaAlihealthDrugCodeKytQuerycodeflowRequest) GetQueryProv() string {
    return r.queryProv
}

func (r *AlibabaAlihealthDrugCodeKytQuerycodeflowRequest) SetQueryCity(queryCity string) error {
    r.queryCity = queryCity
    r.Set("query_city", queryCity)
    return nil
}

func (r AlibabaAlihealthDrugCodeKytQuerycodeflowRequest) GetQueryCity() string {
    return r.queryCity
}

func (r *AlibabaAlihealthDrugCodeKytQuerycodeflowRequest) SetQueryArea(queryArea string) error {
    r.queryArea = queryArea
    r.Set("query_area", queryArea)
    return nil
}

func (r AlibabaAlihealthDrugCodeKytQuerycodeflowRequest) GetQueryArea() string {
    return r.queryArea
}

func (r *AlibabaAlihealthDrugCodeKytQuerycodeflowRequest) SetQueryRegionCode(queryRegionCode string) error {
    r.queryRegionCode = queryRegionCode
    r.Set("query_region_code", queryRegionCode)
    return nil
}

func (r AlibabaAlihealthDrugCodeKytQuerycodeflowRequest) GetQueryRegionCode() string {
    return r.queryRegionCode
}

func (r *AlibabaAlihealthDrugCodeKytQuerycodeflowRequest) SetDetail(detail string) error {
    r.detail = detail
    r.Set("detail", detail)
    return nil
}

func (r AlibabaAlihealthDrugCodeKytQuerycodeflowRequest) GetDetail() string {
    return r.detail
}

