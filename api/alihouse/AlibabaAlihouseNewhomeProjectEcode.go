package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseNewhomeProjectEcode 楼盘e码更新
// alibaba.alihouse.newhome.project.ecode
//
// 更新楼盘ecode
func AlibabaAlihouseNewhomeProjectEcode(clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeProjectEcodeAPIRequest, session string) (*alihouse.AlibabaAlihouseNewhomeProjectEcodeAPIResponse, error) {
	var resp alihouse.AlibabaAlihouseNewhomeProjectEcodeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
