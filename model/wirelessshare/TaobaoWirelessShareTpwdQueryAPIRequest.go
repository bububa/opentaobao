package wirelessshare

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoWirelessShareTpwdQueryAPIRequest
查询解析淘口令 API请求
taobao.wireless.share.tpwd.query

查询解析淘口令 */
type TaobaoWirelessShareTpwdQueryAPIRequest struct {
	model.Params
	// 淘口令
	_passwordContent string
}

// New
