package tbk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbk"
)

// TaobaoTbkScMaterialOptional 淘宝客-服务商-物料搜索
// taobao.tbk.sc.material.optional
//
// 服务商使用。支持入参推广者对应的“推广位”、关键词和相关筛选条件，获取对应的物料信息和推广者对应的推广链接。
func TaobaoTbkScMaterialOptional(clt *core.SDKClient, req *tbk.TaobaoTbkScMaterialOptionalAPIRequest, session string) (*tbk.TaobaoTbkScMaterialOptionalAPIResponse, error) {
	var resp tbk.TaobaoTbkScMaterialOptionalAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
