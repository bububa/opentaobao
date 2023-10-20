package tbk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaotbkdgvegastljstopAPIRequest 淘宝客-推广者-淘礼金暂停发放 API请求
// taobao.tbk.dg.vegas.tlj.stop
//
// 淘宝客推广者可对已经创建的淘礼金暂停发放
type TaobaotbkdgvegastljstopAPIRequest struct {
	model.Params
	// 创建淘礼金时返回的rightsId
	_rightsid string
	// adzoneId
	_adzoneid int64
}

// NewTaobaotbkdgvegastljstopRequest 初始化TaobaotbkdgvegastljstopAPIRequest对象
func NewTaobaotbkdgvegastljstopRequest() *TaobaotbkdgvegastljstopAPIRequest {
	return &TaobaotbkdgvegastljstopAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaotbkdgvegastljstopAPIRequest) GetApiMethodName() string {
	return "taobao.tbk.dg.vegas.tlj.stop"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaotbkdgvegastljstopAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaotbkdgvegastljstopAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRightsid is Rightsid Setter
// 创建淘礼金时返回的rightsId
func (r *TaobaotbkdgvegastljstopAPIRequest) SetRightsid(_rightsid string) error {
	r._rightsid = _rightsid
	r.Set("rights_id", _rightsid)
	return nil
}

// GetRightsid Rightsid Getter
func (r TaobaotbkdgvegastljstopAPIRequest) GetRightsid() string {
	return r._rightsid
}

// SetAdzoneid is Adzoneid Setter
// adzoneId
func (r *TaobaotbkdgvegastljstopAPIRequest) SetAdzoneid(_adzoneid int64) error {
	r._adzoneid = _adzoneid
	r.Set("adzone_id", _adzoneid)
	return nil
}

// GetAdzoneid Adzoneid Getter
func (r TaobaotbkdgvegastljstopAPIRequest) GetAdzoneid() int64 {
	return r._adzoneid
}
