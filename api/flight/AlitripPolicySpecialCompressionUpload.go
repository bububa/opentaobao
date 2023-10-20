package flight

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/flight"
)

// Alitrippolicyspecialcompressionupload 大批量上传特殊类型的单程/往返政策
// alitrip.policy.special.compression.upload
//
// 大批量上传特殊类型的单程/往返政策
func Alitrippolicyspecialcompressionupload(clt *core.SDKClient, req *flight.AlitrippolicyspecialcompressionuploadAPIRequest, session string) (*flight.AlitrippolicyspecialcompressionuploadAPIResponse, error) {
	var resp flight.AlitrippolicyspecialcompressionuploadAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
