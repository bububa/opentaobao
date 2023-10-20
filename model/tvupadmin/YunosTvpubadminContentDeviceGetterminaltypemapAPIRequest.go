package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YunostvpubadmincontentdevicegetterminaltypemapAPIRequest 获取终端类型表 API请求
// yunos.tvpubadmin.content.device.getterminaltypemap
//
// 获取终端类型表
type YunostvpubadmincontentdevicegetterminaltypemapAPIRequest struct {
	model.Params
}

// NewYunostvpubadmincontentdevicegetterminaltypemapRequest 初始化YunostvpubadmincontentdevicegetterminaltypemapAPIRequest对象
func NewYunostvpubadmincontentdevicegetterminaltypemapRequest() *YunostvpubadmincontentdevicegetterminaltypemapAPIRequest {
	return &YunostvpubadmincontentdevicegetterminaltypemapAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunostvpubadmincontentdevicegetterminaltypemapAPIRequest) GetApiMethodName() string {
	return "yunos.tvpubadmin.content.device.getterminaltypemap"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunostvpubadmincontentdevicegetterminaltypemapAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunostvpubadmincontentdevicegetterminaltypemapAPIRequest) GetRawParams() model.Params {
	return r.Params
}
