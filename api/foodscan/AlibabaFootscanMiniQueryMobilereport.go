package foodscan

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/foodscan"
)

// Alibabafootscanminiquerymobilereport 根据scanId查询报告
// alibaba.footscan.mini.query.mobilereport
//
// 根据scanId查询报告
func Alibabafootscanminiquerymobilereport(clt *core.SDKClient, req *foodscan.AlibabafootscanminiquerymobilereportAPIRequest, session string) (*foodscan.AlibabafootscanminiquerymobilereportAPIResponse, error) {
	var resp foodscan.AlibabafootscanminiquerymobilereportAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
