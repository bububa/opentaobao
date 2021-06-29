package nrt

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
摊位信息同步 API请求
tmall.nrt.stall.synchronize

摊位信息同步
*/
type TmallNrtStallSynchronizeRequest struct {
    model.Params
    // 参数对象
    stall   *NrtStoreDTO
}

// 初始化TmallNrtStallSynchronizeRequest对象
func NewTmallNrtStallSynchronizeRequest() *TmallNrtStallSynchronizeRequest{
    return &TmallNrtStallSynchronizeRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallNrtStallSynchronizeRequest) GetApiMethodName() string {
    return "tmall.nrt.stall.synchronize"
}

// IRequest interface 方法, 获取API参数
func (r TmallNrtStallSynchronizeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Stall Setter
// 参数对象
func (r *TmallNrtStallSynchronizeRequest) SetStall(stall *NrtStoreDTO) error {
    r.stall = stall
    r.Set("stall", stall)
    return nil
}

// Stall Getter
func (r TmallNrtStallSynchronizeRequest) GetStall() *NrtStoreDTO {
    return r.stall
}
