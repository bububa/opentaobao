package aliyunocs

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查看Region列表 API请求
m-kvstore.aliyuncs.com.DescribeRegions.2015-03-01

查看Region列表
*/
type M_kvstoreAliyuncsComDescribeRegions2015_03_01Request struct {
    model.Params
}

// 初始化M_kvstoreAliyuncsComDescribeRegions2015_03_01Request对象
func NewM_kvstoreAliyuncsComDescribeRegions2015_03_01Request() *M_kvstoreAliyuncsComDescribeRegions2015_03_01Request{
    return &M_kvstoreAliyuncsComDescribeRegions2015_03_01Request{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r M_kvstoreAliyuncsComDescribeRegions2015_03_01Request) GetApiMethodName() string {
    return "m-kvstore.aliyuncs.com.DescribeRegions.2015-03-01"
}

// IRequest interface 方法, 获取API参数
func (r M_kvstoreAliyuncsComDescribeRegions2015_03_01Request) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
