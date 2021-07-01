package baichuan

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaBaichuanTaopasswordCheckAPIRequest
淘口令检查 API请求
alibaba.baichuan.taopassword.check

检查当前文本是否为淘口令 */
type AlibabaBaichuanTaopasswordCheckAPIRequest struct {
	model.Params
	// 参数DTO
	_paramDto *ParamDto
	// 系统自动生成
	_clientInfo *RichClientInfo
}

// New
