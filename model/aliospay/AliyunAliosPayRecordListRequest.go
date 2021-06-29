package aliospay

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
支付记录批量查询接口 API请求
aliyun.alios.pay.record.list

商户用来对账的接口
*/
type AliyunAliosPayRecordListRequest struct {
    model.Params
    // 请求参数
    _searchRecordRequest   *SearchRecordRequest
}

// 初始化AliyunAliosPayRecordListRequest对象
func NewAliyunAliosPayRecordListRequest() *AliyunAliosPayRecordListRequest{
    return &AliyunAliosPayRecordListRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AliyunAliosPayRecordListRequest) GetApiMethodName() string {
    return "aliyun.alios.pay.record.list"
}

// IRequest interface 方法, 获取API参数
func (r AliyunAliosPayRecordListRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SearchRecordRequest Setter
// 请求参数
func (r *AliyunAliosPayRecordListRequest) SetSearchRecordRequest(_searchRecordRequest *SearchRecordRequest) error {
    r._searchRecordRequest = _searchRecordRequest
    r.Set("search_record_request", _searchRecordRequest)
    return nil
}

// SearchRecordRequest Getter
func (r AliyunAliosPayRecordListRequest) GetSearchRecordRequest() *SearchRecordRequest {
    return r._searchRecordRequest
}
