package aliqin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAliqinFcIotCardInfoAPIRequest
物联卡信息查询 API请求
alibaba.aliqin.fc.iot.cardInfo

物联卡信息查询 */
type AlibabaAliqinFcIotCardInfoAPIRequest struct {
	model.Params
	// SIM卡号
	_iccid string
}

// New
