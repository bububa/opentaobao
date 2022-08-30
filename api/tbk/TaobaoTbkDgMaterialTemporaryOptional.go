package tbk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbk"
)

// TaobaoTbkDgMaterialTemporaryOptional 淘宝客-推广者-物料搜索（临时接口）
// taobao.tbk.dg.material.temporary.optional
//
// 通用物料搜索API（临时接口）
func TaobaoTbkDgMaterialTemporaryOptional(clt *core.SDKClient, req *tbk.TaobaoTbkDgMaterialTemporaryOptionalAPIRequest, session string) (*tbk.TaobaoTbkDgMaterialTemporaryOptionalAPIResponse, error) {
	var resp tbk.TaobaoTbkDgMaterialTemporaryOptionalAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
