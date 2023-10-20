package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// Alibabaalihousenewhomercsync 阿里房产图文草稿信息同步
// alibaba.alihouse.newhome.rc.sync
//
// 接收图文草稿信息
func Alibabaalihousenewhomercsync(clt *core.SDKClient, req *alihouse.AlibabaalihousenewhomercsyncAPIRequest, session string) (*alihouse.AlibabaalihousenewhomercsyncAPIResponse, error) {
	var resp alihouse.AlibabaalihousenewhomercsyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
