package logistic

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoLogisticsExpressPickcodeCheckAPIRequest 快递公司取货码校验 API请求
// taobao.logistics.express.pickcode.check
//
// 快递公司取货码校验
type TaobaoLogisticsExpressPickcodeCheckAPIRequest struct {
	model.Params
	// 取件码校验参数
	_tmsPickCodeRequest *TmsPickCodeRequest
}

// NewTaobaoLogisticsExpressPickcodeCheckRequest 初始化TaobaoLogisticsExpressPickcodeCheckAPIRequest对象
func NewTaobaoLogisticsExpressPickcodeCheckRequest() *TaobaoLogisticsExpressPickcodeCheckAPIRequest {
	return &TaobaoLogisticsExpressPickcodeCheckAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoLogisticsExpressPickcodeCheckAPIRequest) Reset() {
	r._tmsPickCodeRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoLogisticsExpressPickcodeCheckAPIRequest) GetApiMethodName() string {
	return "taobao.logistics.express.pickcode.check"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoLogisticsExpressPickcodeCheckAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoLogisticsExpressPickcodeCheckAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTmsPickCodeRequest is TmsPickCodeRequest Setter
// 取件码校验参数
func (r *TaobaoLogisticsExpressPickcodeCheckAPIRequest) SetTmsPickCodeRequest(_tmsPickCodeRequest *TmsPickCodeRequest) error {
	r._tmsPickCodeRequest = _tmsPickCodeRequest
	r.Set("tms_pick_code_request", _tmsPickCodeRequest)
	return nil
}

// GetTmsPickCodeRequest TmsPickCodeRequest Getter
func (r TaobaoLogisticsExpressPickcodeCheckAPIRequest) GetTmsPickCodeRequest() *TmsPickCodeRequest {
	return r._tmsPickCodeRequest
}

var poolTaobaoLogisticsExpressPickcodeCheckAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoLogisticsExpressPickcodeCheckRequest()
	},
}

// GetTaobaoLogisticsExpressPickcodeCheckRequest 从 sync.Pool 获取 TaobaoLogisticsExpressPickcodeCheckAPIRequest
func GetTaobaoLogisticsExpressPickcodeCheckAPIRequest() *TaobaoLogisticsExpressPickcodeCheckAPIRequest {
	return poolTaobaoLogisticsExpressPickcodeCheckAPIRequest.Get().(*TaobaoLogisticsExpressPickcodeCheckAPIRequest)
}

// ReleaseTaobaoLogisticsExpressPickcodeCheckAPIRequest 将 TaobaoLogisticsExpressPickcodeCheckAPIRequest 放入 sync.Pool
func ReleaseTaobaoLogisticsExpressPickcodeCheckAPIRequest(v *TaobaoLogisticsExpressPickcodeCheckAPIRequest) {
	v.Reset()
	poolTaobaoLogisticsExpressPickcodeCheckAPIRequest.Put(v)
}
