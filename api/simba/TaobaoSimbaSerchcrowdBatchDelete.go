package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// Taobaosimbaserchcrowdbatchdelete 单品搜索人群批量取消溢价
// taobao.simba.serchcrowd.batch.delete
//
// 删除单品搜索人群溢价功能
func Taobaosimbaserchcrowdbatchdelete(clt *core.SDKClient, req *simba.TaobaosimbaserchcrowdbatchdeleteAPIRequest, session string) (*simba.TaobaosimbaserchcrowdbatchdeleteAPIResponse, error) {
	var resp simba.TaobaosimbaserchcrowdbatchdeleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
