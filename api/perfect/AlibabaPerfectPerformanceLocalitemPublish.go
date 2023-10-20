package perfect

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/perfect"
)

// Alibabaperfectperformancelocalitempublish 同城购定制化发品
// alibaba.perfect.performance.localitem.publish
//
// 同城购业务定制化发品接口，同城购业务线专用
func Alibabaperfectperformancelocalitempublish(clt *core.SDKClient, req *perfect.AlibabaperfectperformancelocalitempublishAPIRequest, session string) (*perfect.AlibabaperfectperformancelocalitempublishAPIResponse, error) {
	var resp perfect.AlibabaperfectperformancelocalitempublishAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
