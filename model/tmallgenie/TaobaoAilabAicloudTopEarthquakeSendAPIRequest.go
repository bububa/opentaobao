package tmallgenie

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAilabAicloudTopEarthquakeSendAPIRequest
地震局发送地震消息 API请求
taobao.ailab.aicloud.top.earthquake.send

地震局发送地震消息给天猫精灵，天猫精灵根据地震消息判断发送地震消息给危险区域用户 */
type TaobaoAilabAicloudTopEarthquakeSendAPIRequest struct {
	model.Params
	// 扩展占位字段
	_ext string
	// 签名
	_signature string
	// 随机值
	_nonceStr string
	// 时间戳
	_timestampStr string
	// 地震信息
	_earthquakeInfo string
}

// New
