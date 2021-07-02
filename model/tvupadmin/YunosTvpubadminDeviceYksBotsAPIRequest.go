package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YunosTvpubadminDeviceYksBotsAPIRequest 获取设备列表 API请求
// yunos.tvpubadmin.device.yks.bots
//
// 获取设备列表
type YunosTvpubadminDeviceYksBotsAPIRequest struct {
	model.Params
}

// NewYunosTvpubadminDeviceYksBotsRequest 初始化YunosTvpubadminDeviceYksBotsAPIRequest对象
func NewYunosTvpubadminDeviceYksBotsRequest() *YunosTvpubadminDeviceYksBotsAPIRequest {
	return &YunosTvpubadminDeviceYksBotsAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosTvpubadminDeviceYksBotsAPIRequest) GetApiMethodName() string {
	return "yunos.tvpubadmin.device.yks.bots"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosTvpubadminDeviceYksBotsAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}
