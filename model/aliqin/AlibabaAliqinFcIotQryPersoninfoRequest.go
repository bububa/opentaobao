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
    _iccid   string
    // 指定查询某userid
    _userid   string
    // 由系统根据业务分配
    _midPatChannel   string
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
func (r *AlibabaAliqinFcIotQryPersoninfoRequest) SetIccid(_iccid string) error {
    r._iccid = _iccid
    r.Set("iccid", _iccid)
    return nil
}

// Iccid Getter
func (r AlibabaAliqinFcIotQryPersoninfoRequest) GetIccid() string {
    return r._iccid
}
// Userid Setter
// 指定查询某userid
func (r *AlibabaAliqinFcIotQryPersoninfoRequest) SetUserid(_userid string) error {
    r._userid = _userid
    r.Set("userid", _userid)
    return nil
}

// Userid Getter
func (r AlibabaAliqinFcIotQryPersoninfoRequest) GetUserid() string {
    return r._userid
}
// MidPatChannel Setter
// 由系统根据业务分配
func (r *AlibabaAliqinFcIotQryPersoninfoRequest) SetMidPatChannel(_midPatChannel string) error {
    r._midPatChannel = _midPatChannel
    r.Set("mid_pat_channel", _midPatChannel)
    return nil
}

// MidPatChannel Getter
func (r AlibabaAliqinFcIotQryPersoninfoRequest) GetMidPatChannel() string {
    return r._midPatChannel
}
