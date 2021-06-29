package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询药监码对应的有效期和包装规格 API请求
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

// 初始化AlibabaAlihealthDrugScanQuerycodeRequest对象
func NewAlibabaAlihealthDrugScanQuerycodeRequest() *AlibabaAlihealthDrugScanQuerycodeRequest{
    return &AlibabaAlihealthDrugScanQuerycodeRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugScanQuerycodeRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.scan.querycode"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugScanQuerycodeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Code Setter
// 溯源码
func (r *AlibabaAlihealthDrugScanQuerycodeRequest) SetCode(code string) error {
    r.code = code
    r.Set("code", code)
    return nil
}

// Code Getter
func (r AlibabaAlihealthDrugScanQuerycodeRequest) GetCode() string {
    return r.code
}
// WebchatId Setter
// 用户标识id
func (r *AlibabaAlihealthDrugScanQuerycodeRequest) SetWebchatId(webchatId string) error {
    r.webchatId = webchatId
    r.Set("webchat_id", webchatId)
    return nil
}

// WebchatId Getter
func (r AlibabaAlihealthDrugScanQuerycodeRequest) GetWebchatId() string {
    return r.webchatId
}
// ProvinceCode Setter
// 省编码
func (r *AlibabaAlihealthDrugScanQuerycodeRequest) SetProvinceCode(provinceCode string) error {
    r.provinceCode = provinceCode
    r.Set("province_code", provinceCode)
    return nil
}

// ProvinceCode Getter
func (r AlibabaAlihealthDrugScanQuerycodeRequest) GetProvinceCode() string {
    return r.provinceCode
}
// CityCode Setter
// 市编码
func (r *AlibabaAlihealthDrugScanQuerycodeRequest) SetCityCode(cityCode string) error {
    r.cityCode = cityCode
    r.Set("city_code", cityCode)
    return nil
}

// CityCode Getter
func (r AlibabaAlihealthDrugScanQuerycodeRequest) GetCityCode() string {
    return r.cityCode
}
// AreaCode Setter
// 区编码
func (r *AlibabaAlihealthDrugScanQuerycodeRequest) SetAreaCode(areaCode string) error {
    r.areaCode = areaCode
    r.Set("area_code", areaCode)
    return nil
}

// AreaCode Getter
func (r AlibabaAlihealthDrugScanQuerycodeRequest) GetAreaCode() string {
    return r.areaCode
}
// ScanTime Setter
// 扫码日期
func (r *AlibabaAlihealthDrugScanQuerycodeRequest) SetScanTime(scanTime string) error {
    r.scanTime = scanTime
    r.Set("scan_time", scanTime)
    return nil
}

// ScanTime Getter
func (r AlibabaAlihealthDrugScanQuerycodeRequest) GetScanTime() string {
    return r.scanTime
}
