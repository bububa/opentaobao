package tmallservice

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallServicecenterSpserviceorderEpocTransferAPIRequest 电子保单受保人转移 API请求
// tmall.servicecenter.spserviceorder.epoc.transfer
//
// 电子保单受保人转移
type TmallServicecenterSpserviceorderEpocTransferAPIRequest struct {
	model.Params
	// 服务交易订单号
	_bizOrderId int64
}

// NewTmallServicecenterSpserviceorderEpocTransferRequest 初始化TmallServicecenterSpserviceorderEpocTransferAPIRequest对象
func NewTmallServicecenterSpserviceorderEpocTransferRequest() *TmallServicecenterSpserviceorderEpocTransferAPIRequest {
	return &TmallServicecenterSpserviceorderEpocTransferAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallServicecenterSpserviceorderEpocTransferAPIRequest) Reset() {
	r._bizOrderId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallServicecenterSpserviceorderEpocTransferAPIRequest) GetApiMethodName() string {
	return "tmall.servicecenter.spserviceorder.epoc.transfer"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallServicecenterSpserviceorderEpocTransferAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallServicecenterSpserviceorderEpocTransferAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBizOrderId is BizOrderId Setter
// 服务交易订单号
func (r *TmallServicecenterSpserviceorderEpocTransferAPIRequest) SetBizOrderId(_bizOrderId int64) error {
	r._bizOrderId = _bizOrderId
	r.Set("biz_order_id", _bizOrderId)
	return nil
}

// GetBizOrderId BizOrderId Getter
func (r TmallServicecenterSpserviceorderEpocTransferAPIRequest) GetBizOrderId() int64 {
	return r._bizOrderId
}

var poolTmallServicecenterSpserviceorderEpocTransferAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallServicecenterSpserviceorderEpocTransferRequest()
	},
}

// GetTmallServicecenterSpserviceorderEpocTransferRequest 从 sync.Pool 获取 TmallServicecenterSpserviceorderEpocTransferAPIRequest
func GetTmallServicecenterSpserviceorderEpocTransferAPIRequest() *TmallServicecenterSpserviceorderEpocTransferAPIRequest {
	return poolTmallServicecenterSpserviceorderEpocTransferAPIRequest.Get().(*TmallServicecenterSpserviceorderEpocTransferAPIRequest)
}

// ReleaseTmallServicecenterSpserviceorderEpocTransferAPIRequest 将 TmallServicecenterSpserviceorderEpocTransferAPIRequest 放入 sync.Pool
func ReleaseTmallServicecenterSpserviceorderEpocTransferAPIRequest(v *TmallServicecenterSpserviceorderEpocTransferAPIRequest) {
	v.Reset()
	poolTmallServicecenterSpserviceorderEpocTransferAPIRequest.Put(v)
}
