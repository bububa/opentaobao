package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaVisitorGetidsbyqrcodeAPIRequest 根据访客二维码查访客行程id API请求
// alibaba.visitor.getidsbyqrcode
//
// 根据支付宝阿里访客小程序的动态二维码查询来访行程id
type AlibabaVisitorGetidsbyqrcodeAPIRequest struct {
	model.Params
	// 公司id
	_companyId int64
	// 园区id
	_campusId int64
	// 来访时间
	_date string
	// 二维码字符串
	_qrCode string
}

// NewAlibabaVisitorGetidsbyqrcodeRequest 初始化AlibabaVisitorGetidsbyqrcodeAPIRequest对象
func NewAlibabaVisitorGetidsbyqrcodeRequest() *AlibabaVisitorGetidsbyqrcodeAPIRequest {
	return &AlibabaVisitorGetidsbyqrcodeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaVisitorGetidsbyqrcodeAPIRequest) GetApiMethodName() string {
	return "alibaba.visitor.getidsbyqrcode"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaVisitorGetidsbyqrcodeAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetCompanyId is CompanyId Setter
// 公司id
func (r *AlibabaVisitorGetidsbyqrcodeAPIRequest) SetCompanyId(_companyId int64) error {
	r._companyId = _companyId
	r.Set("company_id", _companyId)
	return nil
}

// GetCompanyId CompanyId Getter
func (r AlibabaVisitorGetidsbyqrcodeAPIRequest) GetCompanyId() int64 {
	return r._companyId
}

// SetCampusId is CampusId Setter
// 园区id
func (r *AlibabaVisitorGetidsbyqrcodeAPIRequest) SetCampusId(_campusId int64) error {
	r._campusId = _campusId
	r.Set("campus_id", _campusId)
	return nil
}

// GetCampusId CampusId Getter
func (r AlibabaVisitorGetidsbyqrcodeAPIRequest) GetCampusId() int64 {
	return r._campusId
}

// SetDate is Date Setter
// 来访时间
func (r *AlibabaVisitorGetidsbyqrcodeAPIRequest) SetDate(_date string) error {
	r._date = _date
	r.Set("date", _date)
	return nil
}

// GetDate Date Getter
func (r AlibabaVisitorGetidsbyqrcodeAPIRequest) GetDate() string {
	return r._date
}

// SetQrCode is QrCode Setter
// 二维码字符串
func (r *AlibabaVisitorGetidsbyqrcodeAPIRequest) SetQrCode(_qrCode string) error {
	r._qrCode = _qrCode
	r.Set("qr_code", _qrCode)
	return nil
}

// GetQrCode QrCode Getter
func (r AlibabaVisitorGetidsbyqrcodeAPIRequest) GetQrCode() string {
	return r._qrCode
}
