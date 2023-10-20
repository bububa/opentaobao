package cloudpush

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cloudpush"
)

// Taobaocloudpushmessageios 百川云推送发送消息给ios
// taobao.cloudpush.message.ios
//
// 百川云推送发送消息给iOS设备.
func Taobaocloudpushmessageios(clt *core.SDKClient, req *cloudpush.TaobaocloudpushmessageiosAPIRequest, session string) (*cloudpush.TaobaocloudpushmessageiosAPIResponse, error) {
	var resp cloudpush.TaobaocloudpushmessageiosAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
