package nrt

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/nrt"
)

// TmallNrtBranddetailQuery 品牌详情查询
// tmall.nrt.branddetail.query
//
// 根据品牌id查询品牌的详细信息
func TmallNrtBranddetailQuery(clt *core.SDKClient, req *nrt.TmallNrtBranddetailQueryAPIRequest, session string) (*nrt.TmallNrtBranddetailQueryAPIResponse, error) {
	var resp nrt.TmallNrtBranddetailQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
