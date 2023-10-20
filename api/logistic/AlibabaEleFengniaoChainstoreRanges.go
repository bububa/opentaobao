package logistic

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

// AlibabaEleFengniaoChainstoreRanges 蜂鸟查询门店配送范围接口
// alibaba.ele.fengniao.chainstore.ranges
//
// 蜂鸟查询门店配送范围接口
func AlibabaEleFengniaoChainstoreRanges(clt *core.SDKClient, req *logistic.AlibabaEleFengniaoChainstoreRangesAPIRequest, resp *logistic.AlibabaEleFengniaoChainstoreRangesAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
