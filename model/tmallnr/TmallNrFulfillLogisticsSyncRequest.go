package tmallnr

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
同城配物流信息回传 APIRequest
tmall.nr.fulfill.logistics.sync

同城配业务物流信息回传，通过接口将物流信息同步给天猫
*/
type TmallNrFulfillLogisticsSyncRequest struct {
    model.Params

    // 物流回传参数
    param0   *NrLogisticsInfoSynReqDto 

}

func NewTmallNrFulfillLogisticsSyncRequest() *TmallNrFulfillLogisticsSyncRequest{
    return &TmallNrFulfillLogisticsSyncRequest{
        Params: model.NewParams(),
    }
}

func (r TmallNrFulfillLogisticsSyncRequest) GetApiMethodName() string {
    return "tmall.nr.fulfill.logistics.sync"
}

func (r TmallNrFulfillLogisticsSyncRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallNrFulfillLogisticsSyncRequest) SetParam0(param0 *NrLogisticsInfoSynReqDto) error {
    r.param0 = param0
    r.Set("param0", param0)
    return nil
}

func (r TmallNrFulfillLogisticsSyncRequest) GetParam0() *NrLogisticsInfoSynReqDto {
    return r.param0
}

