package omniorder

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOmniDealerOdersGetAPIRequest 获取单笔全渠道经销商订单的详细信息 API请求
// taobao.omni.dealer.oders.get
//
// 全渠道经销商获取单笔订单的详细信息
type TaobaoOmniDealerOdersGetAPIRequest struct {
	model.Params
	// 主订单ID
	_mainOrderId int64
}

// NewTaobaoOmniDealerOdersGetRequest 初始化TaobaoOmniDealerOdersGetAPIRequest对象
func NewTaobaoOmniDealerOdersGetRequest() *TaobaoOmniDealerOdersGetAPIRequest {
	return &TaobaoOmniDealerOdersGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoOmniDealerOdersGetAPIRequest) Reset() {
	r._mainOrderId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOmniDealerOdersGetAPIRequest) GetApiMethodName() string {
	return "taobao.omni.dealer.oders.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOmniDealerOdersGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoOmniDealerOdersGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMainOrderId is MainOrderId Setter
// 主订单ID
func (r *TaobaoOmniDealerOdersGetAPIRequest) SetMainOrderId(_mainOrderId int64) error {
	r._mainOrderId = _mainOrderId
	r.Set("main_order_id", _mainOrderId)
	return nil
}

// GetMainOrderId MainOrderId Getter
func (r TaobaoOmniDealerOdersGetAPIRequest) GetMainOrderId() int64 {
	return r._mainOrderId
}

var poolTaobaoOmniDealerOdersGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoOmniDealerOdersGetRequest()
	},
}

// GetTaobaoOmniDealerOdersGetRequest 从 sync.Pool 获取 TaobaoOmniDealerOdersGetAPIRequest
func GetTaobaoOmniDealerOdersGetAPIRequest() *TaobaoOmniDealerOdersGetAPIRequest {
	return poolTaobaoOmniDealerOdersGetAPIRequest.Get().(*TaobaoOmniDealerOdersGetAPIRequest)
}

// ReleaseTaobaoOmniDealerOdersGetAPIRequest 将 TaobaoOmniDealerOdersGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoOmniDealerOdersGetAPIRequest(v *TaobaoOmniDealerOdersGetAPIRequest) {
	v.Reset()
	poolTaobaoOmniDealerOdersGetAPIRequest.Put(v)
}
