package security

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabasecurityjaqwsgriskdatareportAPIRequest 无线保镖SDK风控数据上报 API请求
// alibaba.security.jaq.wsgriskdata.report
//
// 无线保镖sdk根据用户的需要，上报数据到聚安全云端
type AlibabasecurityjaqwsgriskdatareportAPIRequest struct {
	model.Params
	// wua串
	_wua string
	// mtopappkey是mtop的appkey
	_extParam string
}

// NewAlibabasecurityjaqwsgriskdatareportRequest 初始化AlibabasecurityjaqwsgriskdatareportAPIRequest对象
func NewAlibabasecurityjaqwsgriskdatareportRequest() *AlibabasecurityjaqwsgriskdatareportAPIRequest {
	return &AlibabasecurityjaqwsgriskdatareportAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabasecurityjaqwsgriskdatareportAPIRequest) GetApiMethodName() string {
	return "alibaba.security.jaq.wsgriskdata.report"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabasecurityjaqwsgriskdatareportAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabasecurityjaqwsgriskdatareportAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetWua is Wua Setter
// wua串
func (r *AlibabasecurityjaqwsgriskdatareportAPIRequest) SetWua(_wua string) error {
	r._wua = _wua
	r.Set("wua", _wua)
	return nil
}

// GetWua Wua Getter
func (r AlibabasecurityjaqwsgriskdatareportAPIRequest) GetWua() string {
	return r._wua
}

// SetExtParam is ExtParam Setter
// mtopappkey是mtop的appkey
func (r *AlibabasecurityjaqwsgriskdatareportAPIRequest) SetExtParam(_extParam string) error {
	r._extParam = _extParam
	r.Set("ext_param", _extParam)
	return nil
}

// GetExtParam ExtParam Getter
func (r AlibabasecurityjaqwsgriskdatareportAPIRequest) GetExtParam() string {
	return r._extParam
}
