package tbk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaotbkdgtpwdreportgetAPIRequest 淘宝客-推广者-淘口令回流数据查询 API请求
// taobao.tbk.dg.tpwd.report.get
//
// 淘宝客获取单个淘口令的回流PV、UV数据。
type TaobaotbkdgtpwdreportgetAPIRequest struct {
	model.Params
	// 待查询的口令
	_taopassword string
	// mm_xxx_xxx_xxx的第3段数字
	_adzoneid string
}

// NewTaobaotbkdgtpwdreportgetRequest 初始化TaobaotbkdgtpwdreportgetAPIRequest对象
func NewTaobaotbkdgtpwdreportgetRequest() *TaobaotbkdgtpwdreportgetAPIRequest {
	return &TaobaotbkdgtpwdreportgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaotbkdgtpwdreportgetAPIRequest) GetApiMethodName() string {
	return "taobao.tbk.dg.tpwd.report.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaotbkdgtpwdreportgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaotbkdgtpwdreportgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTaopassword is Taopassword Setter
// 待查询的口令
func (r *TaobaotbkdgtpwdreportgetAPIRequest) SetTaopassword(_taopassword string) error {
	r._taopassword = _taopassword
	r.Set("tao_password", _taopassword)
	return nil
}

// GetTaopassword Taopassword Getter
func (r TaobaotbkdgtpwdreportgetAPIRequest) GetTaopassword() string {
	return r._taopassword
}

// SetAdzoneid is Adzoneid Setter
// mm_xxx_xxx_xxx的第3段数字
func (r *TaobaotbkdgtpwdreportgetAPIRequest) SetAdzoneid(_adzoneid string) error {
	r._adzoneid = _adzoneid
	r.Set("adzone_id", _adzoneid)
	return nil
}

// GetAdzoneid Adzoneid Getter
func (r TaobaotbkdgtpwdreportgetAPIRequest) GetAdzoneid() string {
	return r._adzoneid
}
