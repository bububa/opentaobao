package cloudpush

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cloudpush"
)

// TaobaoCloudpushPush 百川用户使用云推送高级推送接口
// taobao.cloudpush.push
//
// 百川用户使用云推送高级推送接口
func TaobaoCloudpushPush(clt *core.SDKClient, req *cloudpush.TaobaoCloudpushPushAPIRequest, session string) (*cloudpush.TaobaoCloudpushPushAPIResponse, error) {
	var resp cloudpush.TaobaoCloudpushPushAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
