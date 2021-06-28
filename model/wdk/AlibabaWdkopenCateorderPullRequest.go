package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商户回传餐饮加工单状态 APIRequest
alibaba.wdkopen.cateorder.pull

商户回传餐饮加工单状态
*/
type AlibabaWdkopenCateorderPullRequest struct {
    model.Params

    // 经营店ID
    storeId   string 

    // 回传状态,PREPARING,准备中，制作中；PRODUCE_FINISH，制作完成；FETCHED 已取餐；  CANCEL，加工失败/取消
    status   string 

    // 主站主订单ID
    outOrderId   string 

    // 主站子订单ID列表, 为空则表示回传整单状态
    subOutOrderIds   []string 

}

func NewAlibabaWdkopenCateorderPullRequest() *AlibabaWdkopenCateorderPullRequest{
    return &AlibabaWdkopenCateorderPullRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkopenCateorderPullRequest) GetApiMethodName() string {
    return "alibaba.wdkopen.cateorder.pull"
}

func (r AlibabaWdkopenCateorderPullRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkopenCateorderPullRequest) SetStoreId(storeId string) error {
    r.storeId = storeId
    r.Set("store_id", storeId)
    return nil
}

func (r AlibabaWdkopenCateorderPullRequest) GetStoreId() string {
    return r.storeId
}

func (r *AlibabaWdkopenCateorderPullRequest) SetStatus(status string) error {
    r.status = status
    r.Set("status", status)
    return nil
}

func (r AlibabaWdkopenCateorderPullRequest) GetStatus() string {
    return r.status
}

func (r *AlibabaWdkopenCateorderPullRequest) SetOutOrderId(outOrderId string) error {
    r.outOrderId = outOrderId
    r.Set("out_order_id", outOrderId)
    return nil
}

func (r AlibabaWdkopenCateorderPullRequest) GetOutOrderId() string {
    return r.outOrderId
}

func (r *AlibabaWdkopenCateorderPullRequest) SetSubOutOrderIds(subOutOrderIds []string) error {
    r.subOutOrderIds = subOutOrderIds
    r.Set("sub_out_order_ids", subOutOrderIds)
    return nil
}

func (r AlibabaWdkopenCateorderPullRequest) GetSubOutOrderIds() []string {
    return r.subOutOrderIds
}

