package baodian

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoBaodianDepositGetWithSdkversionAPIRequest
查询用户宝点信息（带sdk版本，已迁移） API请求
taobao.baodian.deposit.get.with.sdkversion

获取用户宝点信息（带sdk版本，已迁移） */
type TaobaoBaodianDepositGetWithSdkversionAPIRequest struct {
	model.Params
	// 设备型号
	_deviceModel string
	// uuid
	_uuid string
	// sdk版本
	_sdkVersion string
}

// New
