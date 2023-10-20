package cloudpush

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cloudpush"
)

// TaobaoCloudpushNoticeAndroid 百川云推送发送通知给android
// taobao.cloudpush.notice.android
//
// 百川云推送发送通知给android
func TaobaoCloudpushNoticeAndroid(clt *core.SDKClient, req *cloudpush.TaobaoCloudpushNoticeAndroidAPIRequest, resp *cloudpush.TaobaoCloudpushNoticeAndroidAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
