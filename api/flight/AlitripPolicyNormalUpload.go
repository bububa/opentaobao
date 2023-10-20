package flight

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/flight"
)

// Alitrippolicynormalupload 普通政策上传
// alitrip.policy.normal.upload
//
// 上传普通类型的单程/往返政策
func Alitrippolicynormalupload(clt *core.SDKClient, req *flight.AlitrippolicynormaluploadAPIRequest, session string) (*flight.AlitrippolicynormaluploadAPIResponse, error) {
	var resp flight.AlitrippolicynormaluploadAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
