package flight

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/flight"
)

// AlitripPolicyProcess 政策进度查询
// alitrip.policy.process
//
// 上传特殊类型的单程/往返政策
func AlitripPolicyProcess(clt *core.SDKClient, req *flight.AlitripPolicyProcessAPIRequest, session string) (*flight.AlitripPolicyProcessAPIResponse, error) {
	var resp flight.AlitripPolicyProcessAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
