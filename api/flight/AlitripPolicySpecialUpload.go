package flight

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/flight"
)

// Alitrippolicyspecialupload 特殊政策上传
// alitrip.policy.special.upload
//
// 上传特殊类型的单程/往返政策
func Alitrippolicyspecialupload(clt *core.SDKClient, req *flight.AlitrippolicyspecialuploadAPIRequest, session string) (*flight.AlitrippolicyspecialuploadAPIResponse, error) {
	var resp flight.AlitrippolicyspecialuploadAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
