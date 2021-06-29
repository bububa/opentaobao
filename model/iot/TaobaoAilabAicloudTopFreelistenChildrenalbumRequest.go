package iot

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
儿童音频列表 API请求
taobao.ailab.aicloud.top.freelisten.childrenalbum

儿童音频列表
*/
type TaobaoAilabAicloudTopFreelistenChildrenalbumRequest struct {
    model.Params
    // 账户体系隔离
    _schema   string
    // 用户ID，此处传入第三方账户体系的用户id
    _userId   string
    // 用户设备唯一识别码，长度限制32以内，建议使用系统接口获取deviceid,然后做一定的混淆处理来作为此输入参数
    _utdId   string
    // 扩展信息，用于存放APP类型等
    _ext   string
    // 音频类型，目前只支持以下几种类型：英语儿歌 英语故事 双语故事 国学启蒙 古典名著 成语故事 寓言故事 神话故事 诗词朗读 诗词婉唱 谚语故事 胎教音乐 经典儿歌 摇篮曲 睡前故事 绘本故事 儿童故事 儿童百科 经典故事 公主故事 名人故事 胎教故事
    _param1   string
    // 页数
    _param2   int64
    // 每页条目数
    _param3   int64
}

// 初始化TaobaoAilabAicloudTopFreelistenChildrenalbumRequest对象
func NewTaobaoAilabAicloudTopFreelistenChildrenalbumRequest() *TaobaoAilabAicloudTopFreelistenChildrenalbumRequest{
    return &TaobaoAilabAicloudTopFreelistenChildrenalbumRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAilabAicloudTopFreelistenChildrenalbumRequest) GetApiMethodName() string {
    return "taobao.ailab.aicloud.top.freelisten.childrenalbum"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAilabAicloudTopFreelistenChildrenalbumRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Schema Setter
// 账户体系隔离
func (r *TaobaoAilabAicloudTopFreelistenChildrenalbumRequest) SetSchema(_schema string) error {
    r._schema = _schema
    r.Set("schema", _schema)
    return nil
}

// Schema Getter
func (r TaobaoAilabAicloudTopFreelistenChildrenalbumRequest) GetSchema() string {
    return r._schema
}
// UserId Setter
// 用户ID，此处传入第三方账户体系的用户id
func (r *TaobaoAilabAicloudTopFreelistenChildrenalbumRequest) SetUserId(_userId string) error {
    r._userId = _userId
    r.Set("user_id", _userId)
    return nil
}

// UserId Getter
func (r TaobaoAilabAicloudTopFreelistenChildrenalbumRequest) GetUserId() string {
    return r._userId
}
// UtdId Setter
// 用户设备唯一识别码，长度限制32以内，建议使用系统接口获取deviceid,然后做一定的混淆处理来作为此输入参数
func (r *TaobaoAilabAicloudTopFreelistenChildrenalbumRequest) SetUtdId(_utdId string) error {
    r._utdId = _utdId
    r.Set("utd_id", _utdId)
    return nil
}

// UtdId Getter
func (r TaobaoAilabAicloudTopFreelistenChildrenalbumRequest) GetUtdId() string {
    return r._utdId
}
// Ext Setter
// 扩展信息，用于存放APP类型等
func (r *TaobaoAilabAicloudTopFreelistenChildrenalbumRequest) SetExt(_ext string) error {
    r._ext = _ext
    r.Set("ext", _ext)
    return nil
}

// Ext Getter
func (r TaobaoAilabAicloudTopFreelistenChildrenalbumRequest) GetExt() string {
    return r._ext
}
// Param1 Setter
// 音频类型，目前只支持以下几种类型：英语儿歌 英语故事 双语故事 国学启蒙 古典名著 成语故事 寓言故事 神话故事 诗词朗读 诗词婉唱 谚语故事 胎教音乐 经典儿歌 摇篮曲 睡前故事 绘本故事 儿童故事 儿童百科 经典故事 公主故事 名人故事 胎教故事
func (r *TaobaoAilabAicloudTopFreelistenChildrenalbumRequest) SetParam1(_param1 string) error {
    r._param1 = _param1
    r.Set("param1", _param1)
    return nil
}

// Param1 Getter
func (r TaobaoAilabAicloudTopFreelistenChildrenalbumRequest) GetParam1() string {
    return r._param1
}
// Param2 Setter
// 页数
func (r *TaobaoAilabAicloudTopFreelistenChildrenalbumRequest) SetParam2(_param2 int64) error {
    r._param2 = _param2
    r.Set("param2", _param2)
    return nil
}

// Param2 Getter
func (r TaobaoAilabAicloudTopFreelistenChildrenalbumRequest) GetParam2() int64 {
    return r._param2
}
// Param3 Setter
// 每页条目数
func (r *TaobaoAilabAicloudTopFreelistenChildrenalbumRequest) SetParam3(_param3 int64) error {
    r._param3 = _param3
    r.Set("param3", _param3)
    return nil
}

// Param3 Getter
func (r TaobaoAilabAicloudTopFreelistenChildrenalbumRequest) GetParam3() int64 {
    return r._param3
}
