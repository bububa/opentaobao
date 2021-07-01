package brandhub

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/brandhub"
)

/* 
品牌号品牌特秀单元报表数据查询 
taobao.brandhub.specialshow.rpt.adgroup.get

获取品牌号品牌特秀广告adgroup分日报表数据，只能查询近90天内的数据，包括展现量，点击量等
*/
func TaobaoBrandhubSpecialshowRptAdgroupGet(clt *core.SDKClient, req *brandhub.TaobaoBrandhubSpecialshowRptAdgroupGetAPIRequest, session string) (*brandhub.TaobaoBrandhubSpecialshowRptAdgroupGetAPIResponse, error) {
    var resp brandhub.TaobaoBrandhubSpecialshowRptAdgroupGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
