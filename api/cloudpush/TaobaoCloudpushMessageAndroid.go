package cloudpush

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cloudpush"
)

// TaobaoCloudpushMessageAndroid 百川云推送发送消息给android
// taobao.cloudpush.message.android
//
// 百川用户使用云推送发送消息给android
func TaobaoCloudpushMessageAndroid(clt *core.SDKClient, req *cloudpush.TaobaoCloudpushMessageAndroidAPIRequest, resp *cloudpush.TaobaoCloudpushMessageAndroidAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
