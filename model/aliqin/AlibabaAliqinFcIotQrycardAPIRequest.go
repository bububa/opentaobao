package aliqin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAliqinFcIotQrycardAPIRequest
查询终端信息 API请求
alibaba.aliqin.fc.iot.qrycard

查询终端信息 */
type AlibabaAliqinFcIotQrycardAPIRequest struct {
	model.Params
	// 外部计费来源
	_billSource string
	// 外部计费号
	_billReal string
	// ICCID
	_iccid string
}

// New
