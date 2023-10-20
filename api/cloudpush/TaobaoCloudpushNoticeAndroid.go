package cloudpush

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cloudpush"
)

// Taobaocloudpushnoticeandroid 百川云推送发送通知给android
// taobao.cloudpush.notice.android
//
// 百川云推送发送通知给android
func Taobaocloudpushnoticeandroid(clt *core.SDKClient, req *cloudpush.TaobaocloudpushnoticeandroidAPIRequest, session string) (*cloudpush.TaobaocloudpushnoticeandroidAPIResponse, error) {
	var resp cloudpush.TaobaocloudpushnoticeandroidAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
