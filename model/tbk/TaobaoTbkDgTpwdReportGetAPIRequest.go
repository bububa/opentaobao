package tbk

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoTbkDgTpwdReportGetAPIRequest) Reset() {
	r._taoPassword = ""
	r._adzoneId = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTbkDgTpwdReportGetAPIRequest) GetApiMethodName() string {
	return "taobao.tbk.dg.tpwd.report.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTbkDgTpwdReportGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoTbkDgTpwdReportGetAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolTaobaoTbkDgTpwdReportGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoTbkDgTpwdReportGetRequest()
	},
}

// GetTaobaoTbkDgTpwdReportGetRequest 从 sync.Pool 获取 TaobaoTbkDgTpwdReportGetAPIRequest
func GetTaobaoTbkDgTpwdReportGetAPIRequest() *TaobaoTbkDgTpwdReportGetAPIRequest {
	return poolTaobaoTbkDgTpwdReportGetAPIRequest.Get().(*TaobaoTbkDgTpwdReportGetAPIRequest)
}

// ReleaseTaobaoTbkDgTpwdReportGetAPIRequest 将 TaobaoTbkDgTpwdReportGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoTbkDgTpwdReportGetAPIRequest(v *TaobaoTbkDgTpwdReportGetAPIRequest) {
	v.Reset()
	poolTaobaoTbkDgTpwdReportGetAPIRequest.Put(v)
}
