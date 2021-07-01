package scbp

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/scbp"
)

/* 
查询指定产品分组的下一层子分组 
alibaba.scbp.product.group.get

查询指定产品分组的下一层子分组
*/
func AlibabaScbpProductGroupGet(clt *core.SDKClient, req *scbp.AlibabaScbpProductGroupGetAPIRequest, session string) (*scbp.AlibabaScbpProductGroupGetAPIResponse, error) {
    var resp scbp.AlibabaScbpProductGroupGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
