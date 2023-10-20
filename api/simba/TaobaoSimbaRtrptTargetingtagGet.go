package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// Taobaosimbartrpttargetingtagget 搜索人群实时报表
// taobao.simba.rtrpt.targetingtag.get
//
// 获取搜搜人群实时报表
func Taobaosimbartrpttargetingtagget(clt *core.SDKClient, req *simba.TaobaosimbartrpttargetingtaggetAPIRequest, session string) (*simba.TaobaosimbartrpttargetingtaggetAPIResponse, error) {
	var resp simba.TaobaosimbartrpttargetingtaggetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
