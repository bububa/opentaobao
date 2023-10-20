package tmallcar

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallcar"
)

// Tmallaliautoserviceitemget 查询服务商门店已上架服务商品列表
// tmall.aliauto.service.item.get
//
// 根据门店自定义门店编码查询门店【已上架】服务商品列表
func Tmallaliautoserviceitemget(clt *core.SDKClient, req *tmallcar.TmallaliautoserviceitemgetAPIRequest, session string) (*tmallcar.TmallaliautoserviceitemgetAPIResponse, error) {
	var resp tmallcar.TmallaliautoserviceitemgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
