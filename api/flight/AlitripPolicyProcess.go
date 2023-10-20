package flight

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/flight"
)

// Alitrippolicyprocess 政策进度查询
// alitrip.policy.process
//
// 上传特殊类型的单程/往返政策
func Alitrippolicyprocess(clt *core.SDKClient, req *flight.AlitrippolicyprocessAPIRequest, session string) (*flight.AlitrippolicyprocessAPIResponse, error) {
	var resp flight.AlitrippolicyprocessAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
