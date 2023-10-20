package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// Taobaosimbaserchcrowdget 根据推广单元id获取搜索溢价人群
// taobao.simba.serchcrowd.get
//
// 根据推广单元id获取搜索溢价人群
func Taobaosimbaserchcrowdget(clt *core.SDKClient, req *simba.TaobaosimbaserchcrowdgetAPIRequest, session string) (*simba.TaobaosimbaserchcrowdgetAPIResponse, error) {
	var resp simba.TaobaosimbaserchcrowdgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
