package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询药监码对应的有效期和包装规格 APIRequest
alibaba.alihealth.drug.scan.querycode

查询药监码对应的有效期和包装规格
*/
type AlibabaAlihealthDrugScanQuerycodeRequest struct {
    model.Params

    // 溯源码
    code   string 

    // 用户标识id
    webchatId   string 

    // 省编码
    provinceCode   string 

    // 市编码
    cityCode   string 

    // 区编码
    areaCode   string 

    // 扫码日期
    scanTime   string 

}

func NewAlibabaAlihealthDrugScanQuerycodeRequest() *AlibabaAlihealthDrugScanQuerycodeRequest{
    return &AlibabaAlihealthDrugScanQuerycodeRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthDrugScanQuerycodeRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.scan.querycode"
}

func (r AlibabaAlihealthDrugScanQuerycodeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthDrugScanQuerycodeRequest) SetCode(code string) error {
    r.code = code
    r.Set("code", code)
    return nil
}

func (r AlibabaAlihealthDrugScanQuerycodeRequest) GetCode() string {
    return r.code
}

func (r *AlibabaAlihealthDrugScanQuerycodeRequest) SetWebchatId(webchatId string) error {
    r.webchatId = webchatId
    r.Set("webchat_id", webchatId)
    return nil
}

func (r AlibabaAlihealthDrugScanQuerycodeRequest) GetWebchatId() string {
    return r.webchatId
}

func (r *AlibabaAlihealthDrugScanQuerycodeRequest) SetProvinceCode(provinceCode string) error {
    r.provinceCode = provinceCode
    r.Set("province_code", provinceCode)
    return nil
}

func (r AlibabaAlihealthDrugScanQuerycodeRequest) GetProvinceCode() string {
    return r.provinceCode
}

func (r *AlibabaAlihealthDrugScanQuerycodeRequest) SetCityCode(cityCode string) error {
    r.cityCode = cityCode
    r.Set("city_code", cityCode)
    return nil
}

func (r AlibabaAlihealthDrugScanQuerycodeRequest) GetCityCode() string {
    return r.cityCode
}

func (r *AlibabaAlihealthDrugScanQuerycodeRequest) SetAreaCode(areaCode string) error {
    r.areaCode = areaCode
    r.Set("area_code", areaCode)
    return nil
}

func (r AlibabaAlihealthDrugScanQuerycodeRequest) GetAreaCode() string {
    return r.areaCode
}

func (r *AlibabaAlihealthDrugScanQuerycodeRequest) SetScanTime(scanTime string) error {
    r.scanTime = scanTime
    r.Set("scan_time", scanTime)
    return nil
}

func (r AlibabaAlihealthDrugScanQuerycodeRequest) GetScanTime() string {
    return r.scanTime
}

