package alicom

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAliqinFlowCobrandcardSysnAPIRequest
联名卡信息同步 API请求
alibaba.aliqin.flow.cobrandcard.sysn

提供给浙江移动同步联名卡信息接口。 */
type AlibabaAliqinFlowCobrandcardSysnAPIRequest struct {
	model.Params
	// 淘宝nick
	_tbUserNick string
	// 手机号码
	_phone string
	// 联名卡类型cardType:1-大喵卡,2-小喵卡
	_cardType string
	// 1-激活，0-失效
	_action string
}

// New
