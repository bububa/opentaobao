package cloudpush

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cloudpush"
)

// TaobaoCloudpushNoticeAndroid 百川云推送发送通知给android
// taobao.cloudpush.notice.android
//
// 百川云推送发送通知给android
func TaobaoCloudpushNoticeAndroid(clt *core.SDKClient, req *cloudpush.TaobaoCloudpushNoticeAndroidAPIRequest, session string) (*cloudpush.TaobaoCloudpushNoticeAndroidAPIResponse, error) {
	var resp cloudpush.TaobaoCloudpushNoticeAndroidAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
