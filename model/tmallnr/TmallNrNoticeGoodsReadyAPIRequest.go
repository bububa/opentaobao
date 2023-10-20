package tmallnr

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallNrNoticeGoodsReadyAPIRequest) Reset() {
	r._param0 = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallNrNoticeGoodsReadyAPIRequest) GetApiMethodName() string {
	return "tmall.nr.notice.goods.ready"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallNrNoticeGoodsReadyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallNrNoticeGoodsReadyAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolTmallNrNoticeGoodsReadyAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallNrNoticeGoodsReadyRequest()
	},
}

// GetTmallNrNoticeGoodsReadyRequest 从 sync.Pool 获取 TmallNrNoticeGoodsReadyAPIRequest
func GetTmallNrNoticeGoodsReadyAPIRequest() *TmallNrNoticeGoodsReadyAPIRequest {
	return poolTmallNrNoticeGoodsReadyAPIRequest.Get().(*TmallNrNoticeGoodsReadyAPIRequest)
}

// ReleaseTmallNrNoticeGoodsReadyAPIRequest 将 TmallNrNoticeGoodsReadyAPIRequest 放入 sync.Pool
func ReleaseTmallNrNoticeGoodsReadyAPIRequest(v *TmallNrNoticeGoodsReadyAPIRequest) {
	v.Reset()
	poolTmallNrNoticeGoodsReadyAPIRequest.Put(v)
}
