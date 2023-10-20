package flight

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/flight"
)

// Alitrippolicynormalcompressionupload 大批量上传普通类型的单程/往返政策
// alitrip.policy.normal.compression.upload
//
// 大批量上传普通类型的单程/往返政策
func Alitrippolicynormalcompressionupload(clt *core.SDKClient, req *flight.AlitrippolicynormalcompressionuploadAPIRequest, session string) (*flight.AlitrippolicynormalcompressionuploadAPIResponse, error) {
	var resp flight.AlitrippolicynormalcompressionuploadAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
