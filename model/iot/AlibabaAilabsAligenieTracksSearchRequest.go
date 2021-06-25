package iot

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询音频 APIRequest
alibaba.ailabs.aligenie.tracks.search

搜索类目下的音频信息
*/
type AlibabaAilabsAligenieTracksSearchRequest struct {
    model.Params

    // 账户体系隔离
    schema   string 

    // 用户ID，此处传入第三方账户体系的用户id
    userId   string 

    // 用户设备唯一识别码，长度限制32以内，建议使用系统接口获取deviceid,然后做一定的混淆处理来作为此输入参数
    utdId   string 

    // 扩展信息，用于存放APP类型等
    ext   string 

    // 一级类目，如：儿童、新闻、商业财经有声书等
    param1   string 

    // 二级类目，如儿童下有：儿歌、童谣、国学等
    param2   string 

    // 搜索的单个音频名称
    param3   string 

    // 当前页（从1开始, 目前由于底层引擎限制，实际最多能查5000个，超过5000返回为空，请确保页码*分页大小不超过5000）
    param4   int64 

    // 每页数量（不超过50）
    param5   int64 

}

func NewAlibabaAilabsAligenieTracksSearchRequest() *AlibabaAilabsAligenieTracksSearchRequest{
    return &AlibabaAilabsAligenieTracksSearchRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAilabsAligenieTracksSearchRequest) GetApiMethodName() string {
    return "alibaba.ailabs.aligenie.tracks.search"
}

func (r AlibabaAilabsAligenieTracksSearchRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAilabsAligenieTracksSearchRequest) SetSchema(schema string) error {
    r.schema = schema
    r.Set("schema", schema)
    return nil
}

func (r AlibabaAilabsAligenieTracksSearchRequest) GetSchema() string {
    return r.schema
}

func (r *AlibabaAilabsAligenieTracksSearchRequest) SetUserId(userId string) error {
    r.userId = userId
    r.Set("user_id", userId)
    return nil
}

func (r AlibabaAilabsAligenieTracksSearchRequest) GetUserId() string {
    return r.userId
}

func (r *AlibabaAilabsAligenieTracksSearchRequest) SetUtdId(utdId string) error {
    r.utdId = utdId
    r.Set("utd_id", utdId)
    return nil
}

func (r AlibabaAilabsAligenieTracksSearchRequest) GetUtdId() string {
    return r.utdId
}

func (r *AlibabaAilabsAligenieTracksSearchRequest) SetExt(ext string) error {
    r.ext = ext
    r.Set("ext", ext)
    return nil
}

func (r AlibabaAilabsAligenieTracksSearchRequest) GetExt() string {
    return r.ext
}

func (r *AlibabaAilabsAligenieTracksSearchRequest) SetParam1(param1 string) error {
    r.param1 = param1
    r.Set("param1", param1)
    return nil
}

func (r AlibabaAilabsAligenieTracksSearchRequest) GetParam1() string {
    return r.param1
}

func (r *AlibabaAilabsAligenieTracksSearchRequest) SetParam2(param2 string) error {
    r.param2 = param2
    r.Set("param2", param2)
    return nil
}

func (r AlibabaAilabsAligenieTracksSearchRequest) GetParam2() string {
    return r.param2
}

func (r *AlibabaAilabsAligenieTracksSearchRequest) SetParam3(param3 string) error {
    r.param3 = param3
    r.Set("param3", param3)
    return nil
}

func (r AlibabaAilabsAligenieTracksSearchRequest) GetParam3() string {
    return r.param3
}

func (r *AlibabaAilabsAligenieTracksSearchRequest) SetParam4(param4 int64) error {
    r.param4 = param4
    r.Set("param4", param4)
    return nil
}

func (r AlibabaAilabsAligenieTracksSearchRequest) GetParam4() int64 {
    return r.param4
}

func (r *AlibabaAilabsAligenieTracksSearchRequest) SetParam5(param5 int64) error {
    r.param5 = param5
    r.Set("param5", param5)
    return nil
}

func (r AlibabaAilabsAligenieTracksSearchRequest) GetParam5() int64 {
    return r.param5
}

