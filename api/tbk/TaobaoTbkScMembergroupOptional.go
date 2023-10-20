package tbk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbk"
)

// Taobaotbkscmembergroupoptional 工具服务商member组查询、新增接口
// taobao.tbk.sc.membergroup.optional
//
// 工具服务商member组查询、新增接口
func Taobaotbkscmembergroupoptional(clt *core.SDKClient, req *tbk.TaobaotbkscmembergroupoptionalAPIRequest, session string) (*tbk.TaobaotbkscmembergroupoptionalAPIResponse, error) {
	var resp tbk.TaobaotbkscmembergroupoptionalAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
