package tanx

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tanx"
)

// Taobaotanxcreativesget 批量获取DSP用户的创意审核结果
// taobao.tanx.creatives.get
//
// 批量获取DSP用户的创意审核结果
func Taobaotanxcreativesget(clt *core.SDKClient, req *tanx.TaobaotanxcreativesgetAPIRequest, session string) (*tanx.TaobaotanxcreativesgetAPIResponse, error) {
	var resp tanx.TaobaotanxcreativesgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
