package user

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaLsyMiniappUserGetAPIRequest
零售云小程序获取登录用户信息 API请求
alibaba.lsy.miniapp.user.get

零售云小程序，通过授权码获取登录的卖家账号信息 */
type AlibabaLsyMiniappUserGetAPIRequest struct {
	model.Params
	// 当前时间戳，毫秒
	_timeStamp string
	// 获取用户信息的授权码，在小程序中获取
	_code string
	// 请求参数签名，sha1(所有入参+appSecret，按字符串升序排列)
	_signature string
	// 系统分配的小程序ID
	_appId string
}

// New
