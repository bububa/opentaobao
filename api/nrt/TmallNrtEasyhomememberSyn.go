package nrt

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/nrt"
)

/* 
会员信息同 
tmall.nrt.easyhomemember.syn

居然之家将会员信息同步到零售中台 包含基本的会员信息
*/
func TmallNrtEasyhomememberSyn(clt *core.SDKClient, req *nrt.TmallNrtEasyhomememberSynRequest, session string) (*nrt.TmallNrtEasyhomememberSynAPIResponse, error) {
    var resp nrt.TmallNrtEasyhomememberSynAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
