package aliyunocs

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查看Region列表 APIRequest
m-kvstore.aliyuncs.com.DescribeRegions.2015-03-01

查看Region列表
*/
type M-KvstoreAliyuncsComDescribeRegions2015-03-01Request struct {
    model.Params

}

func NewM-KvstoreAliyuncsComDescribeRegions2015-03-01Request() *M-KvstoreAliyuncsComDescribeRegions2015-03-01Request{
    return &M-KvstoreAliyuncsComDescribeRegions2015-03-01Request{
        Params: model.NewParams(),
    }
}

func (r M-KvstoreAliyuncsComDescribeRegions2015-03-01Request) GetApiMethodName() string {
    return "m-kvstore.aliyuncs.com.DescribeRegions.2015-03-01"
}

func (r M-KvstoreAliyuncsComDescribeRegions2015-03-01Request) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


