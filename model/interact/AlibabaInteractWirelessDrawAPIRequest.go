package interact

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabainteractwirelessdrawAPIRequest 双11到店互动无线端抽奖接口鉴权 API请求
// alibaba.interact.wireless.draw
//
// 双11到店互动无线端mtop接口开放鉴权接口，无入参出参，无安全风险，mtop接口开发 坯子
type AlibabainteractwirelessdrawAPIRequest struct {
	model.Params
}

// NewAlibabainteractwirelessdrawRequest 初始化AlibabainteractwirelessdrawAPIRequest对象
func NewAlibabainteractwirelessdrawRequest() *AlibabainteractwirelessdrawAPIRequest {
	return &AlibabainteractwirelessdrawAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabainteractwirelessdrawAPIRequest) GetApiMethodName() string {
	return "alibaba.interact.wireless.draw"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabainteractwirelessdrawAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabainteractwirelessdrawAPIRequest) GetRawParams() model.Params {
	return r.Params
}
