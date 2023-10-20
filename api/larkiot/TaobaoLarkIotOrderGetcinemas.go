package larkiot

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/larkiot"
)

// TaobaoLarkIotOrderGetcinemas 获取iot渠道开放的影院
// taobao.lark.iot.order.getcinemas
//
// iot渠道拉取有权限访问的影院
func TaobaoLarkIotOrderGetcinemas(clt *core.SDKClient, req *larkiot.TaobaoLarkIotOrderGetcinemasAPIRequest, session string) (*larkiot.TaobaoLarkIotOrderGetcinemasAPIResponse, error) {
	var resp larkiot.TaobaoLarkIotOrderGetcinemasAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
