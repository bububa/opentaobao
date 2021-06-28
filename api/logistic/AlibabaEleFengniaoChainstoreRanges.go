package logistic

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/logistic"
)

/* 
蜂鸟查询门店配送范围接口 
alibaba.ele.fengniao.chainstore.ranges

蜂鸟查询门店配送范围接口
*/
func AlibabaEleFengniaoChainstoreRanges(clt *core.SDKClient, req *logistic.AlibabaEleFengniaoChainstoreRangesRequest, session string) (*logistic.AlibabaEleFengniaoChainstoreRangesAPIResponse, error) {
    var resp logistic.AlibabaEleFengniaoChainstoreRangesAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
