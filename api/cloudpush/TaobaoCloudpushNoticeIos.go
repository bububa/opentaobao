package cloudpush

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cloudpush"
)

// TaobaoCloudpushNoticeIos 推送通知给ios设备
// taobao.cloudpush.notice.ios
//
// 推送通知给ios设备
func TaobaoCloudpushNoticeIos(clt *core.SDKClient, req *cloudpush.TaobaoCloudpushNoticeIosAPIRequest, resp *cloudpush.TaobaoCloudpushNoticeIosAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
