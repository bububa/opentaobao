package aliyun

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AccountAliyuncsComCreateAliyunAccountForBid2013_07_01APIRequest
为bid用户创建账号 API请求
account.aliyuncs.com.CreateAliyunAccountForBid.2013-07-01

给指定的bid创建账号，同时账号打上owner bid的标记 */
type AccountAliyuncsComCreateAliyunAccountForBid2013_07_01APIRequest struct {
	model.Params
}

// New
