package tmallnr

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallnrnoticegoodsreadyAPIRequest 同步天猫配送人员信息 API请求
// tmall.nr.notice.goods.ready
//
// 接收商家的配送人员信息，和第三公司信息及提货码
type TmallnrnoticegoodsreadyAPIRequest struct {
	model.Params
	// 入参
	_param0 *NrZqsGoodsReadyReqDto
}

// NewTmallnrnoticegoodsreadyRequest 初始化TmallnrnoticegoodsreadyAPIRequest对象
func NewTmallnrnoticegoodsreadyRequest() *TmallnrnoticegoodsreadyAPIRequest {
	return &TmallnrnoticegoodsreadyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallnrnoticegoodsreadyAPIRequest) GetApiMethodName() string {
	return "tmall.nr.notice.goods.ready"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallnrnoticegoodsreadyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallnrnoticegoodsreadyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 入参
func (r *TmallnrnoticegoodsreadyAPIRequest) SetParam0(_param0 *NrZqsGoodsReadyReqDto) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r TmallnrnoticegoodsreadyAPIRequest) GetParam0() *NrZqsGoodsReadyReqDto {
	return r._param0
}
