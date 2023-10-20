package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// Alibabaalihouseexistinghomecompanyinfosubmit 进件接口
// alibaba.alihouse.existinghome.companyinfo.submit
//
// 进件接口
func Alibabaalihouseexistinghomecompanyinfosubmit(clt *core.SDKClient, req *alihouse.AlibabaalihouseexistinghomecompanyinfosubmitAPIRequest, session string) (*alihouse.AlibabaalihouseexistinghomecompanyinfosubmitAPIResponse, error) {
	var resp alihouse.AlibabaalihouseexistinghomecompanyinfosubmitAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
