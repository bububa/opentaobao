package cloudpush

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoCloudpushNoticeAndroidAPIRequest
百川云推送发送通知给android API请求
taobao.cloudpush.notice.android

百川云推送发送通知给android */
type TaobaoCloudpushNoticeAndroidAPIRequest struct {
	model.Params
	// 通知摘要
	_summary string
	// 推送目标: device:推送给设备; account:推送给指定帐号,all: 推送给全部
	_target string
	// 根据Target来设定，如Target=device, 则对应的值为 设备id1,设备id2. 多个值使用逗号分隔
	_targetValue string
	// 通知的标题.
	_title string
}

// New
