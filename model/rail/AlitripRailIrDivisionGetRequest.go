package rail

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
国际火车票标准城市查询 APIRequest
alitrip.rail.ir.division.get

国际火车票提供给代理商用于查询标准城市信息，全部城市数据量209530条，含除中国大陆以外的全部海外区域。
代理商通过分页查询的方式，拉取飞猪平台方全部境外标准城市，用于自身城市与飞猪平台城市的映射。
*/
type AlitripRailIrDivisionGetRequest struct {
    model.Params

    // 代理商id
    agentId   int64 

    // 层级，1洲，2是国家，3是省，4是市，5是区，6是街道/镇，7是村，8是逻辑行政区，境外火车票业务只需要市级别，传4就可以
    level   int64 

    // 每页条数
    pageSize   int64 

    // 页数，从1开始
    pageIndex   int64 

}

func NewAlitripRailIrDivisionGetRequest() *AlitripRailIrDivisionGetRequest{
    return &AlitripRailIrDivisionGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripRailIrDivisionGetRequest) GetApiMethodName() string {
    return "alitrip.rail.ir.division.get"
}

func (r AlitripRailIrDivisionGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripRailIrDivisionGetRequest) SetAgentId(agentId int64) error {
    r.agentId = agentId
    r.Set("agent_id", agentId)
    return nil
}

func (r AlitripRailIrDivisionGetRequest) GetAgentId() int64 {
    return r.agentId
}

func (r *AlitripRailIrDivisionGetRequest) SetLevel(level int64) error {
    r.level = level
    r.Set("level", level)
    return nil
}

func (r AlitripRailIrDivisionGetRequest) GetLevel() int64 {
    return r.level
}

func (r *AlitripRailIrDivisionGetRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r AlitripRailIrDivisionGetRequest) GetPageSize() int64 {
    return r.pageSize
}

func (r *AlitripRailIrDivisionGetRequest) SetPageIndex(pageIndex int64) error {
    r.pageIndex = pageIndex
    r.Set("page_index", pageIndex)
    return nil
}

func (r AlitripRailIrDivisionGetRequest) GetPageIndex() int64 {
    return r.pageIndex
}

