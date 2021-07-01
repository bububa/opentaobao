package sungari

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoSungariInspectionSubmitAPIRequest
抽检指令录入 API请求
taobao.sungari.inspection.submit

抽检指令录入 */
type TaobaoSungariInspectionSubmitAPIRequest struct {
	model.Params
	// 抽检入参
	_data string
	// 操作类型
	_methodName string
}

// New
