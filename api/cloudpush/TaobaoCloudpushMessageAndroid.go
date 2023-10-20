package cloudpush

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cloudpush"
)

// Taobaocloudpushmessageandroid 百川云推送发送消息给android
// taobao.cloudpush.message.android
//
// 百川用户使用云推送发送消息给android
func Taobaocloudpushmessageandroid(clt *core.SDKClient, req *cloudpush.TaobaocloudpushmessageandroidAPIRequest, session string) (*cloudpush.TaobaocloudpushmessageandroidAPIResponse, error) {
	var resp cloudpush.TaobaocloudpushmessageandroidAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
