package baichuanctg

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaBaichuanCtgToutiaoContentAPIRequest
微博输出头条数据 API请求
alibaba.baichuan.ctg.toutiao.content

百川头条内容获取 */
type AlibabaBaichuanCtgToutiaoContentAPIRequest struct {
	model.Params
	// param0
	_param0 *CtgRequest
}

// NewAlibabaBaichuanCtgToutiaoContentRequest 初始化AlibabaBaichuanCtgToutiaoContentAPIRequest对象
func NewAlibabaBaichuanCtgToutiaoContentRequest() *AlibabaBaichuanCtgToutiaoContentAPIRequest {
	return &AlibabaBaichuanCtgToutiaoContentAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaBaichuanCtgToutiaoContentAPIRequest) GetApiMethodName() string {
	return "alibaba.baichuan.ctg.toutiao.content"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaBaichuanCtgToutiaoContentAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Param0 Setter
// param0
func (r *AlibabaBaichuanCtgToutiaoContentAPIRequest) SetParam0(_param0 *CtgRequest) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// Get Param0 Getter
func (r AlibabaBaichuanCtgToutiaoContentAPIRequest) GetParam0() *CtgRequest {
	return r._param0
}
