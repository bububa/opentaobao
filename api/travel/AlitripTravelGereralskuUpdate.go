package travel

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/travel"
)

// Alitriptravelgereralskuupdate 发布SKU信息（如果properties重复 则更新）
// alitrip.travel.gereralsku.update
//
// 发布SKU信息（如果properties重复 则更新）
func Alitriptravelgereralskuupdate(clt *core.SDKClient, req *travel.AlitriptravelgereralskuupdateAPIRequest, session string) (*travel.AlitriptravelgereralskuupdateAPIResponse, error) {
	var resp travel.AlitriptravelgereralskuupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
