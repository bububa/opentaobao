package aliexpresssumaitong

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/aliexpresssumaitong"
)

/* 
关务所需的申报清关字段 
aliexpress.taxation.calculate.open.query

关务所需的申报清关字段
*/
func AliexpressTaxationCalculateOpenQuery(clt *core.SDKClient, req *aliexpresssumaitong.AliexpressTaxationCalculateOpenQueryRequest, session string) (*aliexpresssumaitong.AliexpressTaxationCalculateOpenQueryAPIResponse, error) {
    var resp aliexpresssumaitong.AliexpressTaxationCalculateOpenQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
