package cloudpush

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoCloudpushMessageAndroidAPIRequest
百川云推送发送消息给android API请求
taobao.cloudpush.message.android

百川用户使用云推送发送消息给android */
type TaobaoCloudpushMessageAndroidAPIRequest struct {
	model.Params
	// 发送的消息内容.
	_body string
	// 推送目标: device:推送给设备; account:推送给指定帐号,all: 推送给全部
	_target string
	// 根据Target来设定，如Target=device, 则对应的值为 设备id1,设备id2. 多个值使用逗号分隔
	_targetValue string
}

// New
