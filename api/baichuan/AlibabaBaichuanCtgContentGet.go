package baichuan

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/baichuan"
)

// Alibababaichuanctgcontentget 百川内容平台内容获取
// alibaba.baichuan.ctg.content.get
//
// 百川内容平台内容获取
func Alibababaichuanctgcontentget(clt *core.SDKClient, req *baichuan.AlibababaichuanctgcontentgetAPIRequest, session string) (*baichuan.AlibababaichuanctgcontentgetAPIResponse, error) {
	var resp baichuan.AlibababaichuanctgcontentgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
