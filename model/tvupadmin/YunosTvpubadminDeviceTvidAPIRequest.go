package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* YunosTvpubadminDeviceTvidAPIRequest
查询终端信息 API请求
yunos.tvpubadmin.device.tvid

通过UUID查询终端信息 */
type YunosTvpubadminDeviceTvidAPIRequest struct {
	model.Params
	// 设备的UUID
	_uuid string
}

// New
