package aliqin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAliqinFcIotCardStatusAPIRequest
物联卡状态查询 API请求
alibaba.aliqin.fc.iot.cardStatus

物联卡状态查询 */
type AlibabaAliqinFcIotCardStatusAPIRequest struct {
	model.Params
	// SIM卡号
	_iccid string
}

// New
