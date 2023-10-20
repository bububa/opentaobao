package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// Taobaosimbarpttargetingtagbaseget 定向基础报表
// taobao.simba.rpt.targetingtagbase.get
//
// 获取定向基础报表
func Taobaosimbarpttargetingtagbaseget(clt *core.SDKClient, req *simba.TaobaosimbarpttargetingtagbasegetAPIRequest, session string) (*simba.TaobaosimbarpttargetingtagbasegetAPIResponse, error) {
	var resp simba.TaobaosimbarpttargetingtagbasegetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
