package servicecenter

import (
    "github.com/bububa/opentaobao/model"
)

/* 
合同下载 APIResponse
tmall.car.contract.download

目前天猫开新车会在线上签署一份合同，协议，需要一个个在已卖出打开，另存为pdf，人工一个个下载比较麻烦，期望通过接口直接读取pdf；
因为比较耗时，建议一个个下载，假设并发下载，很可能限流，每天的调用量有限；
*/
type TmallCarContractDownloadAPIResponse struct {
    model.CommonResponse
    Response *TmallCarContractDownloadResponse `json:"tmall_car_contract_download_response,omitempty"`
}

type TmallCarContractDownloadResponse struct {

    // 结果
    Result   *TmallCarContractDownloadResult `json:"result,omitempty"`

}
