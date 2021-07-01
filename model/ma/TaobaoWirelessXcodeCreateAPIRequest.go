package ma

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoWirelessXcodeCreateAPIRequest
创建二维码/短连接 API请求
taobao.wireless.xcode.create

创建码平台的普通二维码或者长连接转短连接服务 */
type TaobaoWirelessXcodeCreateAPIRequest struct {
	model.Params
	// 码平台开放的业务code
	_bizCode string
	// 原始内容/原始地址
	_content string
	// 普通二维码样式参数
	_style *QrCodeStyle
}

// New
