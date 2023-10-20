package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// Taobaosimbaserchcrowdstatebatchupdate 单品搜索人群修改状态
// taobao.simba.serchcrowd.state.batch.update
//
// 暂停或启用单品推广搜索人群溢价
func Taobaosimbaserchcrowdstatebatchupdate(clt *core.SDKClient, req *simba.TaobaosimbaserchcrowdstatebatchupdateAPIRequest, session string) (*simba.TaobaosimbaserchcrowdstatebatchupdateAPIResponse, error) {
	var resp simba.TaobaosimbaserchcrowdstatebatchupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
