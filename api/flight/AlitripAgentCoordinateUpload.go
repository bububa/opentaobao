package flight

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/flight"
)

// Alitripagentcoordinateupload 协同单附件凭证上传
// alitrip.agent.coordinate.upload
//
// 协同单附件凭证上传
func Alitripagentcoordinateupload(clt *core.SDKClient, req *flight.AlitripagentcoordinateuploadAPIRequest, session string) (*flight.AlitripagentcoordinateuploadAPIResponse, error) {
	var resp flight.AlitripagentcoordinateuploadAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
