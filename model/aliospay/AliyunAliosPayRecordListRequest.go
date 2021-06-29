package aliospay

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
支付记录批量查询接口 APIRequest
aliyun.alios.pay.record.list

商户用来对账的接口
*/
type AliyunAliosPayRecordListRequest struct {
    model.Params

    // 请求参数
    searchRecordRequest   *SearchRecordRequest 

}

func NewAliyunAliosPayRecordListRequest() *AliyunAliosPayRecordListRequest{
    return &AliyunAliosPayRecordListRequest{
        Params: model.NewParams(),
    }
}

func (r AliyunAliosPayRecordListRequest) GetApiMethodName() string {
    return "aliyun.alios.pay.record.list"
}

func (r AliyunAliosPayRecordListRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AliyunAliosPayRecordListRequest) SetSearchRecordRequest(searchRecordRequest *SearchRecordRequest) error {
    r.searchRecordRequest = searchRecordRequest
    r.Set("search_record_request", searchRecordRequest)
    return nil
}

func (r AliyunAliosPayRecordListRequest) GetSearchRecordRequest() *SearchRecordRequest {
    return r.searchRecordRequest
}

