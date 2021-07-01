package xhotelcrm

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/xhotelcrm"
)

/* 
飞猪酒店商家会员绑定 
taobao.xhotel.potential.member.bind

支持互通商家发起会员绑定
*/
func TaobaoXhotelPotentialMemberBind(clt *core.SDKClient, req *xhotelcrm.TaobaoXhotelPotentialMemberBindAPIRequest, session string) (*xhotelcrm.TaobaoXhotelPotentialMemberBindAPIResponse, error) {
    var resp xhotelcrm.TaobaoXhotelPotentialMemberBindAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
