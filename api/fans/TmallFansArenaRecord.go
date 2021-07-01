package fans

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/fans"
)

/* 
记录完成擂台的用户 
tmall.fans.arena.record

记录完成擂台的用户和完成分数
*/
func TmallFansArenaRecord(clt *core.SDKClient, req *fans.TmallFansArenaRecordAPIRequest, session string) (*fans.TmallFansArenaRecordAPIResponse, error) {
    var resp fans.TmallFansArenaRecordAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
