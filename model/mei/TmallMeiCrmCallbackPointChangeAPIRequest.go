package mei

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
品牌积分变更回调API API请求
tmall.mei.crm.callback.point.change

线下品牌积分变更消息回调API，告诉积分扣减或者累加是否成功。
*/
type TmallMeiCrmCallbackPointChangeAPIRequest struct {
    model.Params
    // 混淆会员手机号码
    _mixMobile   string
    // 变更记录ID
    _recordId   int64
    // 0:成功。1：失败
    _result   int64
    // 处理失败的错误码.
    _errorCode   string
    // 拓展信息
    _extInfo   string
    // 积分信息
    _point   int64
}

// 初始化TmallMeiCrmCallbackPointChangeAPIRequest对象
func NewTmallMeiCrmCallbackPointChangeRequest() *TmallMeiCrmCallbackPointChangeAPIRequest{
    return &TmallMeiCrmCallbackPointChangeAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallMeiCrmCallbackPointChangeAPIRequest) GetApiMethodName() string {
    return "tmall.mei.crm.callback.point.change"
}

// IRequest interface 方法, 获取API参数
func (r TmallMeiCrmCallbackPointChangeAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// MixMobile Setter
// 混淆会员手机号码
func (r *TmallMeiCrmCallbackPointChangeAPIRequest) SetMixMobile(_mixMobile string) error {
    r._mixMobile = _mixMobile
    r.Set("mix_mobile", _mixMobile)
    return nil
}

// MixMobile Getter
func (r TmallMeiCrmCallbackPointChangeAPIRequest) GetMixMobile() string {
    return r._mixMobile
}
// RecordId Setter
// 变更记录ID
func (r *TmallMeiCrmCallbackPointChangeAPIRequest) SetRecordId(_recordId int64) error {
    r._recordId = _recordId
    r.Set("record_id", _recordId)
    return nil
}

// RecordId Getter
func (r TmallMeiCrmCallbackPointChangeAPIRequest) GetRecordId() int64 {
    return r._recordId
}
// Result Setter
// 0:成功。1：失败
func (r *TmallMeiCrmCallbackPointChangeAPIRequest) SetResult(_result int64) error {
    r._result = _result
    r.Set("result", _result)
    return nil
}

// Result Getter
func (r TmallMeiCrmCallbackPointChangeAPIRequest) GetResult() int64 {
    return r._result
}
// ErrorCode Setter
// 处理失败的错误码.
func (r *TmallMeiCrmCallbackPointChangeAPIRequest) SetErrorCode(_errorCode string) error {
    r._errorCode = _errorCode
    r.Set("error_code", _errorCode)
    return nil
}

// ErrorCode Getter
func (r TmallMeiCrmCallbackPointChangeAPIRequest) GetErrorCode() string {
    return r._errorCode
}
// ExtInfo Setter
// 拓展信息
func (r *TmallMeiCrmCallbackPointChangeAPIRequest) SetExtInfo(_extInfo string) error {
    r._extInfo = _extInfo
    r.Set("ext_info", _extInfo)
    return nil
}

// ExtInfo Getter
func (r TmallMeiCrmCallbackPointChangeAPIRequest) GetExtInfo() string {
    return r._extInfo
}
// Point Setter
// 积分信息
func (r *TmallMeiCrmCallbackPointChangeAPIRequest) SetPoint(_point int64) error {
    r._point = _point
    r.Set("point", _point)
    return nil
}

// Point Getter
func (r TmallMeiCrmCallbackPointChangeAPIRequest) GetPoint() int64 {
    return r._point
}
