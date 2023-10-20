package nlife

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/nlife"
)

// AlibabaNlifeB2cTradeDownload b2c下载订单
// alibaba.nlife.b2c.trade.download
//
// 下载零售商在零售+平台创建的订单
func AlibabaNlifeB2cTradeDownload(clt *core.SDKClient, req *nlife.AlibabaNlifeB2cTradeDownloadAPIRequest, resp *nlife.AlibabaNlifeB2cTradeDownloadAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
