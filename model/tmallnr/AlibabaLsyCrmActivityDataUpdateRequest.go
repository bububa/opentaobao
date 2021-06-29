package tmallnr

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
私域导购数据回流接口 API请求
alibaba.lsy.crm.activity.data.update

私域导购数据回流接口
*/
type AlibabaLsyCrmActivityDataUpdateRequest struct {
    model.Params
    // 入参对象
    _reqDTO   *NrtCrmActivityStatisticsDataReq
}

// 初始化AlibabaLsyCrmActivityDataUpdateRequest对象
func NewAlibabaLsyCrmActivityDataUpdateRequest() *AlibabaLsyCrmActivityDataUpdateRequest{
    return &AlibabaLsyCrmActivityDataUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaLsyCrmActivityDataUpdateRequest) GetApiMethodName() string {
    return "alibaba.lsy.crm.activity.data.update"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaLsyCrmActivityDataUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ReqDTO Setter
// 入参对象
func (r *AlibabaLsyCrmActivityDataUpdateRequest) SetReqDTO(_reqDTO *NrtCrmActivityStatisticsDataReq) error {
    r._reqDTO = _reqDTO
    r.Set("req_d_t_o", _reqDTO)
    return nil
}

// ReqDTO Getter
func (r AlibabaLsyCrmActivityDataUpdateRequest) GetReqDTO() *NrtCrmActivityStatisticsDataReq {
    return r._reqDTO
}
