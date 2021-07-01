package qimen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoQimenWarehouseinfoSynchronizeAPIRequest
仓库同步接口 API请求
taobao.qimen.warehouseinfo.synchronize

仓库同步接口 */
type TaobaoQimenWarehouseinfoSynchronizeAPIRequest struct {
	model.Params
	// 请求报文
	_request *TaobaoQimenWarehouseinfoSynchronizeRequest
}

// NewTaobaoQimenWarehouseinfoSynchronizeRequest 初始化TaobaoQimenWarehouseinfoSynchronizeAPIRequest对象
func NewTaobaoQimenWarehouseinfoSynchronizeRequest() *TaobaoQimenWarehouseinfoSynchronizeAPIRequest {
	return &TaobaoQimenWarehouseinfoSynchronizeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoQimenWarehouseinfoSynchronizeAPIRequest) GetApiMethodName() string {
	return "taobao.qimen.warehouseinfo.synchronize"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoQimenWarehouseinfoSynchronizeAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Request Setter
// 请求报文
func (r *TaobaoQimenWarehouseinfoSynchronizeAPIRequest) SetRequest(_request *TaobaoQimenWarehouseinfoSynchronizeRequest) error {
	r._request = _request
	r.Set("request", _request)
	return nil
}

// Get Request Getter
func (r TaobaoQimenWarehouseinfoSynchronizeAPIRequest) GetRequest() *TaobaoQimenWarehouseinfoSynchronizeRequest {
	return r._request
}
