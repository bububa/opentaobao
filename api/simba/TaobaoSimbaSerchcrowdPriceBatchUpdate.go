package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// Taobaosimbaserchcrowdpricebatchupdate 单品推广搜索人群修改溢价
// taobao.simba.serchcrowd.price.batch.update
//
// 单品推广搜索人群修改溢价, 不支持跨推广单元修改
func Taobaosimbaserchcrowdpricebatchupdate(clt *core.SDKClient, req *simba.TaobaosimbaserchcrowdpricebatchupdateAPIRequest, session string) (*simba.TaobaosimbaserchcrowdpricebatchupdateAPIResponse, error) {
	var resp simba.TaobaosimbaserchcrowdpricebatchupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
