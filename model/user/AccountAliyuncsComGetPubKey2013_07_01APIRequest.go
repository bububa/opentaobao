package user

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AccountAliyuncsComGetPubKey2013_07_01APIRequest
获取用户公钥 API请求
account.aliyuncs.com.GetPubKey.2013-07-01

根据用户的appkey查询用户的pubkey */
type AccountAliyuncsComGetPubKey2013_07_01APIRequest struct {
	model.Params
	// appkey
	_ownerAppkey string
}

// New
