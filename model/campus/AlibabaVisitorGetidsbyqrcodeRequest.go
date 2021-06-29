package campus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据访客二维码查访客行程id APIRequest
alibaba.visitor.getidsbyqrcode

根据支付宝阿里访客小程序的动态二维码查询来访行程id
*/
type AlibabaVisitorGetidsbyqrcodeRequest struct {
    model.Params

    // 公司id
    companyId   int64 

    // 园区id
    campusId   int64 

    // 来访时间
    date   string 

    // 二维码字符串
    qrCode   string 

}

func NewAlibabaVisitorGetidsbyqrcodeRequest() *AlibabaVisitorGetidsbyqrcodeRequest{
    return &AlibabaVisitorGetidsbyqrcodeRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaVisitorGetidsbyqrcodeRequest) GetApiMethodName() string {
    return "alibaba.visitor.getidsbyqrcode"
}

func (r AlibabaVisitorGetidsbyqrcodeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaVisitorGetidsbyqrcodeRequest) SetCompanyId(companyId int64) error {
    r.companyId = companyId
    r.Set("company_id", companyId)
    return nil
}

func (r AlibabaVisitorGetidsbyqrcodeRequest) GetCompanyId() int64 {
    return r.companyId
}

func (r *AlibabaVisitorGetidsbyqrcodeRequest) SetCampusId(campusId int64) error {
    r.campusId = campusId
    r.Set("campus_id", campusId)
    return nil
}

func (r AlibabaVisitorGetidsbyqrcodeRequest) GetCampusId() int64 {
    return r.campusId
}

func (r *AlibabaVisitorGetidsbyqrcodeRequest) SetDate(date string) error {
    r.date = date
    r.Set("date", date)
    return nil
}

func (r AlibabaVisitorGetidsbyqrcodeRequest) GetDate() string {
    return r.date
}

func (r *AlibabaVisitorGetidsbyqrcodeRequest) SetQrCode(qrCode string) error {
    r.qrCode = qrCode
    r.Set("qr_code", qrCode)
    return nil
}

func (r AlibabaVisitorGetidsbyqrcodeRequest) GetQrCode() string {
    return r.qrCode
}

