package tanx

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tanx"
)

// Taobaotanxbiddingrefusesget tanx竞价失败反馈api
// taobao.tanx.biddingrefuses.get
//
// 竞价失败反馈根据创意id查询API提供
func Taobaotanxbiddingrefusesget(clt *core.SDKClient, req *tanx.TaobaotanxbiddingrefusesgetAPIRequest, session string) (*tanx.TaobaotanxbiddingrefusesgetAPIResponse, error) {
	var resp tanx.TaobaotanxbiddingrefusesgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
