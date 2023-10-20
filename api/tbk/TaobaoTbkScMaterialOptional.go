package tbk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbk"
)

// Taobaotbkscmaterialoptional 淘宝客-服务商-物料搜索
// taobao.tbk.sc.material.optional
//
// 服务商使用。支持入参推广者对应的“推广位”、关键词和相关筛选条件，获取对应的物料信息和推广者对应的推广链接。
func Taobaotbkscmaterialoptional(clt *core.SDKClient, req *tbk.TaobaotbkscmaterialoptionalAPIRequest, session string) (*tbk.TaobaotbkscmaterialoptionalAPIResponse, error) {
	var resp tbk.TaobaotbkscmaterialoptionalAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
