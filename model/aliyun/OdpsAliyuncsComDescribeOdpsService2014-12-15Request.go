package aliyun

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询ODPS服务 APIRequest
odps.aliyuncs.com.DescribeOdpsService.2014-12-15

查询ODPS服务
*/
type OdpsAliyuncsComDescribeOdpsService2014-12-15Request struct {
    model.Params

}

func NewOdpsAliyuncsComDescribeOdpsService2014-12-15Request() *OdpsAliyuncsComDescribeOdpsService2014-12-15Request{
    return &OdpsAliyuncsComDescribeOdpsService2014-12-15Request{
        Params: model.NewParams(),
    }
}

func (r OdpsAliyuncsComDescribeOdpsService2014-12-15Request) GetApiMethodName() string {
    return "odps.aliyuncs.com.DescribeOdpsService.2014-12-15"
}

func (r OdpsAliyuncsComDescribeOdpsService2014-12-15Request) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


