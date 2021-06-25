package nazca

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取文件秘钥 APIRequest
alibaba.nazca.token.filesecret.get

获取文件秘钥
*/
type AlibabaNazcaTokenFilesecretGetRequest struct {
    model.Params

    // 客户在1688的唯一标识
    platformUserId   string 

    // 合同编号
    contractNum   string 

}

func NewAlibabaNazcaTokenFilesecretGetRequest() *AlibabaNazcaTokenFilesecretGetRequest{
    return &AlibabaNazcaTokenFilesecretGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaNazcaTokenFilesecretGetRequest) GetApiMethodName() string {
    return "alibaba.nazca.token.filesecret.get"
}

func (r AlibabaNazcaTokenFilesecretGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaNazcaTokenFilesecretGetRequest) SetPlatformUserId(platformUserId string) error {
    r.platformUserId = platformUserId
    r.Set("platform_user_id", platformUserId)
    return nil
}

func (r AlibabaNazcaTokenFilesecretGetRequest) GetPlatformUserId() string {
    return r.platformUserId
}

func (r *AlibabaNazcaTokenFilesecretGetRequest) SetContractNum(contractNum string) error {
    r.contractNum = contractNum
    r.Set("contract_num", contractNum)
    return nil
}

func (r AlibabaNazcaTokenFilesecretGetRequest) GetContractNum() string {
    return r.contractNum
}

