package cloudpush

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cloudpush"
)

// TaobaoCloudpushNoticeIos 推送通知给ios设备
// taobao.cloudpush.notice.ios
//
// 推送通知给ios设备
func TaobaoCloudpushNoticeIos(clt *core.SDKClient, req *cloudpush.TaobaoCloudpushNoticeIosAPIRequest, session string) (*cloudpush.TaobaoCloudpushNoticeIosAPIResponse, error) {
	var resp cloudpush.TaobaoCloudpushNoticeIosAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
