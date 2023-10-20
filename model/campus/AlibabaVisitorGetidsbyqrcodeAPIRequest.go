package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabavisitorgetidsbyqrcodeAPIRequest 根据访客二维码查访客行程id API请求
// alibaba.visitor.getidsbyqrcode
//
// 根据支付宝阿里访客小程序的动态二维码查询来访行程id
type AlibabavisitorgetidsbyqrcodeAPIRequest struct {
	model.Params
	// 来访时间
	_date string
	// 二维码字符串
	_qrCode string
	// 公司id
	_companyId int64
	// 园区id
	_campusId int64
}

// NewAlibabavisitorgetidsbyqrcodeRequest 初始化AlibabavisitorgetidsbyqrcodeAPIRequest对象
func NewAlibabavisitorgetidsbyqrcodeRequest() *AlibabavisitorgetidsbyqrcodeAPIRequest {
	return &AlibabavisitorgetidsbyqrcodeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabavisitorgetidsbyqrcodeAPIRequest) GetApiMethodName() string {
	return "alibaba.visitor.getidsbyqrcode"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabavisitorgetidsbyqrcodeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabavisitorgetidsbyqrcodeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDate is Date Setter
// 来访时间
func (r *AlibabavisitorgetidsbyqrcodeAPIRequest) SetDate(_date string) error {
	r._date = _date
	r.Set("date", _date)
	return nil
}

// GetDate Date Getter
func (r AlibabavisitorgetidsbyqrcodeAPIRequest) GetDate() string {
	return r._date
}

// SetQrCode is QrCode Setter
// 二维码字符串
func (r *AlibabavisitorgetidsbyqrcodeAPIRequest) SetQrCode(_qrCode string) error {
	r._qrCode = _qrCode
	r.Set("qr_code", _qrCode)
	return nil
}

// GetQrCode QrCode Getter
func (r AlibabavisitorgetidsbyqrcodeAPIRequest) GetQrCode() string {
	return r._qrCode
}

// SetCompanyId is CompanyId Setter
// 公司id
func (r *AlibabavisitorgetidsbyqrcodeAPIRequest) SetCompanyId(_companyId int64) error {
	r._companyId = _companyId
	r.Set("company_id", _companyId)
	return nil
}

// GetCompanyId CompanyId Getter
func (r AlibabavisitorgetidsbyqrcodeAPIRequest) GetCompanyId() int64 {
	return r._companyId
}

// SetCampusId is CampusId Setter
// 园区id
func (r *AlibabavisitorgetidsbyqrcodeAPIRequest) SetCampusId(_campusId int64) error {
	r._campusId = _campusId
	r.Set("campus_id", _campusId)
	return nil
}

// GetCampusId CampusId Getter
func (r AlibabavisitorgetidsbyqrcodeAPIRequest) GetCampusId() int64 {
	return r._campusId
}
