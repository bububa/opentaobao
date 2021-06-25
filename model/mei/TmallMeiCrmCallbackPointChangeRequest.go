package mei

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
品牌积分变更回调API APIRequest
tmall.mei.crm.callback.point.change

线下品牌积分变更消息回调API，告诉积分扣减或者累加是否成功。
*/
type TmallMeiCrmCallbackPointChangeRequest struct {
    model.Params

    // 混淆会员手机号码
    mixMobile   string 

    // 变更记录ID
    recordId   int64 

    // 0:成功。1：失败
    result   int64 

    // 处理失败的错误码.
    errorCode   string 

    // 拓展信息
    extInfo   string 

    // 积分信息
    point   int64 

}

func NewTmallMeiCrmCallbackPointChangeRequest() *TmallMeiCrmCallbackPointChangeRequest{
    return &TmallMeiCrmCallbackPointChangeRequest{
        Params: model.NewParams(),
    }
}

func (r TmallMeiCrmCallbackPointChangeRequest) GetApiMethodName() string {
    return "tmall.mei.crm.callback.point.change"
}

func (r TmallMeiCrmCallbackPointChangeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallMeiCrmCallbackPointChangeRequest) SetMixMobile(mixMobile string) error {
    r.mixMobile = mixMobile
    r.Set("mix_mobile", mixMobile)
    return nil
}

func (r TmallMeiCrmCallbackPointChangeRequest) GetMixMobile() string {
    return r.mixMobile
}

func (r *TmallMeiCrmCallbackPointChangeRequest) SetRecordId(recordId int64) error {
    r.recordId = recordId
    r.Set("record_id", recordId)
    return nil
}

func (r TmallMeiCrmCallbackPointChangeRequest) GetRecordId() int64 {
    return r.recordId
}

func (r *TmallMeiCrmCallbackPointChangeRequest) SetResult(result int64) error {
    r.result = result
    r.Set("result", result)
    return nil
}

func (r TmallMeiCrmCallbackPointChangeRequest) GetResult() int64 {
    return r.result
}

func (r *TmallMeiCrmCallbackPointChangeRequest) SetErrorCode(errorCode string) error {
    r.errorCode = errorCode
    r.Set("error_code", errorCode)
    return nil
}

func (r TmallMeiCrmCallbackPointChangeRequest) GetErrorCode() string {
    return r.errorCode
}

func (r *TmallMeiCrmCallbackPointChangeRequest) SetExtInfo(extInfo string) error {
    r.extInfo = extInfo
    r.Set("ext_info", extInfo)
    return nil
}

func (r TmallMeiCrmCallbackPointChangeRequest) GetExtInfo() string {
    return r.extInfo
}

func (r *TmallMeiCrmCallbackPointChangeRequest) SetPoint(point int64) error {
    r.point = point
    r.Set("point", point)
    return nil
}

func (r TmallMeiCrmCallbackPointChangeRequest) GetPoint() int64 {
    return r.point
}

