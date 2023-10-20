package jst

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoModifyorderOrderCheckAPIRequest 自助改单服务发货订单校验 API请求
// taobao.modifyorder.order.check
//
// 自助改单服务发货后订单回传接口
type TaobaoModifyorderOrderCheckAPIRequest struct {
	model.Params
	// 商品skuId
	_finalSkuId string
	// 订单id
	_bizOrderId string
	// 商品skuId
	_finalOuterId string
	// 子订单id
	_subBizOrderId string
	// 地址oaid
	_oaid string
}

// NewTaobaoModifyorderOrderCheckRequest 初始化TaobaoModifyorderOrderCheckAPIRequest对象
func NewTaobaoModifyorderOrderCheckRequest() *TaobaoModifyorderOrderCheckAPIRequest {
	return &TaobaoModifyorderOrderCheckAPIRequest{
		Params: model.NewParams(5),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoModifyorderOrderCheckAPIRequest) Reset() {
	r._finalSkuId = ""
	r._bizOrderId = ""
	r._finalOuterId = ""
	r._subBizOrderId = ""
	r._oaid = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoModifyorderOrderCheckAPIRequest) GetApiMethodName() string {
	return "taobao.modifyorder.order.check"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoModifyorderOrderCheckAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoModifyorderOrderCheckAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFinalSkuId is FinalSkuId Setter
// 商品skuId
func (r *TaobaoModifyorderOrderCheckAPIRequest) SetFinalSkuId(_finalSkuId string) error {
	r._finalSkuId = _finalSkuId
	r.Set("final_sku_id", _finalSkuId)
	return nil
}

// GetFinalSkuId FinalSkuId Getter
func (r TaobaoModifyorderOrderCheckAPIRequest) GetFinalSkuId() string {
	return r._finalSkuId
}

// SetBizOrderId is BizOrderId Setter
// 订单id
func (r *TaobaoModifyorderOrderCheckAPIRequest) SetBizOrderId(_bizOrderId string) error {
	r._bizOrderId = _bizOrderId
	r.Set("biz_order_id", _bizOrderId)
	return nil
}

// GetBizOrderId BizOrderId Getter
func (r TaobaoModifyorderOrderCheckAPIRequest) GetBizOrderId() string {
	return r._bizOrderId
}

// SetFinalOuterId is FinalOuterId Setter
// 商品skuId
func (r *TaobaoModifyorderOrderCheckAPIRequest) SetFinalOuterId(_finalOuterId string) error {
	r._finalOuterId = _finalOuterId
	r.Set("final_outer_id", _finalOuterId)
	return nil
}

// GetFinalOuterId FinalOuterId Getter
func (r TaobaoModifyorderOrderCheckAPIRequest) GetFinalOuterId() string {
	return r._finalOuterId
}

// SetSubBizOrderId is SubBizOrderId Setter
// 子订单id
func (r *TaobaoModifyorderOrderCheckAPIRequest) SetSubBizOrderId(_subBizOrderId string) error {
	r._subBizOrderId = _subBizOrderId
	r.Set("sub_biz_order_id", _subBizOrderId)
	return nil
}

// GetSubBizOrderId SubBizOrderId Getter
func (r TaobaoModifyorderOrderCheckAPIRequest) GetSubBizOrderId() string {
	return r._subBizOrderId
}

// SetOaid is Oaid Setter
// 地址oaid
func (r *TaobaoModifyorderOrderCheckAPIRequest) SetOaid(_oaid string) error {
	r._oaid = _oaid
	r.Set("oaid", _oaid)
	return nil
}

// GetOaid Oaid Getter
func (r TaobaoModifyorderOrderCheckAPIRequest) GetOaid() string {
	return r._oaid
}

var poolTaobaoModifyorderOrderCheckAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoModifyorderOrderCheckRequest()
	},
}

// GetTaobaoModifyorderOrderCheckRequest 从 sync.Pool 获取 TaobaoModifyorderOrderCheckAPIRequest
func GetTaobaoModifyorderOrderCheckAPIRequest() *TaobaoModifyorderOrderCheckAPIRequest {
	return poolTaobaoModifyorderOrderCheckAPIRequest.Get().(*TaobaoModifyorderOrderCheckAPIRequest)
}

// ReleaseTaobaoModifyorderOrderCheckAPIRequest 将 TaobaoModifyorderOrderCheckAPIRequest 放入 sync.Pool
func ReleaseTaobaoModifyorderOrderCheckAPIRequest(v *TaobaoModifyorderOrderCheckAPIRequest) {
	v.Reset()
	poolTaobaoModifyorderOrderCheckAPIRequest.Put(v)
}
