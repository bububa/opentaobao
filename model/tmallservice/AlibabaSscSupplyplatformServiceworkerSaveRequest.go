package tmallservice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
服务商绑定工人 APIRequest
alibaba.ssc.supplyplatform.serviceworker.save

服务商将上传工人与服务商自己建立关系，需要将工人的服务区域和住址回传
*/
type AlibabaSscSupplyplatformServiceworkerSaveRequest struct {
    model.Params

    // 工人保存参数
    workerSaveForTopReqDto   *WorkerSaveForTopReqDto 

}

func NewAlibabaSscSupplyplatformServiceworkerSaveRequest() *AlibabaSscSupplyplatformServiceworkerSaveRequest{
    return &AlibabaSscSupplyplatformServiceworkerSaveRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaSscSupplyplatformServiceworkerSaveRequest) GetApiMethodName() string {
    return "alibaba.ssc.supplyplatform.serviceworker.save"
}

func (r AlibabaSscSupplyplatformServiceworkerSaveRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaSscSupplyplatformServiceworkerSaveRequest) SetWorkerSaveForTopReqDto(workerSaveForTopReqDto *WorkerSaveForTopReqDto) error {
    r.workerSaveForTopReqDto = workerSaveForTopReqDto
    r.Set("worker_save_for_top_req_dto", workerSaveForTopReqDto)
    return nil
}

func (r AlibabaSscSupplyplatformServiceworkerSaveRequest) GetWorkerSaveForTopReqDto() *WorkerSaveForTopReqDto {
    return r.workerSaveForTopReqDto
}

