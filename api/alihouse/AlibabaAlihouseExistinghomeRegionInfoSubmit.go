package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// Alibabaalihouseexistinghomeregioninfosubmit 商圈专家信息同步
// alibaba.alihouse.existinghome.region.info.submit
//
// 商圈专家信息同步
func Alibabaalihouseexistinghomeregioninfosubmit(clt *core.SDKClient, req *alihouse.AlibabaalihouseexistinghomeregioninfosubmitAPIRequest, session string) (*alihouse.AlibabaalihouseexistinghomeregioninfosubmitAPIResponse, error) {
	var resp alihouse.AlibabaalihouseexistinghomeregioninfosubmitAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
