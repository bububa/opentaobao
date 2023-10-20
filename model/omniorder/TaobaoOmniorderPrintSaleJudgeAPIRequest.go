package omniorder

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOmniorderPrintSaleJudgeAPIRequest 导购员判断 API请求
// taobao.omniorder.print.sale.judge
//
// 用于判断当前子账号是否导购员
type TaobaoOmniorderPrintSaleJudgeAPIRequest struct {
	model.Params
	// 用户子账号ID
	_subUid int64
}

// NewTaobaoOmniorderPrintSaleJudgeRequest 初始化TaobaoOmniorderPrintSaleJudgeAPIRequest对象
func NewTaobaoOmniorderPrintSaleJudgeRequest() *TaobaoOmniorderPrintSaleJudgeAPIRequest {
	return &TaobaoOmniorderPrintSaleJudgeAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoOmniorderPrintSaleJudgeAPIRequest) Reset() {
	r._subUid = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOmniorderPrintSaleJudgeAPIRequest) GetApiMethodName() string {
	return "taobao.omniorder.print.sale.judge"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOmniorderPrintSaleJudgeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoOmniorderPrintSaleJudgeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSubUid is SubUid Setter
// 用户子账号ID
func (r *TaobaoOmniorderPrintSaleJudgeAPIRequest) SetSubUid(_subUid int64) error {
	r._subUid = _subUid
	r.Set("sub_uid", _subUid)
	return nil
}

// GetSubUid SubUid Getter
func (r TaobaoOmniorderPrintSaleJudgeAPIRequest) GetSubUid() int64 {
	return r._subUid
}

var poolTaobaoOmniorderPrintSaleJudgeAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoOmniorderPrintSaleJudgeRequest()
	},
}

// GetTaobaoOmniorderPrintSaleJudgeRequest 从 sync.Pool 获取 TaobaoOmniorderPrintSaleJudgeAPIRequest
func GetTaobaoOmniorderPrintSaleJudgeAPIRequest() *TaobaoOmniorderPrintSaleJudgeAPIRequest {
	return poolTaobaoOmniorderPrintSaleJudgeAPIRequest.Get().(*TaobaoOmniorderPrintSaleJudgeAPIRequest)
}

// ReleaseTaobaoOmniorderPrintSaleJudgeAPIRequest 将 TaobaoOmniorderPrintSaleJudgeAPIRequest 放入 sync.Pool
func ReleaseTaobaoOmniorderPrintSaleJudgeAPIRequest(v *TaobaoOmniorderPrintSaleJudgeAPIRequest) {
	v.Reset()
	poolTaobaoOmniorderPrintSaleJudgeAPIRequest.Put(v)
}
