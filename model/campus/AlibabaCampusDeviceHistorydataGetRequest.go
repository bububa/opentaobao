package campus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
设备历史数据批量获取 APIRequest
alibaba.campus.device.historydata.get

设备历史数据批量获取
*/
type AlibabaCampusDeviceHistorydataGetRequest struct {
    model.Params

    // workbench
    workBenchContext   *WorkBenchContext 

    // 查询条件
    query   *DeviceHistoryBatchQuery 

}

func NewAlibabaCampusDeviceHistorydataGetRequest() *AlibabaCampusDeviceHistorydataGetRequest{
    return &AlibabaCampusDeviceHistorydataGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaCampusDeviceHistorydataGetRequest) GetApiMethodName() string {
    return "alibaba.campus.device.historydata.get"
}

func (r AlibabaCampusDeviceHistorydataGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaCampusDeviceHistorydataGetRequest) SetWorkBenchContext(workBenchContext *WorkBenchContext) error {
    r.workBenchContext = workBenchContext
    r.Set("work_bench_context", workBenchContext)
    return nil
}

func (r AlibabaCampusDeviceHistorydataGetRequest) GetWorkBenchContext() *WorkBenchContext {
    return r.workBenchContext
}

func (r *AlibabaCampusDeviceHistorydataGetRequest) SetQuery(query *DeviceHistoryBatchQuery) error {
    r.query = query
    r.Set("query", query)
    return nil
}

func (r AlibabaCampusDeviceHistorydataGetRequest) GetQuery() *DeviceHistoryBatchQuery {
    return r.query
}

