package aliqin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询物联卡个人实人认证信息 API请求
alibaba.aliqin.fc.iot.qry.personinfo

查询物联卡个人实人认证信息
*/
type AlibabaAliqinFcIotQryPersoninfoRequest struct {
    model.Params
    // 需要查询的iccid
    iccid   string
    // 指定查询某userid
    userid   string
    // 由系统根据业务分配
    midPatChannel   string
}

// 初始化AlibabaAliqinFcIotQryPersoninfoRequest对象
func NewAlibabaAliqinFcIotQryPersoninfoRequest() *AlibabaAliqinFcIotQryPersoninfoRequest{
    return &AlibabaAliqinFcIotQryPersoninfoRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAliqinFcIotQryPersoninfoRequest) GetApiMethodName() string {
    return "alibaba.aliqin.fc.iot.qry.personinfo"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAliqinFcIotQryPersoninfoRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Iccid Setter
// 需要查询的iccid
func (r *AlibabaAliqinFcIotQryPersoninfoRequest) SetIccid(iccid string) error {
    r.iccid = iccid
    r.Set("iccid", iccid)
    return nil
}

// Iccid Getter
func (r AlibabaAliqinFcIotQryPersoninfoRequest) GetIccid() string {
    return r.iccid
}
// Userid Setter
// 指定查询某userid
func (r *AlibabaAliqinFcIotQryPersoninfoRequest) SetUserid(userid string) error {
    r.userid = userid
    r.Set("userid", userid)
    return nil
}

// Userid Getter
func (r AlibabaAliqinFcIotQryPersoninfoRequest) GetUserid() string {
    return r.userid
}
// MidPatChannel Setter
// 由系统根据业务分配
func (r *AlibabaAliqinFcIotQryPersoninfoRequest) SetMidPatChannel(midPatChannel string) error {
    r.midPatChannel = midPatChannel
    r.Set("mid_pat_channel", midPatChannel)
    return nil
}

// MidPatChannel Getter
func (r AlibabaAliqinFcIotQryPersoninfoRequest) GetMidPatChannel() string {
    return r.midPatChannel
}
