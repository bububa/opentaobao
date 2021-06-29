package baichuanctg

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
微博输出头条数据 API请求
alibaba.baichuan.ctg.toutiao.content

百川头条内容获取
*/
type AlibabaBaichuanCtgToutiaoContentRequest struct {
    model.Params
    // param0
    param0   *CtgRequest
}

// 初始化AlibabaBaichuanCtgToutiaoContentRequest对象
func NewAlibabaBaichuanCtgToutiaoContentRequest() *AlibabaBaichuanCtgToutiaoContentRequest{
    return &AlibabaBaichuanCtgToutiaoContentRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaBaichuanCtgToutiaoContentRequest) GetApiMethodName() string {
    return "alibaba.baichuan.ctg.toutiao.content"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaBaichuanCtgToutiaoContentRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param0 Setter
// param0
func (r *AlibabaBaichuanCtgToutiaoContentRequest) SetParam0(param0 *CtgRequest) error {
    r.param0 = param0
    r.Set("param0", param0)
    return nil
}

// Param0 Getter
func (r AlibabaBaichuanCtgToutiaoContentRequest) GetParam0() *CtgRequest {
    return r.param0
}
