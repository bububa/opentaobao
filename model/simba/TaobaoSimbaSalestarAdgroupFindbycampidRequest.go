package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
(销量明星)批量获取推广计划下的推广组信息 API请求
taobao.simba.salestar.adgroup.findbycampid

批量得到推广计划下的推广组
*/
type TaobaoSimbaSalestarAdgroupFindbycampidRequest struct {
    model.Params
    // 推广计划Id
    campaignId   int64
    // 页尺寸，最大200，如果入参adgroup_ids有传入值，则page_size和page_no值不起作用。如果adgrpup_ids为空而campaign_id有值，此时page_size和page_no值才是返回的页数据大小和页码
    pageSize   int64
    // 页码，从1开始
    pageNo   int64
}

// 初始化TaobaoSimbaSalestarAdgroupFindbycampidRequest对象
func NewTaobaoSimbaSalestarAdgroupFindbycampidRequest() *TaobaoSimbaSalestarAdgroupFindbycampidRequest{
    return &TaobaoSimbaSalestarAdgroupFindbycampidRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoSimbaSalestarAdgroupFindbycampidRequest) GetApiMethodName() string {
    return "taobao.simba.salestar.adgroup.findbycampid"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoSimbaSalestarAdgroupFindbycampidRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CampaignId Setter
// 推广计划Id
func (r *TaobaoSimbaSalestarAdgroupFindbycampidRequest) SetCampaignId(campaignId int64) error {
    r.campaignId = campaignId
    r.Set("campaign_id", campaignId)
    return nil
}

// CampaignId Getter
func (r TaobaoSimbaSalestarAdgroupFindbycampidRequest) GetCampaignId() int64 {
    return r.campaignId
}
// PageSize Setter
// 页尺寸，最大200，如果入参adgroup_ids有传入值，则page_size和page_no值不起作用。如果adgrpup_ids为空而campaign_id有值，此时page_size和page_no值才是返回的页数据大小和页码
func (r *TaobaoSimbaSalestarAdgroupFindbycampidRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoSimbaSalestarAdgroupFindbycampidRequest) GetPageSize() int64 {
    return r.pageSize
}
// PageNo Setter
// 页码，从1开始
func (r *TaobaoSimbaSalestarAdgroupFindbycampidRequest) SetPageNo(pageNo int64) error {
    r.pageNo = pageNo
    r.Set("page_no", pageNo)
    return nil
}

// PageNo Getter
func (r TaobaoSimbaSalestarAdgroupFindbycampidRequest) GetPageNo() int64 {
    return r.pageNo
}
