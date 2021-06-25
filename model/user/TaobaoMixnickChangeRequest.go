package user

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
新旧mixnick互转 APIRequest
taobao.mixnick.change

如果采用老的Appkey获取MixNick A’， 发现DB里 new MixNick为null，则使用新的Appkey调API来换取A’’；如果采用新的Appkey获取A’’，发现DB里不存在当前A’’ 的记录时，先用老Appkey调API得到A’ 查询是否存在A用户的记录，如果已经存在，则关联，否则新增一条MixNick为null的新用户记录。
*/
type TaobaoMixnickChangeRequest struct {
    model.Params

    // 输入的混淆nick
    srcMixNick   string 

    // 输入的appkey
    srcAppkey   string 

}

func NewTaobaoMixnickChangeRequest() *TaobaoMixnickChangeRequest{
    return &TaobaoMixnickChangeRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoMixnickChangeRequest) GetApiMethodName() string {
    return "taobao.mixnick.change"
}

func (r TaobaoMixnickChangeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoMixnickChangeRequest) SetSrcMixNick(srcMixNick string) error {
    r.srcMixNick = srcMixNick
    r.Set("src_mix_nick", srcMixNick)
    return nil
}

func (r TaobaoMixnickChangeRequest) GetSrcMixNick() string {
    return r.srcMixNick
}

func (r *TaobaoMixnickChangeRequest) SetSrcAppkey(srcAppkey string) error {
    r.srcAppkey = srcAppkey
    r.Set("src_appkey", srcAppkey)
    return nil
}

func (r TaobaoMixnickChangeRequest) GetSrcAppkey() string {
    return r.srcAppkey
}

