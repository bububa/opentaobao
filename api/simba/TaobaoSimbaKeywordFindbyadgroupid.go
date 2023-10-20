package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// Taobaosimbakeywordfindbyadgroupid 根据推广单元id获取关键词
// taobao.simba.keyword.findbyadgroupid
//
// 根据一个关键词Id列表取得一组关键词
func Taobaosimbakeywordfindbyadgroupid(clt *core.SDKClient, req *simba.TaobaosimbakeywordfindbyadgroupidAPIRequest, session string) (*simba.TaobaosimbakeywordfindbyadgroupidAPIResponse, error) {
	var resp simba.TaobaosimbakeywordfindbyadgroupidAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
