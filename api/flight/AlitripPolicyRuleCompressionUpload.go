package flight

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/flight"
)

// Alitrippolicyrulecompressionupload 大批量上传规则类型的单程/往返政策
// alitrip.policy.rule.compression.upload
//
// 大批量上传规则类型的单程/往返政策
func Alitrippolicyrulecompressionupload(clt *core.SDKClient, req *flight.AlitrippolicyrulecompressionuploadAPIRequest, session string) (*flight.AlitrippolicyrulecompressionuploadAPIResponse, error) {
	var resp flight.AlitrippolicyrulecompressionuploadAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
