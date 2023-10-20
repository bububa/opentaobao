package tmallcar

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallAliautoEticketStatusAPIRequest 查询电子凭证状态 API请求
// tmall.aliauto.eticket.status
//
// 查询天猫汽车二轮车行业门店电子凭证状态
type TmallAliautoEticketStatusAPIRequest struct {
	model.Params
	// 核销码
	_consumeCode string
}

// NewTmallAliautoEticketStatusRequest 初始化TmallAliautoEticketStatusAPIRequest对象
func NewTmallAliautoEticketStatusRequest() *TmallAliautoEticketStatusAPIRequest {
	return &TmallAliautoEticketStatusAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallAliautoEticketStatusAPIRequest) Reset() {
	r._consumeCode = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallAliautoEticketStatusAPIRequest) GetApiMethodName() string {
	return "tmall.aliauto.eticket.status"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallAliautoEticketStatusAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallAliautoEticketStatusAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetConsumeCode is ConsumeCode Setter
// 核销码
func (r *TmallAliautoEticketStatusAPIRequest) SetConsumeCode(_consumeCode string) error {
	r._consumeCode = _consumeCode
	r.Set("consume_code", _consumeCode)
	return nil
}

// GetConsumeCode ConsumeCode Getter
func (r TmallAliautoEticketStatusAPIRequest) GetConsumeCode() string {
	return r._consumeCode
}

var poolTmallAliautoEticketStatusAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallAliautoEticketStatusRequest()
	},
}

// GetTmallAliautoEticketStatusRequest 从 sync.Pool 获取 TmallAliautoEticketStatusAPIRequest
func GetTmallAliautoEticketStatusAPIRequest() *TmallAliautoEticketStatusAPIRequest {
	return poolTmallAliautoEticketStatusAPIRequest.Get().(*TmallAliautoEticketStatusAPIRequest)
}

// ReleaseTmallAliautoEticketStatusAPIRequest 将 TmallAliautoEticketStatusAPIRequest 放入 sync.Pool
func ReleaseTmallAliautoEticketStatusAPIRequest(v *TmallAliautoEticketStatusAPIRequest) {
	v.Reset()
	poolTmallAliautoEticketStatusAPIRequest.Put(v)
}
