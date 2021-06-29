package aliyun

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询ODPS服务 API请求
odps.aliyuncs.com.DescribeOdpsService.2014-12-15

查询ODPS服务
*/
type OdpsAliyuncsComDescribeOdpsService2014_12_15Request struct {
    model.Params
}

// 初始化OdpsAliyuncsComDescribeOdpsService2014_12_15Request对象
func NewOdpsAliyuncsComDescribeOdpsService2014_12_15Request() *OdpsAliyuncsComDescribeOdpsService2014_12_15Request{
    return &OdpsAliyuncsComDescribeOdpsService2014_12_15Request{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r OdpsAliyuncsComDescribeOdpsService2014_12_15Request) GetApiMethodName() string {
    return "odps.aliyuncs.com.DescribeOdpsService.2014-12-15"
}

// IRequest interface 方法, 获取API参数
func (r OdpsAliyuncsComDescribeOdpsService2014_12_15Request) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
