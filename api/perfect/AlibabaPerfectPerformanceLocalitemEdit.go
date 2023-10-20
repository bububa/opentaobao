package perfect

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/perfect"
)

// Alibabaperfectperformancelocalitemedit 同城购定制发品编辑
// alibaba.perfect.performance.localitem.edit
//
// 同城购业务定制化发品接口，同城购业务线专用
func Alibabaperfectperformancelocalitemedit(clt *core.SDKClient, req *perfect.AlibabaperfectperformancelocalitemeditAPIRequest, session string) (*perfect.AlibabaperfectperformancelocalitemeditAPIResponse, error) {
	var resp perfect.AlibabaperfectperformancelocalitemeditAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
