package icbu

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/icbu"
)

/* 
国际站图片银行查询接口 
alibaba.icbu.photobank.list

国际站图片银行查询接口
*/
func AlibabaIcbuPhotobankList(clt *core.SDKClient, req *icbu.AlibabaIcbuPhotobankListRequest, session string) (*icbu.AlibabaIcbuPhotobankListAPIResponse, error) {
    var resp icbu.AlibabaIcbuPhotobankListAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
