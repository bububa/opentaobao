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
    _code   string
    // 用户标识id
    _webchatId   string
    // 省编码
    _provinceCode   string
    // 市编码
    _cityCode   string
    // 区编码
    _areaCode   string
    // 扫码日期
    _scanTime   string
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
func (r *AlibabaAlihealthDrugScanQuerycodeRequest) SetCode(_code string) error {
    r._code = _code
    r.Set("code", _code)
    return nil
}

// Code Getter
func (r AlibabaAlihealthDrugScanQuerycodeRequest) GetCode() string {
    return r._code
}
// WebchatId Setter
// 用户标识id
func (r *AlibabaAlihealthDrugScanQuerycodeRequest) SetWebchatId(_webchatId string) error {
    r._webchatId = _webchatId
    r.Set("webchat_id", _webchatId)
    return nil
}

// WebchatId Getter
func (r AlibabaAlihealthDrugScanQuerycodeRequest) GetWebchatId() string {
    return r._webchatId
}
// ProvinceCode Setter
// 省编码
func (r *AlibabaAlihealthDrugScanQuerycodeRequest) SetProvinceCode(_provinceCode string) error {
    r._provinceCode = _provinceCode
    r.Set("province_code", _provinceCode)
    return nil
}

// ProvinceCode Getter
func (r AlibabaAlihealthDrugScanQuerycodeRequest) GetProvinceCode() string {
    return r._provinceCode
}
// CityCode Setter
// 市编码
func (r *AlibabaAlihealthDrugScanQuerycodeRequest) SetCityCode(_cityCode string) error {
    r._cityCode = _cityCode
    r.Set("city_code", _cityCode)
    return nil
}

// CityCode Getter
func (r AlibabaAlihealthDrugScanQuerycodeRequest) GetCityCode() string {
    return r._cityCode
}
// AreaCode Setter
// 区编码
func (r *AlibabaAlihealthDrugScanQuerycodeRequest) SetAreaCode(_areaCode string) error {
    r._areaCode = _areaCode
    r.Set("area_code", _areaCode)
    return nil
}

// AreaCode Getter
func (r AlibabaAlihealthDrugScanQuerycodeRequest) GetAreaCode() string {
    return r._areaCode
}
// ScanTime Setter
// 扫码日期
func (r *AlibabaAlihealthDrugScanQuerycodeRequest) SetScanTime(_scanTime string) error {
    r._scanTime = _scanTime
    r.Set("scan_time", _scanTime)
    return nil
}

// ScanTime Getter
func (r AlibabaAlihealthDrugScanQuerycodeRequest) GetScanTime() string {
    return r._scanTime
}
