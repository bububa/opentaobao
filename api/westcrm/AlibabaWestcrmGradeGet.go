package westcrm

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/westcrm"
)

/* 
获取等级列表 
alibaba.westcrm.grade.get

获取会员卡等级列表
*/
func AlibabaWestcrmGradeGet(clt *core.SDKClient, req *westcrm.AlibabaWestcrmGradeGetAPIRequest, session string) (*westcrm.AlibabaWestcrmGradeGetAPIResponse, error) {
    var resp westcrm.AlibabaWestcrmGradeGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
