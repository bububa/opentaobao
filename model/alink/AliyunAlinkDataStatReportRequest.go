package alink

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
外部离线统计数据上报 API请求
aliyun.alink.data.stat.report

外部合作厂商上报设备的明细数据，或者离线统计数据。
*/
type AliyunAlinkDataStatReportRequest struct {
    model.Params
    // 入参对象
    paramBean   *OuterDataBean
}

// 初始化AliyunAlinkDataStatReportRequest对象
func NewAliyunAlinkDataStatReportRequest() *AliyunAlinkDataStatReportRequest{
    return &AliyunAlinkDataStatReportRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AliyunAlinkDataStatReportRequest) GetApiMethodName() string {
    return "aliyun.alink.data.stat.report"
}

// IRequest interface 方法, 获取API参数
func (r AliyunAlinkDataStatReportRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamBean Setter
// 入参对象
func (r *AliyunAlinkDataStatReportRequest) SetParamBean(paramBean *OuterDataBean) error {
    r.paramBean = paramBean
    r.Set("param_bean", paramBean)
    return nil
}

// ParamBean Getter
func (r AliyunAlinkDataStatReportRequest) GetParamBean() *OuterDataBean {
    return r.paramBean
}
