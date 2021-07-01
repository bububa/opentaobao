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
type AlibabaLsyCrmCreateAPIRequest struct {
    model.Params
    // 客资记录对象
    _nrtRecordDto   *NrtRecordDTO
}

// 初始化AlibabaLsyCrmCreateAPIRequest对象
func NewAlibabaLsyCrmCreateRequest() *AlibabaLsyCrmCreateAPIRequest{
    return &AlibabaLsyCrmCreateAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaLsyCrmCreateAPIRequest) GetApiMethodName() string {
    return "alibaba.lsy.crm.create"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaLsyCrmCreateAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// NrtRecordDto Setter
// 客资记录对象
func (r *AlibabaLsyCrmCreateAPIRequest) SetNrtRecordDto(_nrtRecordDto *NrtRecordDTO) error {
    r._nrtRecordDto = _nrtRecordDto
    r.Set("nrt_record_dto", _nrtRecordDto)
    return nil
}

// NrtRecordDto Getter
func (r AlibabaLsyCrmCreateAPIRequest) GetNrtRecordDto() *NrtRecordDTO {
    return r._nrtRecordDto
}
