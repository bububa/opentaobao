package yunosdm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* YunosDmSysGetDomainAPIRequest
获取动态域名 API请求
yunos.dm.sys.get.domain

返回alios ucp后端域名 */
type YunosDmSysGetDomainAPIRequest struct {
	model.Params
	// 制造商
	_make string
	// 设备类型
	_model string
	// 序列号
	_sn string
}

// New
