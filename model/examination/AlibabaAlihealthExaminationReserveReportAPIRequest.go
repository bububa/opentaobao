package examination

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
体检机构对接_体检报告查询 API请求
alibaba.alihealth.examination.reserve.report

体检机构对接_体检报告获取
*/
type AlibabaAlihealthExaminationReserveReportAPIRequest struct {
    model.Params
    // 商户唯一码
    _merchantCode   string
    // 阿里健康预约唯一标识
    _reserveNumber   string
    // 到检唯一标识
    _checkNo   string
    // 体检机构预约唯一标识码
    _uniqReserveCode   string
    // 查询报告卡号
    _searchNo   string
    // 查询报告密码
    _searchPwd   string
}

// 初始化AlibabaAlihealthExaminationReserveReportAPIRequest对象
func NewAlibabaAlihealthExaminationReserveReportRequest() *AlibabaAlihealthExaminationReserveReportAPIRequest{
    return &AlibabaAlihealthExaminationReserveReportAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthExaminationReserveReportAPIRequest) GetApiMethodName() string {
    return "alibaba.alihealth.examination.reserve.report"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthExaminationReserveReportAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// MerchantCode Setter
// 商户唯一码
func (r *AlibabaAlihealthExaminationReserveReportAPIRequest) SetMerchantCode(_merchantCode string) error {
    r._merchantCode = _merchantCode
    r.Set("merchant_code", _merchantCode)
    return nil
}

// MerchantCode Getter
func (r AlibabaAlihealthExaminationReserveReportAPIRequest) GetMerchantCode() string {
    return r._merchantCode
}
// ReserveNumber Setter
// 阿里健康预约唯一标识
func (r *AlibabaAlihealthExaminationReserveReportAPIRequest) SetReserveNumber(_reserveNumber string) error {
    r._reserveNumber = _reserveNumber
    r.Set("reserve_number", _reserveNumber)
    return nil
}

// ReserveNumber Getter
func (r AlibabaAlihealthExaminationReserveReportAPIRequest) GetReserveNumber() string {
    return r._reserveNumber
}
// CheckNo Setter
// 到检唯一标识
func (r *AlibabaAlihealthExaminationReserveReportAPIRequest) SetCheckNo(_checkNo string) error {
    r._checkNo = _checkNo
    r.Set("check_no", _checkNo)
    return nil
}

// CheckNo Getter
func (r AlibabaAlihealthExaminationReserveReportAPIRequest) GetCheckNo() string {
    return r._checkNo
}
// UniqReserveCode Setter
// 体检机构预约唯一标识码
func (r *AlibabaAlihealthExaminationReserveReportAPIRequest) SetUniqReserveCode(_uniqReserveCode string) error {
    r._uniqReserveCode = _uniqReserveCode
    r.Set("uniq_reserve_code", _uniqReserveCode)
    return nil
}

// UniqReserveCode Getter
func (r AlibabaAlihealthExaminationReserveReportAPIRequest) GetUniqReserveCode() string {
    return r._uniqReserveCode
}
// SearchNo Setter
// 查询报告卡号
func (r *AlibabaAlihealthExaminationReserveReportAPIRequest) SetSearchNo(_searchNo string) error {
    r._searchNo = _searchNo
    r.Set("search_no", _searchNo)
    return nil
}

// SearchNo Getter
func (r AlibabaAlihealthExaminationReserveReportAPIRequest) GetSearchNo() string {
    return r._searchNo
}
// SearchPwd Setter
// 查询报告密码
func (r *AlibabaAlihealthExaminationReserveReportAPIRequest) SetSearchPwd(_searchPwd string) error {
    r._searchPwd = _searchPwd
    r.Set("search_pwd", _searchPwd)
    return nil
}

// SearchPwd Getter
func (r AlibabaAlihealthExaminationReserveReportAPIRequest) GetSearchPwd() string {
    return r._searchPwd
}
