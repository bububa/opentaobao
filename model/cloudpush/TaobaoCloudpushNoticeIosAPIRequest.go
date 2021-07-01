package cloudpush

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoCloudpushNoticeIosAPIRequest
推送通知给ios设备 API请求
taobao.cloudpush.notice.ios

推送通知给ios设备 */
type TaobaoCloudpushNoticeIosAPIRequest struct {
	model.Params
	// 通知摘要
	_summary string
	// 推送目标: device:推送给设备; account:推送给指定帐号,all: 推送给全部
	_target string
	// 根据Target来设定，如Target=device, 则对应的值为 设备id1,设备id2. 多个值使用逗号分隔
	_targetValue string
	// iOS的通知是通过APNS中心来发送的，需要填写对应的环境信息.  DEV:表示开发环境, PRODUCT: 表示生产环境.
	_env string
	// 提供给IOS通知的扩展属性，如角标或者声音等,注意：参数值为json
	_ext string
}

// New
