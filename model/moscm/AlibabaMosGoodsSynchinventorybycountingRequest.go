package moscm

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
以盘点方式调整库存：传入商品实际库存 API请求
alibaba.mos.goods.synchinventorybycounting

以盘点方式调整库存：传入商品实际库存
盘点单自动判断数量增减
*/
type AlibabaMosGoodsSynchinventorybycountingRequest struct {
    model.Params
    // 专柜信息
    _paramCountingInfoDto   *CountingInfoDto
    // 盘点库存项（最大列表长度：20）
    _countingItemDto   []CountingItemDto
}

// 初始化AlibabaMosGoodsSynchinventorybycountingRequest对象
func NewAlibabaMosGoodsSynchinventorybycountingRequest() *AlibabaMosGoodsSynchinventorybycountingRequest{
    return &AlibabaMosGoodsSynchinventorybycountingRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaMosGoodsSynchinventorybycountingRequest) GetApiMethodName() string {
    return "alibaba.mos.goods.synchinventorybycounting"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaMosGoodsSynchinventorybycountingRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamCountingInfoDto Setter
// 专柜信息
func (r *AlibabaMosGoodsSynchinventorybycountingRequest) SetParamCountingInfoDto(_paramCountingInfoDto *CountingInfoDto) error {
    r._paramCountingInfoDto = _paramCountingInfoDto
    r.Set("param_counting_info_dto", _paramCountingInfoDto)
    return nil
}

// ParamCountingInfoDto Getter
func (r AlibabaMosGoodsSynchinventorybycountingRequest) GetParamCountingInfoDto() *CountingInfoDto {
    return r._paramCountingInfoDto
}
// CountingItemDto Setter
// 盘点库存项（最大列表长度：20）
func (r *AlibabaMosGoodsSynchinventorybycountingRequest) SetCountingItemDto(_countingItemDto []CountingItemDto) error {
    r._countingItemDto = _countingItemDto
    r.Set("counting_item_dto", _countingItemDto)
    return nil
}

// CountingItemDto Getter
func (r AlibabaMosGoodsSynchinventorybycountingRequest) GetCountingItemDto() []CountingItemDto {
    return r._countingItemDto
}
