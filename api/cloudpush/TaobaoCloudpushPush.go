package cloudpush

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cloudpush"
)

// Taobaocloudpushpush 百川用户使用云推送高级推送接口
// taobao.cloudpush.push
//
// 百川用户使用云推送高级推送接口
func Taobaocloudpushpush(clt *core.SDKClient, req *cloudpush.TaobaocloudpushpushAPIRequest, session string) (*cloudpush.TaobaocloudpushpushAPIResponse, error) {
	var resp cloudpush.TaobaocloudpushpushAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
