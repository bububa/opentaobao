package jae

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAplatformWeakgetAPIRequest
活动平台弱登录接口 API请求
taobao.aplatform.weakget

无线活动平台的开放接口，提供商品信息等的读操作 */
type TaobaoAplatformWeakgetAPIRequest struct {
	model.Params
	// 客户端自带参数
	_paramRichClientInfo *RichClientInfo
	// 业务自定义参数
	_paramDto *ParamDto
}

// New
