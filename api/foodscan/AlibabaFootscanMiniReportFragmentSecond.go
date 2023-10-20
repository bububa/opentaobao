package foodscan

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/foodscan"
)

// Alibabafootscanminireportfragmentsecond 第二只脚生成报告接口
// alibaba.footscan.mini.report.fragment.second
//
// 第二只脚生成报告接口
func Alibabafootscanminireportfragmentsecond(clt *core.SDKClient, req *foodscan.AlibabafootscanminireportfragmentsecondAPIRequest, session string) (*foodscan.AlibabafootscanminireportfragmentsecondAPIResponse, error) {
	var resp foodscan.AlibabafootscanminireportfragmentsecondAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
