package campus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// Alibabacampusspacefloorgetbyid 根据id获取楼层
// alibaba.campus.space.floor.getbyid
//
// 根据id获取楼层
func Alibabacampusspacefloorgetbyid(clt *core.SDKClient, req *campus.AlibabacampusspacefloorgetbyidAPIRequest, session string) (*campus.AlibabacampusspacefloorgetbyidAPIResponse, error) {
	var resp campus.AlibabacampusspacefloorgetbyidAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
