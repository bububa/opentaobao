package sungari

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoSungariDisposeSubmitAPIRequest
商品商家处置提交任务 API请求
taobao.sungari.dispose.submit

商品商家处置信息接口，提供政府部门发送处置信息给阿里 */
type TaobaoSungariDisposeSubmitAPIRequest struct {
	model.Params
	// 平台处置信息入参
	_info *DisposeInfoDo
}

// New
