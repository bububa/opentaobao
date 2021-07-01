package user

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOpensecurityUidGetAPIRequest
淘宝open security uid 获取接口 API请求
taobao.opensecurity.uid.get

根据明文 taobao user id 换取 app的 open_uid */
type TaobaoOpensecurityUidGetAPIRequest struct {
	model.Params
	// 淘宝用户ID
	_tbUserId int64
}

// New
