package flight

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/flight"
)

// AlitripPolicyProcess 政策进度查询
// alitrip.policy.process
//
// 上传特殊类型的单程/往返政策
func AlitripPolicyProcess(clt *core.SDKClient, req *flight.AlitripPolicyProcessAPIRequest, resp *flight.AlitripPolicyProcessAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
