package nrt

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/nrt"
)

// TmallNrtBranddetailQuery 品牌详情查询
// tmall.nrt.branddetail.query
//
// 根据品牌id查询品牌的详细信息
func TmallNrtBranddetailQuery(clt *core.SDKClient, req *nrt.TmallNrtBranddetailQueryAPIRequest, resp *nrt.TmallNrtBranddetailQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
