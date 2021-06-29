package tmallnr

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
同步天猫配送人员信息 API请求
tmall.nr.notice.goods.ready

接收商家的配送人员信息，和第三公司信息及提货码
*/
type TmallNrNoticeGoodsReadyRequest struct {
    model.Params
    // 入参
    param0   *NrZqsGoodsReadyReqDto
}

// 初始化TmallNrNoticeGoodsReadyRequest对象
func NewTmallNrNoticeGoodsReadyRequest() *TmallNrNoticeGoodsReadyRequest{
    return &TmallNrNoticeGoodsReadyRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallNrNoticeGoodsReadyRequest) GetApiMethodName() string {
    return "tmall.nr.notice.goods.ready"
}

// IRequest interface 方法, 获取API参数
func (r TmallNrNoticeGoodsReadyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param0 Setter
// 入参
func (r *TmallNrNoticeGoodsReadyRequest) SetParam0(param0 *NrZqsGoodsReadyReqDto) error {
    r.param0 = param0
    r.Set("param0", param0)
    return nil
}

// Param0 Getter
func (r TmallNrNoticeGoodsReadyRequest) GetParam0() *NrZqsGoodsReadyReqDto {
    return r.param0
}
