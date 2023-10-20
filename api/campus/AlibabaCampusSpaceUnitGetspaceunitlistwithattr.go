package campus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// Alibabacampusspaceunitgetspaceunitlistwithattr 空间单元列表带业务属性实例
// alibaba.campus.space.unit.getspaceunitlistwithattr
//
// 空间单元列表带业务属性实例
func Alibabacampusspaceunitgetspaceunitlistwithattr(clt *core.SDKClient, req *campus.AlibabacampusspaceunitgetspaceunitlistwithattrAPIRequest, session string) (*campus.AlibabacampusspaceunitgetspaceunitlistwithattrAPIResponse, error) {
	var resp campus.AlibabacampusspaceunitgetspaceunitlistwithattrAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
