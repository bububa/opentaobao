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
type AlibabaAlihealthExaminationReserveReportRequest struct {
    model.Params
    // 商户唯一码
    merchantCode   string
    // 阿里健康预约唯一标识
    reserveNumber   string
    // 到检唯一标识
    checkNo   string
    // 体检机构预约唯一标识码
    uniqReserveCode   string
    // 查询报告卡号
    searchNo   string
    // 查询报告密码
    searchPwd   string
}

// 初始化AlibabaAlihealthExaminationReserveReportRequest对象
func NewAlibabaAlihealthExaminationReserveReportRequest() *AlibabaAlihealthExaminationReserveReportRequest{
    return &AlibabaAlihealthExaminationReserveReportRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthExaminationReserveReportRequest) GetApiMethodName() string {
    return "alibaba.alihealth.examination.reserve.report"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthExaminationReserveReportRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// MerchantCode Setter
// 商户唯一码
func (r *AlibabaAlihealthExaminationReserveReportRequest) SetMerchantCode(merchantCode string) error {
    r.merchantCode = merchantCode
    r.Set("merchant_code", merchantCode)
    return nil
}

// MerchantCode Getter
func (r AlibabaAlihealthExaminationReserveReportRequest) GetMerchantCode() string {
    return r.merchantCode
}
// ReserveNumber Setter
// 阿里健康预约唯一标识
func (r *AlibabaAlihealthExaminationReserveReportRequest) SetReserveNumber(reserveNumber string) error {
    r.reserveNumber = reserveNumber
    r.Set("reserve_number", reserveNumber)
    return nil
}

// ReserveNumber Getter
func (r AlibabaAlihealthExaminationReserveReportRequest) GetReserveNumber() string {
    return r.reserveNumber
}
// CheckNo Setter
// 到检唯一标识
func (r *AlibabaAlihealthExaminationReserveReportRequest) SetCheckNo(checkNo string) error {
    r.checkNo = checkNo
    r.Set("check_no", checkNo)
    return nil
}

// CheckNo Getter
func (r AlibabaAlihealthExaminationReserveReportRequest) GetCheckNo() string {
    return r.checkNo
}
// UniqReserveCode Setter
// 体检机构预约唯一标识码
func (r *AlibabaAlihealthExaminationReserveReportRequest) SetUniqReserveCode(uniqReserveCode string) error {
    r.uniqReserveCode = uniqReserveCode
    r.Set("uniq_reserve_code", uniqReserveCode)
    return nil
}

// UniqReserveCode Getter
func (r AlibabaAlihealthExaminationReserveReportRequest) GetUniqReserveCode() string {
    return r.uniqReserveCode
}
// SearchNo Setter
// 查询报告卡号
func (r *AlibabaAlihealthExaminationReserveReportRequest) SetSearchNo(searchNo string) error {
    r.searchNo = searchNo
    r.Set("search_no", searchNo)
    return nil
}

// SearchNo Getter
func (r AlibabaAlihealthExaminationReserveReportRequest) GetSearchNo() string {
    return r.searchNo
}
// SearchPwd Setter
// 查询报告密码
func (r *AlibabaAlihealthExaminationReserveReportRequest) SetSearchPwd(searchPwd string) error {
    r.searchPwd = searchPwd
    r.Set("search_pwd", searchPwd)
    return nil
}

// SearchPwd Getter
func (r AlibabaAlihealthExaminationReserveReportRequest) GetSearchPwd() string {
    return r.searchPwd
}
