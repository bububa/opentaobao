package lstlogistics

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLstLogisticsThirdpartSendAPIRequest 供应商-异云-使用第三方物流发货 API请求
// alibaba.lst.logistics.thirdpart.send
//
// 异地云仓的订单，使用第三方物流的发货方式来变更订单发货状态
type AlibabaLstLogisticsThirdpartSendAPIRequest struct {
	model.Params
	// 入参
	_param *SendOfflineOrderParam
}

// NewAlibabaLstLogisticsThirdpartSendRequest 初始化AlibabaLstLogisticsThirdpartSendAPIRequest对象
func NewAlibabaLstLogisticsThirdpartSendRequest() *AlibabaLstLogisticsThirdpartSendAPIRequest {
	return &AlibabaLstLogisticsThirdpartSendAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaLstLogisticsThirdpartSendAPIRequest) Reset() {
	r._param = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaLstLogisticsThirdpartSendAPIRequest) GetApiMethodName() string {
	return "alibaba.lst.logistics.thirdpart.send"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaLstLogisticsThirdpartSendAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaLstLogisticsThirdpartSendAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 入参
func (r *AlibabaLstLogisticsThirdpartSendAPIRequest) SetParam(_param *SendOfflineOrderParam) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabaLstLogisticsThirdpartSendAPIRequest) GetParam() *SendOfflineOrderParam {
	return r._param
}

var poolAlibabaLstLogisticsThirdpartSendAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaLstLogisticsThirdpartSendRequest()
	},
}

// GetAlibabaLstLogisticsThirdpartSendRequest 从 sync.Pool 获取 AlibabaLstLogisticsThirdpartSendAPIRequest
func GetAlibabaLstLogisticsThirdpartSendAPIRequest() *AlibabaLstLogisticsThirdpartSendAPIRequest {
	return poolAlibabaLstLogisticsThirdpartSendAPIRequest.Get().(*AlibabaLstLogisticsThirdpartSendAPIRequest)
}

// ReleaseAlibabaLstLogisticsThirdpartSendAPIRequest 将 AlibabaLstLogisticsThirdpartSendAPIRequest 放入 sync.Pool
func ReleaseAlibabaLstLogisticsThirdpartSendAPIRequest(v *AlibabaLstLogisticsThirdpartSendAPIRequest) {
	v.Reset()
	poolAlibabaLstLogisticsThirdpartSendAPIRequest.Put(v)
}
