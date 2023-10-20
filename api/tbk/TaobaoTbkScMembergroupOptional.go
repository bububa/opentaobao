package tbk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbk"
)

// TaobaoTbkScMembergroupOptional 工具服务商member组查询、新增接口
// taobao.tbk.sc.membergroup.optional
//
// 工具服务商member组查询、新增接口
func TaobaoTbkScMembergroupOptional(clt *core.SDKClient, req *tbk.TaobaoTbkScMembergroupOptionalAPIRequest, session string) (*tbk.TaobaoTbkScMembergroupOptionalAPIResponse, error) {
	var resp tbk.TaobaoTbkScMembergroupOptionalAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
