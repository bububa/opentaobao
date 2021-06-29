package nazca

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取文件秘钥 API请求
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

// 初始化AlibabaNazcaTokenFilesecretGetRequest对象
func NewAlibabaNazcaTokenFilesecretGetRequest() *AlibabaNazcaTokenFilesecretGetRequest{
    return &AlibabaNazcaTokenFilesecretGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaNazcaTokenFilesecretGetRequest) GetApiMethodName() string {
    return "alibaba.nazca.token.filesecret.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaNazcaTokenFilesecretGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// PlatformUserId Setter
// 客户在1688的唯一标识
func (r *AlibabaNazcaTokenFilesecretGetRequest) SetPlatformUserId(platformUserId string) error {
    r.platformUserId = platformUserId
    r.Set("platform_user_id", platformUserId)
    return nil
}

// PlatformUserId Getter
func (r AlibabaNazcaTokenFilesecretGetRequest) GetPlatformUserId() string {
    return r.platformUserId
}
// ContractNum Setter
// 合同编号
func (r *AlibabaNazcaTokenFilesecretGetRequest) SetContractNum(contractNum string) error {
    r.contractNum = contractNum
    r.Set("contract_num", contractNum)
    return nil
}

// ContractNum Getter
func (r AlibabaNazcaTokenFilesecretGetRequest) GetContractNum() string {
    return r.contractNum
}
