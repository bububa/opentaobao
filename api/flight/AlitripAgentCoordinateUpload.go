package flight

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/flight"
)

// AlitripAgentCoordinateUpload 协同单附件凭证上传
// alitrip.agent.coordinate.upload
//
// 协同单附件凭证上传
func AlitripAgentCoordinateUpload(clt *core.SDKClient, req *flight.AlitripAgentCoordinateUploadAPIRequest, session string) (*flight.AlitripAgentCoordinateUploadAPIResponse, error) {
	var resp flight.AlitripAgentCoordinateUploadAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
