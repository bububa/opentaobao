package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YunostvpubadmindeviceyksbotsAPIRequest 获取设备列表 API请求
// yunos.tvpubadmin.device.yks.bots
//
// 获取设备列表
type YunostvpubadmindeviceyksbotsAPIRequest struct {
	model.Params
}

// NewYunostvpubadmindeviceyksbotsRequest 初始化YunostvpubadmindeviceyksbotsAPIRequest对象
func NewYunostvpubadmindeviceyksbotsRequest() *YunostvpubadmindeviceyksbotsAPIRequest {
	return &YunostvpubadmindeviceyksbotsAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunostvpubadmindeviceyksbotsAPIRequest) GetApiMethodName() string {
	return "yunos.tvpubadmin.device.yks.bots"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunostvpubadmindeviceyksbotsAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunostvpubadmindeviceyksbotsAPIRequest) GetRawParams() model.Params {
	return r.Params
}
