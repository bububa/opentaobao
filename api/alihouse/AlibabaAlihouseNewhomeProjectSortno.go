package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseNewhomeProjectSortno 新房排序值同步
// alibaba.alihouse.newhome.project.sortno
//
// 新房排序值同步
func AlibabaAlihouseNewhomeProjectSortno(clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeProjectSortnoAPIRequest, session string) (*alihouse.AlibabaAlihouseNewhomeProjectSortnoAPIResponse, error) {
	var resp alihouse.AlibabaAlihouseNewhomeProjectSortnoAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
