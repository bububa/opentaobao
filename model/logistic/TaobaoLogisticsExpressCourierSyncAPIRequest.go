package logistic

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoLogisticsExpressCourierSyncAPIRequest 快递公司同步小件员信息 API请求
// taobao.logistics.express.courier.sync
//
// 快递公司同步小件员信息
type TaobaoLogisticsExpressCourierSyncAPIRequest struct {
	model.Params
	// 小件员信息
	_tmsCourierRequest *TmsCourierRequest
}

// NewTaobaoLogisticsExpressCourierSyncRequest 初始化TaobaoLogisticsExpressCourierSyncAPIRequest对象
func NewTaobaoLogisticsExpressCourierSyncRequest() *TaobaoLogisticsExpressCourierSyncAPIRequest {
	return &TaobaoLogisticsExpressCourierSyncAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoLogisticsExpressCourierSyncAPIRequest) Reset() {
	r._tmsCourierRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoLogisticsExpressCourierSyncAPIRequest) GetApiMethodName() string {
	return "taobao.logistics.express.courier.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoLogisticsExpressCourierSyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoLogisticsExpressCourierSyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTmsCourierRequest is TmsCourierRequest Setter
// 小件员信息
func (r *TaobaoLogisticsExpressCourierSyncAPIRequest) SetTmsCourierRequest(_tmsCourierRequest *TmsCourierRequest) error {
	r._tmsCourierRequest = _tmsCourierRequest
	r.Set("tms_courier_request", _tmsCourierRequest)
	return nil
}

// GetTmsCourierRequest TmsCourierRequest Getter
func (r TaobaoLogisticsExpressCourierSyncAPIRequest) GetTmsCourierRequest() *TmsCourierRequest {
	return r._tmsCourierRequest
}

var poolTaobaoLogisticsExpressCourierSyncAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoLogisticsExpressCourierSyncRequest()
	},
}

// GetTaobaoLogisticsExpressCourierSyncRequest 从 sync.Pool 获取 TaobaoLogisticsExpressCourierSyncAPIRequest
func GetTaobaoLogisticsExpressCourierSyncAPIRequest() *TaobaoLogisticsExpressCourierSyncAPIRequest {
	return poolTaobaoLogisticsExpressCourierSyncAPIRequest.Get().(*TaobaoLogisticsExpressCourierSyncAPIRequest)
}

// ReleaseTaobaoLogisticsExpressCourierSyncAPIRequest 将 TaobaoLogisticsExpressCourierSyncAPIRequest 放入 sync.Pool
func ReleaseTaobaoLogisticsExpressCourierSyncAPIRequest(v *TaobaoLogisticsExpressCourierSyncAPIRequest) {
	v.Reset()
	poolTaobaoLogisticsExpressCourierSyncAPIRequest.Put(v)
}
