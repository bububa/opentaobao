package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* YunosTvpubadminEpgDesktopOperationAPIRequest
影视桌面运营通用接口 API请求
yunos.tvpubadmin.epg.desktop.operation

影视桌面运营通用接口 */
type YunosTvpubadminEpgDesktopOperationAPIRequest struct {
	model.Params
	// 操作对象实体
	_entityType string
	// 操作类型
	_actionType string
	// 具体入参
	_parameter string
}

// New
