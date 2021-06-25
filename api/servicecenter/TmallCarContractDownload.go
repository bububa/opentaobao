package servicecenter

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/servicecenter"
)

/* 
合同下载 
tmall.car.contract.download

目前天猫开新车会在线上签署一份合同，协议，需要一个个在已卖出打开，另存为pdf，人工一个个下载比较麻烦，期望通过接口直接读取pdf；
因为比较耗时，建议一个个下载，假设并发下载，很可能限流，每天的调用量有限；
*/
func TmallCarContractDownload(clt *core.SDKClient, req *servicecenter.TmallCarContractDownloadRequest, session string) (*servicecenter.TmallCarContractDownloadResponse, error) {
    var resp servicecenter.TmallCarContractDownloadAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
