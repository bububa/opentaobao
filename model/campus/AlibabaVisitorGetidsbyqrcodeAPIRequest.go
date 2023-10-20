package campus

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaVisitorGetidsbyqrcodeAPIRequest 根据访客二维码查访客行程id API请求
// alibaba.visitor.getidsbyqrcode
//
// 根据支付宝阿里访客小程序的动态二维码查询来访行程id
type AlibabaVisitorGetidsbyqrcodeAPIRequest struct {
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

// NewAlibabaVisitorGetidsbyqrcodeRequest 初始化AlibabaVisitorGetidsbyqrcodeAPIRequest对象
func NewAlibabaVisitorGetidsbyqrcodeRequest() *AlibabaVisitorGetidsbyqrcodeAPIRequest {
	return &AlibabaVisitorGetidsbyqrcodeAPIRequest{
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaVisitorGetidsbyqrcodeAPIRequest) Reset() {
	r._date = ""
	r._qrCode = ""
	r._companyId = 0
	r._campusId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaVisitorGetidsbyqrcodeAPIRequest) GetApiMethodName() string {
	return "alibaba.visitor.getidsbyqrcode"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaVisitorGetidsbyqrcodeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaVisitorGetidsbyqrcodeAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolAlibabaVisitorGetidsbyqrcodeAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaVisitorGetidsbyqrcodeRequest()
	},
}

// GetAlibabaVisitorGetidsbyqrcodeRequest 从 sync.Pool 获取 AlibabaVisitorGetidsbyqrcodeAPIRequest
func GetAlibabaVisitorGetidsbyqrcodeAPIRequest() *AlibabaVisitorGetidsbyqrcodeAPIRequest {
	return poolAlibabaVisitorGetidsbyqrcodeAPIRequest.Get().(*AlibabaVisitorGetidsbyqrcodeAPIRequest)
}

// ReleaseAlibabaVisitorGetidsbyqrcodeAPIRequest 将 AlibabaVisitorGetidsbyqrcodeAPIRequest 放入 sync.Pool
func ReleaseAlibabaVisitorGetidsbyqrcodeAPIRequest(v *AlibabaVisitorGetidsbyqrcodeAPIRequest) {
	v.Reset()
	poolAlibabaVisitorGetidsbyqrcodeAPIRequest.Put(v)
}
