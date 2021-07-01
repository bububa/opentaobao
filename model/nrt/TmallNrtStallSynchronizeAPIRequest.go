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
type TmallNrtStallSynchronizeAPIRequest struct {
    model.Params
    // 参数对象
    _stall   *NrtStoreDto
}

// 初始化TmallNrtStallSynchronizeAPIRequest对象
func NewTmallNrtStallSynchronizeRequest() *TmallNrtStallSynchronizeAPIRequest{
    return &TmallNrtStallSynchronizeAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallNrtStallSynchronizeAPIRequest) GetApiMethodName() string {
    return "tmall.nrt.stall.synchronize"
}

// IRequest interface 方法, 获取API参数
func (r TmallNrtStallSynchronizeAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Stall Setter
// 参数对象
func (r *TmallNrtStallSynchronizeAPIRequest) SetStall(_stall *NrtStoreDto) error {
    r._stall = _stall
    r.Set("stall", _stall)
    return nil
}

// Stall Getter
func (r TmallNrtStallSynchronizeAPIRequest) GetStall() *NrtStoreDto {
    return r._stall
}
