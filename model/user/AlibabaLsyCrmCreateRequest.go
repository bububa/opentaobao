package user

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
创建客资 API请求
alibaba.lsy.crm.create

欢客调用该接口进行客资创建
*/
type AlibabaLsyCrmCreateRequest struct {
    model.Params
    // 客资记录对象
    _nrtRecordDto   *NrtRecordDto
}

// 初始化AlibabaLsyCrmCreateRequest对象
func NewAlibabaLsyCrmCreateRequest() *AlibabaLsyCrmCreateRequest{
    return &AlibabaLsyCrmCreateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaLsyCrmCreateRequest) GetApiMethodName() string {
    return "alibaba.lsy.crm.create"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaLsyCrmCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// NrtRecordDto Setter
// 客资记录对象
func (r *AlibabaLsyCrmCreateRequest) SetNrtRecordDto(_nrtRecordDto *NrtRecordDto) error {
    r._nrtRecordDto = _nrtRecordDto
    r.Set("nrt_record_dto", _nrtRecordDto)
    return nil
}

// NrtRecordDto Getter
func (r AlibabaLsyCrmCreateRequest) GetNrtRecordDto() *NrtRecordDto {
    return r._nrtRecordDto
}
