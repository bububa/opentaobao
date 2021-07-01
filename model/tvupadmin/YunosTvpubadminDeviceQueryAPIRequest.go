package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* YunosTvpubadminDeviceQueryAPIRequest
获取设备列表 API请求
yunos.tvpubadmin.device.query

获取设备列表 */
type YunosTvpubadminDeviceQueryAPIRequest struct {
	model.Params
	// 终端类型
	_terminalType string
	// 品牌ID
	_brandId int64
	// 牌照方
	_license int64
	// 页码值
	_pageNo int64
	// 每页条数
	_pageSize int64
}

// New
