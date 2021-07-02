package tmallnr

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallNrNoticeGoodsReadyAPIRequest 同步天猫配送人员信息 API请求
// tmall.nr.notice.goods.ready
//
// 接收商家的配送人员信息，和第三公司信息及提货码
type TmallNrNoticeGoodsReadyAPIRequest struct {
	model.Params
	// 入参
	_param0 *NrZqsGoodsReadyReqDto
}

// NewTmallNrNoticeGoodsReadyRequest 初始化TmallNrNoticeGoodsReadyAPIRequest对象
func NewTmallNrNoticeGoodsReadyRequest() *TmallNrNoticeGoodsReadyAPIRequest {
	return &TmallNrNoticeGoodsReadyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallNrNoticeGoodsReadyAPIRequest) GetApiMethodName() string {
	return "tmall.nr.notice.goods.ready"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallNrNoticeGoodsReadyAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetParam0 is Param0 Setter
// 入参
func (r *TmallNrNoticeGoodsReadyAPIRequest) SetParam0(_param0 *NrZqsGoodsReadyReqDto) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r TmallNrNoticeGoodsReadyAPIRequest) GetParam0() *NrZqsGoodsReadyReqDto {
	return r._param0
}
