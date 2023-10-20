package cloudpush

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cloudpush"
)

// Taobaocloudpushnoticeios 推送通知给ios设备
// taobao.cloudpush.notice.ios
//
// 推送通知给ios设备
func Taobaocloudpushnoticeios(clt *core.SDKClient, req *cloudpush.TaobaocloudpushnoticeiosAPIRequest, session string) (*cloudpush.TaobaocloudpushnoticeiosAPIResponse, error) {
	var resp cloudpush.TaobaocloudpushnoticeiosAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
