package alsc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询卡模板列表(支持数据下行) API请求
alibaba.alsc.crm.card.pagetmp

查询卡模板列表(支持数据下行)
当传递了数据下行参数:
     * isDeleted,lastMaxId,gmtModified,num时,进行数据下行处理,返回结果不带分页信息
     * 否则分页查询卡模板,返回结果带有分页信息
*/
type AlibabaAlscCrmCardPagetmpRequest struct {
    model.Params
    // 请求结果
    _paramPullCardTemplateOpenReq   *PullCardTemplateOpenReq
}

// 初始化AlibabaAlscCrmCardPagetmpRequest对象
func NewAlibabaAlscCrmCardPagetmpRequest() *AlibabaAlscCrmCardPagetmpRequest{
    return &AlibabaAlscCrmCardPagetmpRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlscCrmCardPagetmpRequest) GetApiMethodName() string {
    return "alibaba.alsc.crm.card.pagetmp"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlscCrmCardPagetmpRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamPullCardTemplateOpenReq Setter
// 请求结果
func (r *AlibabaAlscCrmCardPagetmpRequest) SetParamPullCardTemplateOpenReq(_paramPullCardTemplateOpenReq *PullCardTemplateOpenReq) error {
    r._paramPullCardTemplateOpenReq = _paramPullCardTemplateOpenReq
    r.Set("param_pull_card_template_open_req", _paramPullCardTemplateOpenReq)
    return nil
}

// ParamPullCardTemplateOpenReq Getter
func (r AlibabaAlscCrmCardPagetmpRequest) GetParamPullCardTemplateOpenReq() *PullCardTemplateOpenReq {
    return r._paramPullCardTemplateOpenReq
}
