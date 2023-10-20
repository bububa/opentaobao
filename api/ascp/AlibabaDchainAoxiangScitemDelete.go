package ascp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// Alibabadchainaoxiangscitemdelete 货品删除
// alibaba.dchain.aoxiang.scitem.delete
//
// 货品删除
func Alibabadchainaoxiangscitemdelete(clt *core.SDKClient, req *ascp.AlibabadchainaoxiangscitemdeleteAPIRequest, session string) (*ascp.AlibabadchainaoxiangscitemdeleteAPIResponse, error) {
	var resp ascp.AlibabadchainaoxiangscitemdeleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
