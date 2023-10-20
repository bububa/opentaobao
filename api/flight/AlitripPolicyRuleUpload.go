package flight

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/flight"
)

// Alitrippolicyruleupload 规则政策上传
// alitrip.policy.rule.upload
//
// 上传特殊类型的单程/往返政策
func Alitrippolicyruleupload(clt *core.SDKClient, req *flight.AlitrippolicyruleuploadAPIRequest, session string) (*flight.AlitrippolicyruleuploadAPIResponse, error) {
	var resp flight.AlitrippolicyruleuploadAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
