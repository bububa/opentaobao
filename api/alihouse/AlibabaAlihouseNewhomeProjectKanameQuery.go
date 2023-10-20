package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseNewhomeProjectKanameQuery 查询KA楼盘名称
// alibaba.alihouse.newhome.project.kaname.query
//
// 查询KA楼盘名称
func AlibabaAlihouseNewhomeProjectKanameQuery(clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeProjectKanameQueryAPIRequest, session string) (*alihouse.AlibabaAlihouseNewhomeProjectKanameQueryAPIResponse, error) {
	var resp alihouse.AlibabaAlihouseNewhomeProjectKanameQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
