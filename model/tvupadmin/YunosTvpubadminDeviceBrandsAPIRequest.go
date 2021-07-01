package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* YunosTvpubadminDeviceBrandsAPIRequest
获取终端类型下品牌列表 API请求
yunos.tvpubadmin.device.brands

获取终端类型下品牌列表 */
type YunosTvpubadminDeviceBrandsAPIRequest struct {
	model.Params
	// 终端类型
	_terminalType string
	// 牌照方
	_license int64
}

// New
