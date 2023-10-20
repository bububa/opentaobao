package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseNewhomeDynamicdataSubmit 提交小区动态信息
// alibaba.alihouse.newhome.dynamicdata.submit
//
// 提交小区动态信息
func AlibabaAlihouseNewhomeDynamicdataSubmit(clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeDynamicdataSubmitAPIRequest, session string) (*alihouse.AlibabaAlihouseNewhomeDynamicdataSubmitAPIResponse, error) {
	var resp alihouse.AlibabaAlihouseNewhomeDynamicdataSubmitAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
