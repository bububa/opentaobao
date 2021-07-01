package tmallnr

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
同城配物流信息回传 API请求
tmall.nr.fulfill.logistics.sync

同城配业务物流信息回传，通过接口将物流信息同步给天猫
*/
type TmallNrFulfillLogisticsSyncAPIRequest struct {
    model.Params
    // 物流回传参数
    _param0   *NrLogisticsInfoSynReqDTO
}

// 初始化TmallNrFulfillLogisticsSyncAPIRequest对象
func NewTmallNrFulfillLogisticsSyncRequest() *TmallNrFulfillLogisticsSyncAPIRequest{
    return &TmallNrFulfillLogisticsSyncAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallNrFulfillLogisticsSyncAPIRequest) GetApiMethodName() string {
    return "tmall.nr.fulfill.logistics.sync"
}

// IRequest interface 方法, 获取API参数
func (r TmallNrFulfillLogisticsSyncAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param0 Setter
// 物流回传参数
func (r *TmallNrFulfillLogisticsSyncAPIRequest) SetParam0(_param0 *NrLogisticsInfoSynReqDTO) error {
    r._param0 = _param0
    r.Set("param0", _param0)
    return nil
}

// Param0 Getter
func (r TmallNrFulfillLogisticsSyncAPIRequest) GetParam0() *NrLogisticsInfoSynReqDTO {
    return r._param0
}
