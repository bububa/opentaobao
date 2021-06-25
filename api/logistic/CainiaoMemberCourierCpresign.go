package logistic

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/logistic"
)

/* 
cp清理离职用户信息 
cainiao.member.courier.cpresign

CP清理内部离职的用户信息
*/
func CainiaoMemberCourierCpresign(clt *core.SDKClient, req *logistic.CainiaoMemberCourierCpresignRequest, session string) (*logistic.CainiaoMemberCourierCpresignResponse, error) {
    var resp logistic.CainiaoMemberCourierCpresignAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
