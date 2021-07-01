package aliyun

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AccountAliyuncsComDeleteAppForBid2013_07_01APIRequest
运营商删除用户的appkey API请求
account.aliyuncs.com.DeleteAppForBid.2013-07-01

删除用户的appkey，会校验调用的用户是否为当前运营商所创建的。 */
type AccountAliyuncsComDeleteAppForBid2013_07_01APIRequest struct {
	model.Params
	// 要删除的appkey的所有者用户的pk
	_ownerId string
	// 要删除的appkey
	_ownerAppkey string
}

// New
