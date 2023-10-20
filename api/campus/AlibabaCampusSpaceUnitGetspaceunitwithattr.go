package campus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// Alibabacampusspaceunitgetspaceunitwithattr 空间单元id查业务属性实例
// alibaba.campus.space.unit.getspaceunitwithattr
//
// 空间单元id查业务属性实例
func Alibabacampusspaceunitgetspaceunitwithattr(clt *core.SDKClient, req *campus.AlibabacampusspaceunitgetspaceunitwithattrAPIRequest, session string) (*campus.AlibabacampusspaceunitgetspaceunitwithattrAPIResponse, error) {
	var resp campus.AlibabacampusspaceunitgetspaceunitwithattrAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
