package omniorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoomniorderprintsalejudgeAPIRequest 导购员判断 API请求
// taobao.omniorder.print.sale.judge
//
// 用于判断当前子账号是否导购员
type TaobaoomniorderprintsalejudgeAPIRequest struct {
	model.Params
	// 用户子账号ID
	_subUid int64
}

// NewTaobaoomniorderprintsalejudgeRequest 初始化TaobaoomniorderprintsalejudgeAPIRequest对象
func NewTaobaoomniorderprintsalejudgeRequest() *TaobaoomniorderprintsalejudgeAPIRequest {
	return &TaobaoomniorderprintsalejudgeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoomniorderprintsalejudgeAPIRequest) GetApiMethodName() string {
	return "taobao.omniorder.print.sale.judge"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoomniorderprintsalejudgeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoomniorderprintsalejudgeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSubUid is SubUid Setter
// 用户子账号ID
func (r *TaobaoomniorderprintsalejudgeAPIRequest) SetSubUid(_subUid int64) error {
	r._subUid = _subUid
	r.Set("sub_uid", _subUid)
	return nil
}

// GetSubUid SubUid Getter
func (r TaobaoomniorderprintsalejudgeAPIRequest) GetSubUid() int64 {
	return r._subUid
}
