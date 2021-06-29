package nlife

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/nlife"
)

/* 
b2c下载订单 
alibaba.nlife.b2c.trade.download

下载零售商在零售+平台创建的订单
*/
func AlibabaNlifeB2cTradeDownload(clt *core.SDKClient, req *nlife.AlibabaNlifeB2cTradeDownloadRequest, session string) (*nlife.AlibabaNlifeB2cTradeDownloadAPIResponse, error) {
    var resp nlife.AlibabaNlifeB2cTradeDownloadAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
