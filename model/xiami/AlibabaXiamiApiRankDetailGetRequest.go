package xiami

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
排行榜详情 API请求
alibaba.xiami.api.rank.detail.get

虾米排行榜详情数据
*/
type AlibabaXiamiApiRankDetailGetRequest struct {
    model.Params
    // 榜单ID
    _billboardId   int64
    // 调用来源
    _bizCode   string
}

// 初始化AlibabaXiamiApiRankDetailGetRequest对象
func NewAlibabaXiamiApiRankDetailGetRequest() *AlibabaXiamiApiRankDetailGetRequest{
    return &AlibabaXiamiApiRankDetailGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaXiamiApiRankDetailGetRequest) GetApiMethodName() string {
    return "alibaba.xiami.api.rank.detail.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaXiamiApiRankDetailGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// BillboardId Setter
// 榜单ID
func (r *AlibabaXiamiApiRankDetailGetRequest) SetBillboardId(_billboardId int64) error {
    r._billboardId = _billboardId
    r.Set("billboard_id", _billboardId)
    return nil
}

// BillboardId Getter
func (r AlibabaXiamiApiRankDetailGetRequest) GetBillboardId() int64 {
    return r._billboardId
}
// BizCode Setter
// 调用来源
func (r *AlibabaXiamiApiRankDetailGetRequest) SetBizCode(_bizCode string) error {
    r._bizCode = _bizCode
    r.Set("biz_code", _bizCode)
    return nil
}

// BizCode Getter
func (r AlibabaXiamiApiRankDetailGetRequest) GetBizCode() string {
    return r._bizCode
}
