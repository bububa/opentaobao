package campus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// Alibabacampusspaceattrsetattr 新增业务属性实例接口
// alibaba.campus.space.attr.setattr
//
// 新增业务属性实例接口
func Alibabacampusspaceattrsetattr(clt *core.SDKClient, req *campus.AlibabacampusspaceattrsetattrAPIRequest, session string) (*campus.AlibabacampusspaceattrsetattrAPIResponse, error) {
	var resp campus.AlibabacampusspaceattrsetattrAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
