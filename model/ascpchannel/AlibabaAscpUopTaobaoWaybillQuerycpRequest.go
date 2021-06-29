package ascpchannel

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询电子面单开放的CP列表 API请求
alibaba.ascp.uop.taobao.waybill.querycp

查询电子面单开放的CP列表
*/
type AlibabaAscpUopTaobaoWaybillQuerycpRequest struct {
    model.Params
    // 系统自动生成
    _queryCpRequest   *Querycprequest
}

// 初始化AlibabaAscpUopTaobaoWaybillQuerycpRequest对象
func NewAlibabaAscpUopTaobaoWaybillQuerycpRequest() *AlibabaAscpUopTaobaoWaybillQuerycpRequest{
    return &AlibabaAscpUopTaobaoWaybillQuerycpRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAscpUopTaobaoWaybillQuerycpRequest) GetApiMethodName() string {
    return "alibaba.ascp.uop.taobao.waybill.querycp"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAscpUopTaobaoWaybillQuerycpRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// QueryCpRequest Setter
// 系统自动生成
func (r *AlibabaAscpUopTaobaoWaybillQuerycpRequest) SetQueryCpRequest(_queryCpRequest *Querycprequest) error {
    r._queryCpRequest = _queryCpRequest
    r.Set("query_cp_request", _queryCpRequest)
    return nil
}

// QueryCpRequest Getter
func (r AlibabaAscpUopTaobaoWaybillQuerycpRequest) GetQueryCpRequest() *Querycprequest {
    return r._queryCpRequest
}
