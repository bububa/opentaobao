package tbk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTbkDgVegasTljReportAPIRequest 淘宝客-推广者-淘礼金效果数据 API请求
// taobao.tbk.dg.vegas.tlj.report
//
// 淘宝客推广者可查询淘礼金发放和使用等效果数据，只提供近150天的数据
type TaobaoTbkDgVegasTljReportAPIRequest struct {
	model.Params
	// 创建淘礼金时返回的rightsId
	_rightsId string
	// adzoneId
	_adzoneId int64
}

// NewTaobaoTbkDgVegasTljReportRequest 初始化TaobaoTbkDgVegasTljReportAPIRequest对象
func NewTaobaoTbkDgVegasTljReportRequest() *TaobaoTbkDgVegasTljReportAPIRequest {
	return &TaobaoTbkDgVegasTljReportAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoTbkDgVegasTljReportAPIRequest) Reset() {
	r._rightsId = ""
	r._adzoneId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTbkDgVegasTljReportAPIRequest) GetApiMethodName() string {
	return "taobao.tbk.dg.vegas.tlj.report"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTbkDgVegasTljReportAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoTbkDgVegasTljReportAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRightsId is RightsId Setter
// 创建淘礼金时返回的rightsId
func (r *TaobaoTbkDgVegasTljReportAPIRequest) SetRightsId(_rightsId string) error {
	r._rightsId = _rightsId
	r.Set("rights_id", _rightsId)
	return nil
}

// GetRightsId RightsId Getter
func (r TaobaoTbkDgVegasTljReportAPIRequest) GetRightsId() string {
	return r._rightsId
}

// SetAdzoneId is AdzoneId Setter
// adzoneId
func (r *TaobaoTbkDgVegasTljReportAPIRequest) SetAdzoneId(_adzoneId int64) error {
	r._adzoneId = _adzoneId
	r.Set("adzone_id", _adzoneId)
	return nil
}

// GetAdzoneId AdzoneId Getter
func (r TaobaoTbkDgVegasTljReportAPIRequest) GetAdzoneId() int64 {
	return r._adzoneId
}

var poolTaobaoTbkDgVegasTljReportAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoTbkDgVegasTljReportRequest()
	},
}

// GetTaobaoTbkDgVegasTljReportRequest 从 sync.Pool 获取 TaobaoTbkDgVegasTljReportAPIRequest
func GetTaobaoTbkDgVegasTljReportAPIRequest() *TaobaoTbkDgVegasTljReportAPIRequest {
	return poolTaobaoTbkDgVegasTljReportAPIRequest.Get().(*TaobaoTbkDgVegasTljReportAPIRequest)
}

// ReleaseTaobaoTbkDgVegasTljReportAPIRequest 将 TaobaoTbkDgVegasTljReportAPIRequest 放入 sync.Pool
func ReleaseTaobaoTbkDgVegasTljReportAPIRequest(v *TaobaoTbkDgVegasTljReportAPIRequest) {
	v.Reset()
	poolTaobaoTbkDgVegasTljReportAPIRequest.Put(v)
}
