package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseNewhomeCommunitySubmit 提交小区信息
// alibaba.alihouse.newhome.community.submit
//
// 提交小区信息
func AlibabaAlihouseNewhomeCommunitySubmit(clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeCommunitySubmitAPIRequest, session string) (*alihouse.AlibabaAlihouseNewhomeCommunitySubmitAPIResponse, error) {
	var resp alihouse.AlibabaAlihouseNewhomeCommunitySubmitAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
