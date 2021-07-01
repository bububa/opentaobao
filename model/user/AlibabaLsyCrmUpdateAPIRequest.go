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
type AlibabaLsyCrmUpdateAPIRequest struct {
    model.Params
    // 更新客资状态对象
    _nrtUpdateRecordStatusDto   *NrtUpdateRecordStatusDto
}

// 初始化AlibabaLsyCrmUpdateAPIRequest对象
func NewAlibabaLsyCrmUpdateRequest() *AlibabaLsyCrmUpdateAPIRequest{
    return &AlibabaLsyCrmUpdateAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaLsyCrmUpdateAPIRequest) GetApiMethodName() string {
    return "alibaba.lsy.crm.update"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaLsyCrmUpdateAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// NrtUpdateRecordStatusDto Setter
// 更新客资状态对象
func (r *AlibabaLsyCrmUpdateAPIRequest) SetNrtUpdateRecordStatusDto(_nrtUpdateRecordStatusDto *NrtUpdateRecordStatusDto) error {
    r._nrtUpdateRecordStatusDto = _nrtUpdateRecordStatusDto
    r.Set("nrt_update_record_status_dto", _nrtUpdateRecordStatusDto)
    return nil
}

// NrtUpdateRecordStatusDto Getter
func (r AlibabaLsyCrmUpdateAPIRequest) GetNrtUpdateRecordStatusDto() *NrtUpdateRecordStatusDto {
    return r._nrtUpdateRecordStatusDto
}
