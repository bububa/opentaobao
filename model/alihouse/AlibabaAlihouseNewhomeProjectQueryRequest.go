package alihouse

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询楼盘相关信息 API请求
alibaba.alihouse.newhome.project.query

根据outerid查询楼盘相关信息
*/
type AlibabaAlihouseNewhomeProjectQueryRequest struct {
    model.Params
    // 外部楼盘/小区id
    outerId   string
}

// 初始化AlibabaAlihouseNewhomeProjectQueryRequest对象
func NewAlibabaAlihouseNewhomeProjectQueryRequest() *AlibabaAlihouseNewhomeProjectQueryRequest{
    return &AlibabaAlihouseNewhomeProjectQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseNewhomeProjectQueryRequest) GetApiMethodName() string {
    return "alibaba.alihouse.newhome.project.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseNewhomeProjectQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OuterId Setter
// 外部楼盘/小区id
func (r *AlibabaAlihouseNewhomeProjectQueryRequest) SetOuterId(outerId string) error {
    r.outerId = outerId
    r.Set("outer_id", outerId)
    return nil
}

// OuterId Getter
func (r AlibabaAlihouseNewhomeProjectQueryRequest) GetOuterId() string {
    return r.outerId
}
