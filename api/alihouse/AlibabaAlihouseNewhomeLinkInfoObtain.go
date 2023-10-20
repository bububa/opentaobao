package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseNewhomeLinkInfoObtain 落地页获取
// alibaba.alihouse.newhome.link.info.obtain
//
// 落地页获取
func AlibabaAlihouseNewhomeLinkInfoObtain(clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeLinkInfoObtainAPIRequest, session string) (*alihouse.AlibabaAlihouseNewhomeLinkInfoObtainAPIResponse, error) {
	var resp alihouse.AlibabaAlihouseNewhomeLinkInfoObtainAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
