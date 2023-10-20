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
	_taoPassword string
	// mm_xxx_xxx_xxx的第3段数字
	_adzoneId string
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

// SetTaoPassword is TaoPassword Setter
// 待查询的口令
func (r *TaobaotbkdgtpwdreportgetAPIRequest) SetTaoPassword(_taoPassword string) error {
	r._taoPassword = _taoPassword
	r.Set("tao_password", _taoPassword)
	return nil
}

// GetTaoPassword TaoPassword Getter
func (r TaobaotbkdgtpwdreportgetAPIRequest) GetTaoPassword() string {
	return r._taoPassword
}

// SetAdzoneId is AdzoneId Setter
// mm_xxx_xxx_xxx的第3段数字
func (r *TaobaotbkdgtpwdreportgetAPIRequest) SetAdzoneId(_adzoneId string) error {
	r._adzoneId = _adzoneId
	r.Set("adzone_id", _adzoneId)
	return nil
}

// GetAdzoneId AdzoneId Getter
func (r TaobaotbkdgtpwdreportgetAPIRequest) GetAdzoneId() string {
	return r._adzoneId
}
