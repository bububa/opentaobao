package lstlogistics

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/lstlogistics"
)

/* 
供应商-异云-第三方物流公司列表 
alibaba.lst.logistics.thirdpart.company.list

异地云仓发货时，需填写的第三方物流公司列表
*/
func AlibabaLstLogisticsThirdpartCompanyList(clt *core.SDKClient, req *lstlogistics.AlibabaLstLogisticsThirdpartCompanyListRequest, session string) (*lstlogistics.AlibabaLstLogisticsThirdpartCompanyListAPIResponse, error) {
    var resp lstlogistics.AlibabaLstLogisticsThirdpartCompanyListAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
