package baichuan

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/baichuan"
)

// Alibababaichuantaopasswordquery 查询解析淘口令
// alibaba.baichuan.taopassword.query
//
// 查询，解析淘口令
func Alibababaichuantaopasswordquery(clt *core.SDKClient, req *baichuan.AlibababaichuantaopasswordqueryAPIRequest, session string) (*baichuan.AlibababaichuantaopasswordqueryAPIResponse, error) {
	var resp baichuan.AlibababaichuantaopasswordqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
