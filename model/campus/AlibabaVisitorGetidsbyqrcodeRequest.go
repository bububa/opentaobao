package campus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据访客二维码查访客行程id API请求
alibaba.visitor.getidsbyqrcode

根据支付宝阿里访客小程序的动态二维码查询来访行程id
*/
type AlibabaVisitorGetidsbyqrcodeRequest struct {
    model.Params
    // 公司id
    _companyId   int64
    // 园区id
    _campusId   int64
    // 来访时间
    _date   string
    // 二维码字符串
    _qrCode   string
}

// 初始化AlibabaVisitorGetidsbyqrcodeRequest对象
func NewAlibabaVisitorGetidsbyqrcodeRequest() *AlibabaVisitorGetidsbyqrcodeRequest{
    return &AlibabaVisitorGetidsbyqrcodeRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaVisitorGetidsbyqrcodeRequest) GetApiMethodName() string {
    return "alibaba.visitor.getidsbyqrcode"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaVisitorGetidsbyqrcodeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CompanyId Setter
// 公司id
func (r *AlibabaVisitorGetidsbyqrcodeRequest) SetCompanyId(_companyId int64) error {
    r._companyId = _companyId
    r.Set("company_id", _companyId)
    return nil
}

// CompanyId Getter
func (r AlibabaVisitorGetidsbyqrcodeRequest) GetCompanyId() int64 {
    return r._companyId
}
// CampusId Setter
// 园区id
func (r *AlibabaVisitorGetidsbyqrcodeRequest) SetCampusId(_campusId int64) error {
    r._campusId = _campusId
    r.Set("campus_id", _campusId)
    return nil
}

// CampusId Getter
func (r AlibabaVisitorGetidsbyqrcodeRequest) GetCampusId() int64 {
    return r._campusId
}
// Date Setter
// 来访时间
func (r *AlibabaVisitorGetidsbyqrcodeRequest) SetDate(_date string) error {
    r._date = _date
    r.Set("date", _date)
    return nil
}

// Date Getter
func (r AlibabaVisitorGetidsbyqrcodeRequest) GetDate() string {
    return r._date
}
// QrCode Setter
// 二维码字符串
func (r *AlibabaVisitorGetidsbyqrcodeRequest) SetQrCode(_qrCode string) error {
    r._qrCode = _qrCode
    r.Set("qr_code", _qrCode)
    return nil
}

// QrCode Getter
func (r AlibabaVisitorGetidsbyqrcodeRequest) GetQrCode() string {
    return r._qrCode
}
