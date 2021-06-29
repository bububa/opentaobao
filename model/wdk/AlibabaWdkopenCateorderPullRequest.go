package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商户回传餐饮加工单状态 API请求
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

// 初始化AlibabaWdkopenCateorderPullRequest对象
func NewAlibabaWdkopenCateorderPullRequest() *AlibabaWdkopenCateorderPullRequest{
    return &AlibabaWdkopenCateorderPullRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkopenCateorderPullRequest) GetApiMethodName() string {
    return "alibaba.wdkopen.cateorder.pull"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkopenCateorderPullRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// StoreId Setter
// 经营店ID
func (r *AlibabaWdkopenCateorderPullRequest) SetStoreId(storeId string) error {
    r.storeId = storeId
    r.Set("store_id", storeId)
    return nil
}

// StoreId Getter
func (r AlibabaWdkopenCateorderPullRequest) GetStoreId() string {
    return r.storeId
}
// Status Setter
// 回传状态,PREPARING,准备中，制作中；PRODUCE_FINISH，制作完成；FETCHED 已取餐；  CANCEL，加工失败/取消
func (r *AlibabaWdkopenCateorderPullRequest) SetStatus(status string) error {
    r.status = status
    r.Set("status", status)
    return nil
}

// Status Getter
func (r AlibabaWdkopenCateorderPullRequest) GetStatus() string {
    return r.status
}
// OutOrderId Setter
// 主站主订单ID
func (r *AlibabaWdkopenCateorderPullRequest) SetOutOrderId(outOrderId string) error {
    r.outOrderId = outOrderId
    r.Set("out_order_id", outOrderId)
    return nil
}

// OutOrderId Getter
func (r AlibabaWdkopenCateorderPullRequest) GetOutOrderId() string {
    return r.outOrderId
}
// SubOutOrderIds Setter
// 主站子订单ID列表, 为空则表示回传整单状态
func (r *AlibabaWdkopenCateorderPullRequest) SetSubOutOrderIds(subOutOrderIds []string) error {
    r.subOutOrderIds = subOutOrderIds
    r.Set("sub_out_order_ids", subOutOrderIds)
    return nil
}

// SubOutOrderIds Getter
func (r AlibabaWdkopenCateorderPullRequest) GetSubOutOrderIds() []string {
    return r.subOutOrderIds
}
