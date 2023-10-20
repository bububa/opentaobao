package campus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// Alibabacampusspacetypegetbycode 根据类别编码查询类别
// alibaba.campus.space.type.getbycode
//
// 根据类别编码查询类别
// HSF接口名称：com.alibaba.campus.space.api.top.SpaceTypeApiTopService
// HSF方法名称：getByCode
func Alibabacampusspacetypegetbycode(clt *core.SDKClient, req *campus.AlibabacampusspacetypegetbycodeAPIRequest, session string) (*campus.AlibabacampusspacetypegetbycodeAPIResponse, error) {
	var resp campus.AlibabacampusspacetypegetbycodeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
