package tbk

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tbk"
)

/* 
淘宝客-公用-私域用户备案信息查询 
taobao.tbk.sc.publisher.info.get

查询已生成的渠道id或会员运营id的相关信息。
*/
func TaobaoTbkScPublisherInfoGet(clt *core.SDKClient, req *tbk.TaobaoTbkScPublisherInfoGetRequest, session string) (*tbk.TaobaoTbkScPublisherInfoGetAPIResponse, error) {
    var resp tbk.TaobaoTbkScPublisherInfoGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
