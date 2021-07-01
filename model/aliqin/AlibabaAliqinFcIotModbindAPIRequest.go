package aliqin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAliqinFcIotModbindAPIRequest
物联网绑定/换绑API API请求
alibaba.aliqin.fc.iot.modbind

支持用户的设备的换绑和解绑操作 */
type AlibabaAliqinFcIotModbindAPIRequest struct {
	model.Params
	// chgBind：换绑；unBind：解绑
	_opionType string
	// 物联卡和iccid保持一致
	_billReal string
	// iccid （20位）
	_iccid string
	// 换绑的时候必传，换的新设备imei
	_newimei string
	// 目前绑定的设备imei
	_imei string
	// 物联卡业务：若无特殊为ICCID
	_billSource string
	// 若无特殊物联卡传入122
	_midPatChannel string
}

// New
