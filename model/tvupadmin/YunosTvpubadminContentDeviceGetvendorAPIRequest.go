package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* YunosTvpubadminContentDeviceGetvendorAPIRequest
查询设备Vendor信息 API请求
yunos.tvpubadmin.content.device.getvendor

查询设备Vendor信息 */
type YunosTvpubadminContentDeviceGetvendorAPIRequest struct {
	model.Params
	// license
	_license int64
	// brand_id
	_brandId int64
}

// New
