package campus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// Alibabacampusspacegetbyids 根据ids和类型查询空间列表
// alibaba.campus.space.getbyids
//
// 根据ids和类型查询空间列表
func Alibabacampusspacegetbyids(clt *core.SDKClient, req *campus.AlibabacampusspacegetbyidsAPIRequest, session string) (*campus.AlibabacampusspacegetbyidsAPIResponse, error) {
	var resp campus.AlibabacampusspacegetbyidsAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
