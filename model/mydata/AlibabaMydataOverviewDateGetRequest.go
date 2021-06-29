package mydata

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
我的效果-获取数据周期 API请求
alibaba.mydata.overview.date.get

获取数据管家我的效果API可以使用的数据周期
*/
type AlibabaMydataOverviewDateGetRequest struct {
    model.Params
}

// 初始化AlibabaMydataOverviewDateGetRequest对象
func NewAlibabaMydataOverviewDateGetRequest() *AlibabaMydataOverviewDateGetRequest{
    return &AlibabaMydataOverviewDateGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaMydataOverviewDateGetRequest) GetApiMethodName() string {
    return "alibaba.mydata.overview.date.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaMydataOverviewDateGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
