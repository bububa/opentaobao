package tbk

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tbk"
)

/* 
淘宝客-推广者-物料精选 
taobao.tbk.dg.optimus.material

支持入参对应的“推广位”和官方提供的“物料id”，获取指定物料信息和推广链接，还可入参用户信息提供智能推荐（需智能推荐请先前协议https://pub.alimama.com/fourth/protocol/common.htm?key=hangye_laxin）
*/
func TaobaoTbkDgOptimusMaterial(clt *core.SDKClient, req *tbk.TaobaoTbkDgOptimusMaterialRequest, session string) (*tbk.TaobaoTbkDgOptimusMaterialAPIResponse, error) {
    var resp tbk.TaobaoTbkDgOptimusMaterialAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
