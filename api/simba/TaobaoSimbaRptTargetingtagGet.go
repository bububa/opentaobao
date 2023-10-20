package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// Taobaosimbarpttargetingtagget 搜索人群离线报表
// taobao.simba.rpt.targetingtag.get
//
// 获取搜搜人群实时报表
func Taobaosimbarpttargetingtagget(clt *core.SDKClient, req *simba.TaobaosimbarpttargetingtaggetAPIRequest, session string) (*simba.TaobaosimbarpttargetingtaggetAPIResponse, error) {
	var resp simba.TaobaosimbarpttargetingtaggetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
