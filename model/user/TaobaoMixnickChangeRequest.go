package user

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
新旧mixnick互转 API请求
taobao.mixnick.change

如果采用老的Appkey获取MixNick A’， 发现DB里 new MixNick为null，则使用新的Appkey调API来换取A’’；如果采用新的Appkey获取A’’，发现DB里不存在当前A’’ 的记录时，先用老Appkey调API得到A’ 查询是否存在A用户的记录，如果已经存在，则关联，否则新增一条MixNick为null的新用户记录。
*/
type TaobaoMixnickChangeRequest struct {
    model.Params
    // 输入的混淆nick
    _srcMixNick   string
    // 输入的appkey
    _srcAppkey   string
}

// 初始化TaobaoMixnickChangeRequest对象
func NewTaobaoMixnickChangeRequest() *TaobaoMixnickChangeRequest{
    return &TaobaoMixnickChangeRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoMixnickChangeRequest) GetApiMethodName() string {
    return "taobao.mixnick.change"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoMixnickChangeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SrcMixNick Setter
// 输入的混淆nick
func (r *TaobaoMixnickChangeRequest) SetSrcMixNick(_srcMixNick string) error {
    r._srcMixNick = _srcMixNick
    r.Set("src_mix_nick", _srcMixNick)
    return nil
}

// SrcMixNick Getter
func (r TaobaoMixnickChangeRequest) GetSrcMixNick() string {
    return r._srcMixNick
}
// SrcAppkey Setter
// 输入的appkey
func (r *TaobaoMixnickChangeRequest) SetSrcAppkey(_srcAppkey string) error {
    r._srcAppkey = _srcAppkey
    r.Set("src_appkey", _srcAppkey)
    return nil
}

// SrcAppkey Getter
func (r TaobaoMixnickChangeRequest) GetSrcAppkey() string {
    return r._srcAppkey
}
