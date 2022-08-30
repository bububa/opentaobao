package tbk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbk"
)

// TaobaoTbkScMaterialTemporaryOptional 淘宝客-服务商-物料搜索（临时接口）
// taobao.tbk.sc.material.temporary.optional
//
// 服务商使用。支持入参推广者对应的“推广位”、关键词和相关筛选条件，获取对应的物料信息和推广者对应的推广链接。
func TaobaoTbkScMaterialTemporaryOptional(clt *core.SDKClient, req *tbk.TaobaoTbkScMaterialTemporaryOptionalAPIRequest, session string) (*tbk.TaobaoTbkScMaterialTemporaryOptionalAPIResponse, error) {
	var resp tbk.TaobaoTbkScMaterialTemporaryOptionalAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
