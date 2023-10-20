package campus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// Alibabacampusspacegroupgetspacegroupwithattr 空间分组id查业务属性实例
// alibaba.campus.space.group.getspacegroupwithattr
//
// 空间分组id查业务属性实例
func Alibabacampusspacegroupgetspacegroupwithattr(clt *core.SDKClient, req *campus.AlibabacampusspacegroupgetspacegroupwithattrAPIRequest, session string) (*campus.AlibabacampusspacegroupgetspacegroupwithattrAPIResponse, error) {
	var resp campus.AlibabacampusspacegroupgetspacegroupwithattrAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
