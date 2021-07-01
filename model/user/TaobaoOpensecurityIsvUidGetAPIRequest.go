package user

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOpensecurityIsvUidGetAPIRequest
获取open security uid for isv API请求
taobao.opensecurity.isv.uid.get

根据 open_uid 获取 open_uid_isv 用于同一个 isv的多个app间数据关联 */
type TaobaoOpensecurityIsvUidGetAPIRequest struct {
	model.Params
	// 基于Appkey生成的open security淘宝用户Id
	_openUid string
}

// New
