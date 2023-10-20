package waybill

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// CainiaoWaybillPrivacySellerOrderGetAPIRequest 隐私面单商家订单查询 API请求
// cainiao.waybill.privacy.seller.order.get
//
// 商家查询最近100天隐私面单记录
type CainiaoWaybillPrivacySellerOrderGetAPIRequest struct {
	model.Params
}

// NewCainiaoWaybillPrivacySellerOrderGetRequest 初始化CainiaoWaybillPrivacySellerOrderGetAPIRequest对象
func NewCainiaoWaybillPrivacySellerOrderGetRequest() *CainiaoWaybillPrivacySellerOrderGetAPIRequest {
	return &CainiaoWaybillPrivacySellerOrderGetAPIRequest{
		Params: model.NewParams(0),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *CainiaoWaybillPrivacySellerOrderGetAPIRequest) Reset() {
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaoWaybillPrivacySellerOrderGetAPIRequest) GetApiMethodName() string {
	return "cainiao.waybill.privacy.seller.order.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaoWaybillPrivacySellerOrderGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r CainiaoWaybillPrivacySellerOrderGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

var poolCainiaoWaybillPrivacySellerOrderGetAPIRequest = sync.Pool{
	New: func() any {
		return NewCainiaoWaybillPrivacySellerOrderGetRequest()
	},
}

// GetCainiaoWaybillPrivacySellerOrderGetRequest 从 sync.Pool 获取 CainiaoWaybillPrivacySellerOrderGetAPIRequest
func GetCainiaoWaybillPrivacySellerOrderGetAPIRequest() *CainiaoWaybillPrivacySellerOrderGetAPIRequest {
	return poolCainiaoWaybillPrivacySellerOrderGetAPIRequest.Get().(*CainiaoWaybillPrivacySellerOrderGetAPIRequest)
}

// ReleaseCainiaoWaybillPrivacySellerOrderGetAPIRequest 将 CainiaoWaybillPrivacySellerOrderGetAPIRequest 放入 sync.Pool
func ReleaseCainiaoWaybillPrivacySellerOrderGetAPIRequest(v *CainiaoWaybillPrivacySellerOrderGetAPIRequest) {
	v.Reset()
	poolCainiaoWaybillPrivacySellerOrderGetAPIRequest.Put(v)
}
