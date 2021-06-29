package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
同城零售履约异常中心异常单处理结果回调接口 API请求
tmall.cityretail.fulfill.abnormal.center.abnormal.status.change

同城零售履约异常中心异常单处理结果回调接口
*/
type TmallCityretailFulfillAbnormalCenterAbnormalStatusChangeRequest struct {
    model.Params
    // 入参
    _abnormalStatusChangeDto   []AbnormalStatusChangeDTO
}

// 初始化TmallCityretailFulfillAbnormalCenterAbnormalStatusChangeRequest对象
func NewTmallCityretailFulfillAbnormalCenterAbnormalStatusChangeRequest() *TmallCityretailFulfillAbnormalCenterAbnormalStatusChangeRequest{
    return &TmallCityretailFulfillAbnormalCenterAbnormalStatusChangeRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallCityretailFulfillAbnormalCenterAbnormalStatusChangeRequest) GetApiMethodName() string {
    return "tmall.cityretail.fulfill.abnormal.center.abnormal.status.change"
}

// IRequest interface 方法, 获取API参数
func (r TmallCityretailFulfillAbnormalCenterAbnormalStatusChangeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AbnormalStatusChangeDto Setter
// 入参
func (r *TmallCityretailFulfillAbnormalCenterAbnormalStatusChangeRequest) SetAbnormalStatusChangeDto(_abnormalStatusChangeDto []AbnormalStatusChangeDTO) error {
    r._abnormalStatusChangeDto = _abnormalStatusChangeDto
    r.Set("abnormal_status_change_dto", _abnormalStatusChangeDto)
    return nil
}

// AbnormalStatusChangeDto Getter
func (r TmallCityretailFulfillAbnormalCenterAbnormalStatusChangeRequest) GetAbnormalStatusChangeDto() []AbnormalStatusChangeDTO {
    return r._abnormalStatusChangeDto
}
