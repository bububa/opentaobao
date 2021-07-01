package aliqin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAliqinFcIotCardofferAPIRequest
查询物联网卡上订购的offer API请求
alibaba.aliqin.fc.iot.cardoffer

查询物联网卡上订购的offer */
type AlibabaAliqinFcIotCardofferAPIRequest struct {
	model.Params
	// 具体ICCID的值
	_billreal string
	// ICCID
	_billsource string
}

// New
