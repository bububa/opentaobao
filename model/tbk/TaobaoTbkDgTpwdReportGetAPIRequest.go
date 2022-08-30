package tbk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTbkDgTpwdReportGetAPIRequest 淘宝客-推广者-淘口令回流数据查询 API请求
// taobao.tbk.dg.tpwd.report.get
//
// 淘宝客获取单个淘口令的回流PV、UV数据。
type TaobaoTbkDgTpwdReportGetAPIRequest struct {
	model.Params
	// 待查询的口令
	_taoPassword string
	// mm_xxx_xxx_xxx的第3段数字
	_adzoneId string
}

// NewTaobaoTbkDgTpwdReportGetRequest 初始化TaobaoTbkDgTpwdReportGetAPIRequest对象
func NewTaobaoTbkDgTpwdReportGetRequest() *TaobaoTbkDgTpwdReportGetAPIRequest {
	return &TaobaoTbkDgTpwdReportGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTbkDgTpwdReportGetAPIRequest) GetApiMethodName() string {
	return "taobao.tbk.dg.tpwd.report.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTbkDgTpwdReportGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetTaoPassword is TaoPassword Setter
// 待查询的口令
func (r *TaobaoTbkDgTpwdReportGetAPIRequest) SetTaoPassword(_taoPassword string) error {
	r._taoPassword = _taoPassword
	r.Set("tao_password", _taoPassword)
	return nil
}

// GetTaoPassword TaoPassword Getter
func (r TaobaoTbkDgTpwdReportGetAPIRequest) GetTaoPassword() string {
	return r._taoPassword
}

// SetAdzoneId is AdzoneId Setter
// mm_xxx_xxx_xxx的第3段数字
func (r *TaobaoTbkDgTpwdReportGetAPIRequest) SetAdzoneId(_adzoneId string) error {
	r._adzoneId = _adzoneId
	r.Set("adzone_id", _adzoneId)
	return nil
}

// GetAdzoneId AdzoneId Getter
func (r TaobaoTbkDgTpwdReportGetAPIRequest) GetAdzoneId() string {
	return r._adzoneId
}
