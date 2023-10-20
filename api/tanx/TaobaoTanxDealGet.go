package tanx

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tanx"
)

// Taobaotanxdealget 对外部dsp提供交易id查询接口
// taobao.tanx.deal.get
//
// 对外部dsp提供交易id查询接口
func Taobaotanxdealget(clt *core.SDKClient, req *tanx.TaobaotanxdealgetAPIRequest, session string) (*tanx.TaobaotanxdealgetAPIResponse, error) {
	var resp tanx.TaobaotanxdealgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
