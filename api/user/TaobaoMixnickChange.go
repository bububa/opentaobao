package user

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/user"
)

/* 
新旧mixnick互转 
taobao.mixnick.change

如果采用老的Appkey获取MixNick A’， 发现DB里 new MixNick为null，则使用新的Appkey调API来换取A’’；如果采用新的Appkey获取A’’，发现DB里不存在当前A’’ 的记录时，先用老Appkey调API得到A’ 查询是否存在A用户的记录，如果已经存在，则关联，否则新增一条MixNick为null的新用户记录。
*/
func TaobaoMixnickChange(clt *core.SDKClient, req *user.TaobaoMixnickChangeRequest, session string) (*user.TaobaoMixnickChangeResponse, error) {
    var resp user.TaobaoMixnickChangeAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
