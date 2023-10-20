package campus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// Alibabacampusspaceunitgetlistwithattrbygroupid 根据空间分组id、ids查空间单元信息【带空间单元业务属性信息】
// alibaba.campus.space.unit.getlistwithattrbygroupid
//
// 根据空间分组id、ids查空间单元信息【带空间单元业务属性信息】
func Alibabacampusspaceunitgetlistwithattrbygroupid(clt *core.SDKClient, req *campus.AlibabacampusspaceunitgetlistwithattrbygroupidAPIRequest, session string) (*campus.AlibabacampusspaceunitgetlistwithattrbygroupidAPIResponse, error) {
	var resp campus.AlibabacampusspaceunitgetlistwithattrbygroupidAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
