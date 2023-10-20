package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// WdkMealPosGetfetchmealcodeAPIRequest 五道口餐饮取餐号获取接口 API请求
// wdk.meal.pos.getfetchmealcode
//
// pos机创建订单前获取餐饮取餐号
type WdkMealPosGetfetchmealcodeAPIRequest struct {
	model.Params
	// 渠道店id
	_channelShopId string
}

// NewWdkMealPosGetfetchmealcodeRequest 初始化WdkMealPosGetfetchmealcodeAPIRequest对象
func NewWdkMealPosGetfetchmealcodeRequest() *WdkMealPosGetfetchmealcodeAPIRequest {
	return &WdkMealPosGetfetchmealcodeAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *WdkMealPosGetfetchmealcodeAPIRequest) Reset() {
	r._channelShopId = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r WdkMealPosGetfetchmealcodeAPIRequest) GetApiMethodName() string {
	return "wdk.meal.pos.getfetchmealcode"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r WdkMealPosGetfetchmealcodeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r WdkMealPosGetfetchmealcodeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetChannelShopId is ChannelShopId Setter
// 渠道店id
func (r *WdkMealPosGetfetchmealcodeAPIRequest) SetChannelShopId(_channelShopId string) error {
	r._channelShopId = _channelShopId
	r.Set("channel_shop_id", _channelShopId)
	return nil
}

// GetChannelShopId ChannelShopId Getter
func (r WdkMealPosGetfetchmealcodeAPIRequest) GetChannelShopId() string {
	return r._channelShopId
}

var poolWdkMealPosGetfetchmealcodeAPIRequest = sync.Pool{
	New: func() any {
		return NewWdkMealPosGetfetchmealcodeRequest()
	},
}

// GetWdkMealPosGetfetchmealcodeRequest 从 sync.Pool 获取 WdkMealPosGetfetchmealcodeAPIRequest
func GetWdkMealPosGetfetchmealcodeAPIRequest() *WdkMealPosGetfetchmealcodeAPIRequest {
	return poolWdkMealPosGetfetchmealcodeAPIRequest.Get().(*WdkMealPosGetfetchmealcodeAPIRequest)
}

// ReleaseWdkMealPosGetfetchmealcodeAPIRequest 将 WdkMealPosGetfetchmealcodeAPIRequest 放入 sync.Pool
func ReleaseWdkMealPosGetfetchmealcodeAPIRequest(v *WdkMealPosGetfetchmealcodeAPIRequest) {
	v.Reset()
	poolWdkMealPosGetfetchmealcodeAPIRequest.Put(v)
}
