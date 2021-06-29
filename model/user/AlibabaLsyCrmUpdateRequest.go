package user

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
跟进客资状态接口 API请求
alibaba.lsy.crm.update

同步客资状态接口
*/
type AlibabaLsyCrmUpdateRequest struct {
    model.Params
    // 更新客资状态对象
    _nrtUpdateRecordStatusDto   *NrtUpdateRecordStatusDTO
}

// 初始化AlibabaLsyCrmUpdateRequest对象
func NewAlibabaLsyCrmUpdateRequest() *AlibabaLsyCrmUpdateRequest{
    return &AlibabaLsyCrmUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaLsyCrmUpdateRequest) GetApiMethodName() string {
    return "alibaba.lsy.crm.update"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaLsyCrmUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// NrtUpdateRecordStatusDto Setter
// 更新客资状态对象
func (r *AlibabaLsyCrmUpdateRequest) SetNrtUpdateRecordStatusDto(_nrtUpdateRecordStatusDto *NrtUpdateRecordStatusDTO) error {
    r._nrtUpdateRecordStatusDto = _nrtUpdateRecordStatusDto
    r.Set("nrt_update_record_status_dto", _nrtUpdateRecordStatusDto)
    return nil
}

// NrtUpdateRecordStatusDto Getter
func (r AlibabaLsyCrmUpdateRequest) GetNrtUpdateRecordStatusDto() *NrtUpdateRecordStatusDTO {
    return r._nrtUpdateRecordStatusDto
}
