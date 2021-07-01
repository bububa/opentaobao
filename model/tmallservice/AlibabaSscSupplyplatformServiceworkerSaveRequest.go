package tmallservice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
服务商绑定工人 API请求
alibaba.ssc.supplyplatform.serviceworker.save

服务商将上传工人与服务商自己建立关系，需要将工人的服务区域和住址回传
*/
type AlibabaSscSupplyplatformServiceworkerSaveAPIRequest struct {
    model.Params
    // 工人保存参数
    _workerSaveForTopReqDto   *WorkerSaveForTopReqDTO
}

// 初始化AlibabaSscSupplyplatformServiceworkerSaveAPIRequest对象
func NewAlibabaSscSupplyplatformServiceworkerSaveRequest() *AlibabaSscSupplyplatformServiceworkerSaveAPIRequest{
    return &AlibabaSscSupplyplatformServiceworkerSaveAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaSscSupplyplatformServiceworkerSaveAPIRequest) GetApiMethodName() string {
    return "alibaba.ssc.supplyplatform.serviceworker.save"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaSscSupplyplatformServiceworkerSaveAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// WorkerSaveForTopReqDto Setter
// 工人保存参数
func (r *AlibabaSscSupplyplatformServiceworkerSaveAPIRequest) SetWorkerSaveForTopReqDto(_workerSaveForTopReqDto *WorkerSaveForTopReqDTO) error {
    r._workerSaveForTopReqDto = _workerSaveForTopReqDto
    r.Set("worker_save_for_top_req_dto", _workerSaveForTopReqDto)
    return nil
}

// WorkerSaveForTopReqDto Getter
func (r AlibabaSscSupplyplatformServiceworkerSaveAPIRequest) GetWorkerSaveForTopReqDto() *WorkerSaveForTopReqDTO {
    return r._workerSaveForTopReqDto
}
