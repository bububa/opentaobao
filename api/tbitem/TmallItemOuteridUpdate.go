package tbitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbitem"
)

// Tmallitemouteridupdate 天猫商品/SKU商家编码更新接口
// tmall.item.outerid.update
//
// 天猫商品/SKU商家编码更新接口；支持商品、SKU的商家编码同时更新；支持同一商品下的SKU批量更新。（感谢sample小雨提供接口命名）
func Tmallitemouteridupdate(clt *core.SDKClient, req *tbitem.TmallitemouteridupdateAPIRequest, session string) (*tbitem.TmallitemouteridupdateAPIResponse, error) {
	var resp tbitem.TmallitemouteridupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
