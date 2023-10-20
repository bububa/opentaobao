package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// Taobaosimbasearchcrowdbatchadd 推广单元增加搜索人群
// taobao.simba.searchcrowd.batch.add
//
// 推广单元新增搜索人群
func Taobaosimbasearchcrowdbatchadd(clt *core.SDKClient, req *simba.TaobaosimbasearchcrowdbatchaddAPIRequest, session string) (*simba.TaobaosimbasearchcrowdbatchaddAPIResponse, error) {
	var resp simba.TaobaosimbasearchcrowdbatchaddAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
