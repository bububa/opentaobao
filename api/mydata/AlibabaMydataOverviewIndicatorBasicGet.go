package mydata

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/mydata"
)

/* 
我的效果-获取公司询盘流量行业表现 
alibaba.mydata.overview.indicator.basic.get

获取公司询盘流量行业表现
*/
func AlibabaMydataOverviewIndicatorBasicGet(clt *core.SDKClient, req *mydata.AlibabaMydataOverviewIndicatorBasicGetRequest, session string) (*mydata.AlibabaMydataOverviewIndicatorBasicGetAPIResponse, error) {
    var resp mydata.AlibabaMydataOverviewIndicatorBasicGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
