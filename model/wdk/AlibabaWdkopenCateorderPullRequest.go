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
    _storeId   string
    // 回传状态,PREPARING,准备中，制作中；PRODUCE_FINISH，制作完成；FETCHED 已取餐；  CANCEL，加工失败/取消
    _status   string
    // 主站主订单ID
    _outOrderId   string
    // 主站子订单ID列表, 为空则表示回传整单状态
    _subOutOrderIds   []string
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
func (r *AlibabaWdkopenCateorderPullRequest) SetStoreId(_storeId string) error {
    r._storeId = _storeId
    r.Set("store_id", _storeId)
    return nil
}

// StoreId Getter
func (r AlibabaWdkopenCateorderPullRequest) GetStoreId() string {
    return r._storeId
}
// Status Setter
// 回传状态,PREPARING,准备中，制作中；PRODUCE_FINISH，制作完成；FETCHED 已取餐；  CANCEL，加工失败/取消
func (r *AlibabaWdkopenCateorderPullRequest) SetStatus(_status string) error {
    r._status = _status
    r.Set("status", _status)
    return nil
}

// Status Getter
func (r AlibabaWdkopenCateorderPullRequest) GetStatus() string {
    return r._status
}
// OutOrderId Setter
// 主站主订单ID
func (r *AlibabaWdkopenCateorderPullRequest) SetOutOrderId(_outOrderId string) error {
    r._outOrderId = _outOrderId
    r.Set("out_order_id", _outOrderId)
    return nil
}

// OutOrderId Getter
func (r AlibabaWdkopenCateorderPullRequest) GetOutOrderId() string {
    return r._outOrderId
}
// SubOutOrderIds Setter
// 主站子订单ID列表, 为空则表示回传整单状态
func (r *AlibabaWdkopenCateorderPullRequest) SetSubOutOrderIds(_subOutOrderIds []string) error {
    r._subOutOrderIds = _subOutOrderIds
    r.Set("sub_out_order_ids", _subOutOrderIds)
    return nil
}

// SubOutOrderIds Getter
func (r AlibabaWdkopenCateorderPullRequest) GetSubOutOrderIds() []string {
    return r._subOutOrderIds
}
