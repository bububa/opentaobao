package alink

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
外部离线统计数据上报 APIRequest
aliyun.alink.data.stat.report

外部合作厂商上报设备的明细数据，或者离线统计数据。
*/
type AliyunAlinkDataStatReportRequest struct {
    model.Params

    // 入参对象
    paramBean   *OuterDataBean 

}

func NewAliyunAlinkDataStatReportRequest() *AliyunAlinkDataStatReportRequest{
    return &AliyunAlinkDataStatReportRequest{
        Params: model.NewParams(),
    }
}

func (r AliyunAlinkDataStatReportRequest) GetApiMethodName() string {
    return "aliyun.alink.data.stat.report"
}

func (r AliyunAlinkDataStatReportRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AliyunAlinkDataStatReportRequest) SetParamBean(paramBean *OuterDataBean) error {
    r.paramBean = paramBean
    r.Set("param_bean", paramBean)
    return nil
}

func (r AliyunAlinkDataStatReportRequest) GetParamBean() *OuterDataBean {
    return r.paramBean
}

