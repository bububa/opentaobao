package yunosappstore

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* YunosAppstorePadHpApplistAPIRequest
查询HpPad appList API请求
yunos.appstore.pad.hp.applist

提供hp pad应用群数据 */
type YunosAppstorePadHpApplistAPIRequest struct {
	model.Params
	// 获取的应用群code
	_code string
}

// New
