package qimen

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQimenWarehouseinfoSynchronizeAPIRequest 仓库同步接口 API请求
// taobao.qimen.warehouseinfo.synchronize
//
// 仓库同步接口
type TaobaoQimenWarehouseinfoSynchronizeAPIRequest struct {
	model.Params
	// 请求报文
	_request *TaobaoQimenWarehouseinfoSynchronizeRequest
}

// NewTaobaoQimenWarehouseinfoSynchronizeRequest 初始化TaobaoQimenWarehouseinfoSynchronizeAPIRequest对象
func NewTaobaoQimenWarehouseinfoSynchronizeRequest() *TaobaoQimenWarehouseinfoSynchronizeAPIRequest {
	return &TaobaoQimenWarehouseinfoSynchronizeAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoQimenWarehouseinfoSynchronizeAPIRequest) Reset() {
	r._request = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoQimenWarehouseinfoSynchronizeAPIRequest) GetApiMethodName() string {
	return "taobao.qimen.warehouseinfo.synchronize"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoQimenWarehouseinfoSynchronizeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoQimenWarehouseinfoSynchronizeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRequest is Request Setter
// 请求报文
func (r *TaobaoQimenWarehouseinfoSynchronizeAPIRequest) SetRequest(_request *TaobaoQimenWarehouseinfoSynchronizeRequest) error {
	r._request = _request
	r.Set("request", _request)
	return nil
}

// GetRequest Request Getter
func (r TaobaoQimenWarehouseinfoSynchronizeAPIRequest) GetRequest() *TaobaoQimenWarehouseinfoSynchronizeRequest {
	return r._request
}

var poolTaobaoQimenWarehouseinfoSynchronizeAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoQimenWarehouseinfoSynchronizeRequest()
	},
}

// GetTaobaoQimenWarehouseinfoSynchronizeRequest 从 sync.Pool 获取 TaobaoQimenWarehouseinfoSynchronizeAPIRequest
func GetTaobaoQimenWarehouseinfoSynchronizeAPIRequest() *TaobaoQimenWarehouseinfoSynchronizeAPIRequest {
	return poolTaobaoQimenWarehouseinfoSynchronizeAPIRequest.Get().(*TaobaoQimenWarehouseinfoSynchronizeAPIRequest)
}

// ReleaseTaobaoQimenWarehouseinfoSynchronizeAPIRequest 将 TaobaoQimenWarehouseinfoSynchronizeAPIRequest 放入 sync.Pool
func ReleaseTaobaoQimenWarehouseinfoSynchronizeAPIRequest(v *TaobaoQimenWarehouseinfoSynchronizeAPIRequest) {
	v.Reset()
	poolTaobaoQimenWarehouseinfoSynchronizeAPIRequest.Put(v)
}
