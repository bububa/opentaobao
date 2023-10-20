package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkorderSharestockOrderGetAPIRequest 猫超商户订单拉取 API请求
// alibaba.wdkorder.sharestock.order.get
//
// 商户拉取猫超订单数据
type AlibabaWdkorderSharestockOrderGetAPIRequest struct {
	model.Params
	// 淘宝主订单ID
	_tbOrderId int64
}

// NewAlibabaWdkorderSharestockOrderGetRequest 初始化AlibabaWdkorderSharestockOrderGetAPIRequest对象
func NewAlibabaWdkorderSharestockOrderGetRequest() *AlibabaWdkorderSharestockOrderGetAPIRequest {
	return &AlibabaWdkorderSharestockOrderGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkorderSharestockOrderGetAPIRequest) Reset() {
	r._tbOrderId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkorderSharestockOrderGetAPIRequest) GetApiMethodName() string {
	return "alibaba.wdkorder.sharestock.order.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkorderSharestockOrderGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkorderSharestockOrderGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTbOrderId is TbOrderId Setter
// 淘宝主订单ID
func (r *AlibabaWdkorderSharestockOrderGetAPIRequest) SetTbOrderId(_tbOrderId int64) error {
	r._tbOrderId = _tbOrderId
	r.Set("tb_order_id", _tbOrderId)
	return nil
}

// GetTbOrderId TbOrderId Getter
func (r AlibabaWdkorderSharestockOrderGetAPIRequest) GetTbOrderId() int64 {
	return r._tbOrderId
}

var poolAlibabaWdkorderSharestockOrderGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkorderSharestockOrderGetRequest()
	},
}

// GetAlibabaWdkorderSharestockOrderGetRequest 从 sync.Pool 获取 AlibabaWdkorderSharestockOrderGetAPIRequest
func GetAlibabaWdkorderSharestockOrderGetAPIRequest() *AlibabaWdkorderSharestockOrderGetAPIRequest {
	return poolAlibabaWdkorderSharestockOrderGetAPIRequest.Get().(*AlibabaWdkorderSharestockOrderGetAPIRequest)
}

// ReleaseAlibabaWdkorderSharestockOrderGetAPIRequest 将 AlibabaWdkorderSharestockOrderGetAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkorderSharestockOrderGetAPIRequest(v *AlibabaWdkorderSharestockOrderGetAPIRequest) {
	v.Reset()
	poolAlibabaWdkorderSharestockOrderGetAPIRequest.Put(v)
}
